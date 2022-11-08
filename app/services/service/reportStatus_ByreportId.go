package service

import "tpk-backend/app/models/model"

func (s serviceTPK) ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error) {
	reportStatus, err := s.repo.ReportStatusByReportId(reportId)
	if err != nil {
		return nil, err
	}
	return reportStatus, nil
}
