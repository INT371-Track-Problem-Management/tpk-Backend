package repositories

import (
	"fmt"
	"log"
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
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
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
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	WHERE 
		r.createBy = %v`, req.CreatedBy)
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
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
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
	log.Println(req)
	var err error
	err = conn.Table("reports").Create(&req).Error
	if err != nil {
		return nil, err
	}

	var id int
	err = conn.Table("reports").Select("reportId").Where("title = ?", req.Title).Where("createBy = ?", req.CreateBy).Where("createAt = ?", req.CreateAt).Scan(&id).Error
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req entity.ReportChangeStatus) error {
	err := conn.Exec("UPDATE reports SET status = ?, updateBy = ?, updateAt = ? WHERE reportId = ?", req.Status, req.EmployeeId, req.UpdateAt, req.ReportId).Error
	if err != nil {
		return err
	}
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
	err = session.Exec("DELETE FROM reportStatus WHERE reportId = ?", req.ReportId).Error
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

func ReportEndJob(ctx echo.Context, conn *gorm.DB, req entity.EndReport) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Table("reports").Where("reportId = ?", req.ReportId).Update("status = ?", "S7").Error
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

func ReportStatus(ctx echo.Context, conn *gorm.DB, status request.ReportStatus) error {
	err := conn.Table("reportStatus").Create(status).Error
	if err != nil {
		return err
	}
	return nil
}
