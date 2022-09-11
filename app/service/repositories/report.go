package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Report(ctx echo.Context, conn *gorm.DB) (*[]entity.Report, error) {
	var report []entity.Report
	sql :=
		`
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r.successDate as successDate ,
		r.reportDate as reportDate ,
		r.createdBy as createdBy 
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	`
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportByCreatedBy(ctx echo.Context, conn *gorm.DB, req request.ReportByCreatedBy) (*[]entity.Report, error) {
	var report []entity.Report
	sql := fmt.Sprintf(`	
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r.successDate as successDate ,
		r.reportDate as reportDate ,
		r.createdBy as createdBy 
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	WHERE 
		r.createdBy = %v`, req.CreatedBy)
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportById(ctx echo.Context, conn *gorm.DB, req request.Report) (*entity.Report, error) {
	var report entity.Report
	sql := fmt.Sprintf(`	
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r.successDate as successDate ,
		r.reportDate as reportDate ,
		r.createdBy as createdBy 
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	WHERE 
		r.reportId = %v`, req.ReportId)
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req entity.ReportInsert) (*int, error) {
	stmt := conn.Begin()
	var err error
	err = stmt.Table("reports").Create(&req).Error
	if err != nil {
		return nil, err
	}

	var id int
	err = stmt.Table("reports").Select("reportId").Where("title = ?", req.Title).Where("createdBy = ?", req.CreatedBy).Where("reportDate = ?", req.ReportDate).Scan(&id).Error
	if err != nil {
		return nil, err
	}

	stmt.Commit()
	return &id, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) error {
	var err error
	stmt := conn.Begin()
	err = stmt.Exec("UPDATE reports SET status = ? WHERE reportId = ?", req.Status, req.ReportId).Error
	if err != nil {
		return err
	}
	if req.Status == "S2" {
		err = stmt.Exec("UPDATE historyReport SET employeeId = ? WHERE reportId = ?", req.EmployeeId, req.ReportId).Error
		if err != nil {
			return err
		}
	}
	stmt.Commit()
	return nil
}

func DeleteReportById(ctx echo.Context, conn *gorm.DB, req request.Report) error {
	var err error
	session := conn.Begin()
	err = session.Exec("DELETE FROM reviewReports WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM reportEngage WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM assignReport WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM historyReport WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM reports WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	session.Commit()
	return nil
}

func ReportByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Report, error) {
	var data *[]entity.Report
	err := conn.Table("report").Where("dormId = ?", dormId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReportEndJob(ctx echo.Context, conn *gorm.DB, req entity.EndReport) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Table("reports").Where("reportId = ?", req.ReportId).Update("status = ?", "S7").Error
	if err != nil {
		return err
	}

	err = stmt.Table("historyReport").Where("reportId = ?", req.ReportId).Update("dateOfIssue = ?", req.DateOfIssue).Error
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(
		`
		INSERT INTO reviewReports (des, reportId, score)
		VALUES (%v, %v, %v)
		`, req.Description, req.ReportId, req.Score)
	err = stmt.Exec(sql).Error
	if err != nil {
		return err
	}

	stmt.Commit()
	return nil
}
