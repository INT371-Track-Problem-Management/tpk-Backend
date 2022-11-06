package service

import "tpk-backend/app/models/request"

func (s serviceTPK) SelectedPlanFixDate(req request.SelectedPlanFixDate) error {
	if err := s.repo.SelectedPlanFixDate(req); err != nil {
		return err
	}
	return nil
}
