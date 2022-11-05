package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"

	"gorm.io/gorm"
)

func (r mysqlRepository) Report() (*[]model.ReportJoinEngage, error) {
	var report []model.ReportJoinEngage
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
		re.selectedDate as selectedDate,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r
	LEFT JOIN
		statusMaster sm 
	ON 
		r.status = sm.statusMasterId
	LEFT JOIN
		room r2 
	ON 
		r.roomId  = r2.roomId
	LEFT JOIN
		reportEngage re
	ON
		r.reportId  = re.reportId 
	ORDER BY 
		r.updateAt DESC
	`
	err := r.conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r mysqlRepository) ReportByCreatedBy(customerId string) (*[]model.Report, error) {
	var report []model.Report
	sql := fmt.Sprintf(`
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r2.roomNum  as roomNum,
		r2.buildingId as buildingId,
		re.selectedDate as selectedDate,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r
	LEFT JOIN
		statusMaster sm 
	ON 
		r.status = sm.statusMasterId
	LEFT JOIN
		room r2 
	ON 
		r.roomId  = r2.roomId
	LEFT JOIN
		reportEngage re
	ON
		r.reportId  = re.reportId 
	WHERE 
		r.createBy = %v`, customerId)
	err := r.conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r mysqlRepository) ReportById(reportId int) (*model.Report, error) {
	var report model.Report
	sql := fmt.Sprintf(`	
	SELECT 
		r.reportId as reportId,
		r.title as title,
		r.categoriesReport as categoriesReport ,
		r.reportDes as reportDes ,
		sm.status as status,
		r2.roomNum  as roomNum,
		r2.buildingId as buildingId,
		re.selectedDate as selectedDate,
		r.updateAt as updateAt ,
		r.createAt as createAt ,
		r.createBy as createBy
	FROM 
		reports r
	LEFT JOIN
		statusMaster sm 
	ON 
		r.status = sm.statusMasterId
	LEFT JOIN
		room r2 
	ON 
		r.roomId  = r2.roomId
	LEFT JOIN
		reportEngage re
	ON
		r.reportId  = re.reportId 
	WHERE 
		r.reportId = %v`, reportId)
	err := r.conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r mysqlRepository) ReportInsert(req model.ReportInsert) (*int, *model.Customer, error) {
	timenow := pkg.GetDatetime()
	session := r.conn.Begin()
	reportId, err := r.reportInsert(req, session)
	if err != nil {
		return nil, nil, err
	}

	status := request.ReportStatus{
		ReportId:  *reportId,
		Status:    req.Status,
		UpdateAt:  timenow,
		UpdateBy:  req.CreateBy,
		CreatedAt: timenow,
	}

	if err := r.reportStatus(status, session); err != nil {
		return nil, nil, err
	}

	customer, err := r.GetCustomerById(req.CreateBy)
	if err != nil {
		return nil, nil, err
	}

	session.Commit()

	return reportId, customer, nil
}

func (r mysqlRepository) reportInsert(req model.ReportInsert, conn *gorm.DB) (*int, error) {
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

func (r mysqlRepository) ReportChangeStatus(req model.ReportChangeStatus) error {
	err := r.conn.Exec("UPDATE reports SET status = ?, updateBy = ?, updateAt = ? WHERE reportId = ?", req.Status, req.UpdateBy, req.UpdateAt, req.ReportId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r mysqlRepository) DeleteReportById(req request.Report) error {
	var err error
	session := r.conn.Begin()
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

func (r mysqlRepository) ReportByRoomId(roomId string) (*[]model.ReportJoinEngage, error) {
	var report []model.ReportJoinEngage
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
	err := r.conn.Raw(sql).Scan(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r mysqlRepository) ReportListForCustomer(customerId string) (*[]model.ReportJoinEngage, error) {
	var reports []model.ReportJoinEngage
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
	err := r.conn.Raw(sql).Scan(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}
