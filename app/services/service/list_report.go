package service

import "tpk-backend/app/models/response"

func (s serviceTPK) ListReport() (*[]response.ReportList, error) {
	reports, err := s.repo.ListReport()
	if err != nil {
		return nil, err
	}
	return reports, nil
}
