package service

import "tpk-backend/app/models/request"

func (s serviceTPK) ChangeCategory(req request.ReportChangeCategory) error {
	if err := s.repo.ChangeCategory(req); err != nil {
		return err
	}
	return nil
}
