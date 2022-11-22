package service

import "tpk-backend/app/pkg"

func (s serviceTPK) RemoveCustomerFromRoom(id int) error {
	now := pkg.GetDatetime()
	if err := s.repo.RemoveCustomerFromRoom(id, now); err != nil {
		return err
	}
	return nil
}
