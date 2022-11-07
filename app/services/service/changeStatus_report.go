package service

import (
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) ChangeStatusReport(req request.ReportChangeStatus) error {
	session := s.database.Begin()
	now := pkg.GetDatetime()
	status := request.ReportStatus{
		ReportId:  req.ReportId,
		Status:    req.Status,
		Detail:    req.Detail,
		UpdateAt:  now,
		UpdateBy:  req.UpdateBy,
		CreatedAt: now,
	}

	if err := s.repo.ChangeStatusReport(status, session); err != nil {
		return err
	}

	if err := s.repo.CreateReportStatus(status, session); err != nil {
		return err
	}
	session.Commit()
	return nil
}
