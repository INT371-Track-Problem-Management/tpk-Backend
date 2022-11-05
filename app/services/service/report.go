package service

import (
	"tpk-backend/app/constants"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) Report() (*[]model.ReportJoinEngage, error) {
	report, err := s.repo.Report()
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s serviceTPK) ReportById(req request.Report) (*response.Report, error) {
	data, err := s.repo.ReportById(req.ReportId)
	if err != nil {
		return nil, err
	}
	res := &response.Report{
		ReportId:         data.ReportId,
		Title:            data.Title,
		CategoriesReport: data.CategoriesReport,
		ReportDes:        data.ReportDes,
		Status:           data.Status,
		UpdateAt:         data.UpdateAt,
		UpdateBy:         data.UpdateBy,
		CreateAt:         data.CreateAt,
		CreateBy:         data.CreateBy,
		RoomId:           data.RoomId,
		RoomNum:          data.RoomNum,
		BuildingId:       data.BuildingId,
		SelectedDate:     data.SelectedDate,
	}
	return res, nil
}

func (s serviceTPK) ReportByCreatedBy(customerId string) (*[]model.Report, error) {
	res, err := s.repo.ReportByCreatedBy(customerId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s serviceTPK) ReportInsert(req request.ReportInsert) (*int, error) {
	timenow := pkg.GetDatetime()
	data := model.ReportInsert{
		Title:            req.Title,
		CategoriesReport: req.CategoriesReport,
		ReportDes:        req.ReportDes,
		Status:           req.Status,
		CreateAt:         timenow,
		CreateBy:         req.CreateBy,
		UpdateAt:         timenow,
		UpdateBy:         req.CreateBy,
		RoomId:           req.RoomId,
	}

	reportid, customer, err := s.repo.ReportInsert(data)
	if err != nil {
		return nil, err
	}
	err = pkg.Smtp2(constants.SUBJECT_EMAIL_SENDING_REPORT, customer.Email, "ส่งการรายงาน")
	if err != nil {
		return nil, err
	}

	return reportid, nil
}
