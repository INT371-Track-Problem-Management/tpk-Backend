package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CreateReport(req request.ReportInsert) (*int, error) {
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
	}
	session := s.database.Begin()
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
	return reportId, nil
}
