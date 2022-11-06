package service

import (
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

func (s serviceTPK) ListReport(fillter *request.FillterReport) (*[]response.ReportList, error) {
	reports, err := s.repo.ListReport(fillter)
	if err != nil {
		return nil, err
	}
	return reports, nil
}
