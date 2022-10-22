package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"

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
		r2.roomNum  as roomNum,
		r2.buildingId as buildingId,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	JOIN
		room r2 
	ON r.roomId  = r2.roomNum
	`
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportByCreatedBy(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.Report, error) {
	var report []entity.Report
	sql := fmt.Sprintf(`
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r.roomId as roomId,
		ro.roomNum as roomNum,
		ro.BuildingId as buildingId,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	JOIN
		room ro
	ON r.roomId = ro.roomId
	WHERE 
		r.createBy = %v`, customerId)
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
		r.roomId as roomId,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy,
		r.updateBy as updateBy,
		ro.roomNum as roomNum,
        ro.buildingId as buildingId
	FROM 
		reports r 
	JOIN
		statusMaster sm 
	ON r.status = sm.statusMasterId
	JOIN
		room ro
	ON r.roomId = ro.roomId
	WHERE 
		r.reportId = %v`, req.ReportId)
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req entity.ReportInsert) (*int, error) {
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

func YearConfig(ctx echo.Context, conn *gorm.DB) (*response.Year, error) {
	var year response.Year
	err := conn.Table("yearConfig").Order("year desc").Select("year").Scan(&year.Year).Error
	if err != nil {
		return nil, err
	}
	return &year, nil
}

func ReportByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*[]entity.ReportJoinEngage, error) {
	var report []entity.ReportJoinEngage
	sql := fmt.Sprintf(
		`
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport,
		r.reportDes as reportDes,
		sm.status as status,
		r2.roomNum as roomNum,
		r2.buildingId as buildingId,
		re.selectedDate as selectedDate,
		r.updateAt as updateAt,
		r.updateBy as updateBy,
		r.createAt as createAt,
		r.createBy as createBy
	FROM
		reports r
	LEFT JOIN
		statusMaster sm 
	ON
		r.status = sm.statusMasterId 
	LEFT JOIN
		reportEngage re
	ON
		r.reportId  = re.reportId 
	LEFT JOIN
		room r2
	ON
		r.roomId = r2.roomId 
	WHERE 
		r.roomId = %v
	`, roomId)
	err := conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportStatusByReportId(ctx echo.Context, conn *gorm.DB, reportId string) (*[]entity.ReportStatus, error) {
	var status []entity.ReportStatus
	sql := fmt.Sprintf(`
		SELECT 
			rs.statusId,
			rs.reportId,
			sm.status,
			rs.createAt
		FROM
			reportStatus rs
		LEFT JOIN
			statusMaster sm 
		ON
			rs.status = sm.statusMasterId 
		WHERE
			rs.reportId = %v
		ORDER BY
			rs.createAt DESC;
	`, reportId)
	err := conn.Raw(sql).Scan(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func ReportListForCustomer(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.ReportJoinEngage, error) {
	var reports []entity.ReportJoinEngage
	sql := fmt.Sprintf(`
		SELECT 
			r.reportId as reportId,
			r.title as title,
			r.categoriesReport as categoriesReport,
			r.reportDes as reportDes,
			sm.status as status,
			r2.roomNum as roomNum,
			r2.buildingId as buildingId,
			re.selectedDate as selectedDate,
			r.updateAt as updateAt,
			r.updateBy as updateBy,
			r.createAt as createAt,
			r.createBy as createBy
		FROM
			reports r
		LEFT JOIN
			statusMaster sm 
		ON
			r.status = sm.statusMasterId 
		LEFT JOIN
			reportEngage re
		ON
			r.reportId  = re.reportId 
		LEFT JOIN
			room r2
		ON
			r.roomId = r2.roomId 
		WHERE
			r.createBy = %v;
	`, customerId)
	err := conn.Raw(sql).Scan(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}
