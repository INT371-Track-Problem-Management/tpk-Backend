package service

import (
	"fmt"
	"mime/multipart"
	"tpk-backend/app/constants"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CreateReport(req request.ReportInsert, image *multipart.FileHeader) (*int, error) {
	var err error
	var file *model.ReportMedia
	if image != nil {
		file, err = pkg.UploadFile(image, constants.IMAGE_DES_REPORT)
		if err != nil {
			return nil, err
		}
	}
	now := pkg.GetDatetime()
	report := model.ReportInsert{
		Title:            req.Title,
		CategoriesReport: req.CategoriesReport,
		ReportDes:        req.ReportDes,
		Status:           req.Status,
		RoomId:           req.RoomId,
		UpdateAt:         now,
		UpdateBy:         req.UpdateBy,
		CreateAt:         now,
		CreateBy:         req.UpdateBy,
		ImageId:          file.Id,
	}
	session := s.database.Begin()
	if err := s.repo.CreateReportMedia(*file, session); err != nil {
		return nil, err
	}
	reportId, err := s.repo.CreateReport(report, session)
	if err != nil {
		return nil, err
	}

	reportStatus := request.ReportStatus{
		ReportId:  *reportId,
		Status:    req.Status,
		UpdateAt:  now,
		UpdateBy:  req.UpdateBy,
		CreatedAt: now,
	}
	if err := s.repo.CreateReportStatus(reportStatus, session); err != nil {
		return nil, err
	}

	engage := model.InsertReportEngage{
		Step:       req.Step,
		ReportId:   *reportId,
		BuildingId: req.BuildingId,
		CreateBy:   req.UpdateBy,
		CreateAt:   now,
		UpdateAt:   now,
		UpdateBy:   req.UpdateBy,
	}
	engageId, err := s.repo.CreateReporEngage(engage, session)
	if err != nil {
		return nil, err
	}

	for _, v := range req.Dates {
		date := model.CreateFixdate{
			Date:     v.Date,
			Step:     req.Step,
			CreateAt: now,
			EngageId: *engageId,
		}
		if err := s.repo.CreateFixdate(date, session); err != nil {
			return nil, err
		}
	}
	session.Commit()

	customer, err := s.repo.GetCustomerById(req.UpdateBy)
	if err != nil {
		return nil, err
	}

	body := fmt.Sprintf("แจ้งเตือนการรายงานปัญหา รหัส: %v กรุณารอเจ้าหน้าที่ตรวจสอบ", reportId)
	if err := pkg.Smtp2(constants.SUBJECT_EMAIL_SENDING_REPORT, customer.Email, body); err != nil {
		return nil, err
	}

	return reportId, nil
}
