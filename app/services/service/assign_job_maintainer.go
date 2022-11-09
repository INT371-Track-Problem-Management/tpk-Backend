package service

import "tpk-backend/app/models/request"

func (s serviceTPK) AssignJobMaintainer(req request.AssignReport) error {
	if err := s.repo.AssignJobMaintainer(req); err != nil {
		return err
	}
	return nil
}
