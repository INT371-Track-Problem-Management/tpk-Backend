package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) EndJobReport(req request.EndJobReport) error {
	now := pkg.GetDatetime()
	entity := model.EndJobReport{
		Des:         req.Des,
		ReportId:    req.ReportId,
		Score:       req.Score,
		DateOfIssue: now,
		UpdateBy:    req.UpdateBy,
	}
	status := request.ReportStatus{
		ReportId:  req.ReportId,
		Status:    "S7",
		UpdateAt:  now,
		UpdateBy:  req.UpdateBy,
		CreatedAt: now,
	}

	session := s.database.Begin()
	if err := s.repo.EndJobReport(session, entity); err != nil {
		return err
	}
	if err := s.repo.CreateReportStatus(status, session); err != nil {
		return err
	}
	session.Commit()
	return nil
}
