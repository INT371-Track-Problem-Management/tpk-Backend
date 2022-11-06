package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) ReportEnagegeFixDateDetail(reportId string) (*response.ReportEngageFixDate, error) {
	engage, err := s.repo.ReportEnagegeByReportId(reportId)
	if err != nil {
		return nil, err
	}
	fixDate, err := s.repo.FixdateByEngageId(*engage)
	if err != nil {
		return nil, err
	}

	engageDetail := response.ReportEngageFixDate{
		EngageId:     engage.EngageId,
		Step:         engage.Step,
		SelectedDate: engage.SelectedDate,
		ReportId:     engage.ReportId,
		BuildingId:   engage.BuildingId,
		CreateBy:     engage.CreateBy,
		CreateAt:     engage.CreateAt,
		UpdateAt:     engage.UpdateAt,
		UpdateBy:     engage.UpdateBy,
		MaintainerId: engage.MaintainerId,
		Fixdate:      *fixDate,
	}

	return &engageDetail, nil
}

func (s serviceTPK) NewFixDate(req request.ReportEngage) error {
	now := pkg.GetDatetime()
	session := s.database.Begin()

	if err := s.repo.EditEngage(req, session); err != nil {
		return err
	}

	for _, v := range req.Dates {
		date := model.CreateFixdate{
			Date:     v.Date,
			Step:     req.Step,
			CreateAt: now,
			EngageId: req.EngageId,
		}
		if err := s.repo.CreateFixdate(date, session); err != nil {
			return err
		}
	}

	session.Commit()
	return nil
}
