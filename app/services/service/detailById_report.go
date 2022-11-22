package service

import "tpk-backend/app/models/response"

func (s serviceTPK) ReportDetailById(reportId int) (*response.ReportDetailById, error) {
	report, err := s.repo.ReportDetailById(reportId)
	if err != nil {
		return nil, err
	}
	return report, nil
}
