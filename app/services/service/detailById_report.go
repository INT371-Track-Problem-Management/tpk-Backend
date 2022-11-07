package service

import "tpk-backend/app/models/model"

func (s serviceTPK) ReportDetailById(reportId int) (*model.Report, error) {
	report, err := s.repo.ReportDetailById(reportId)
	if err != nil {
		return nil, err
	}
	return report, nil
}
