package service

func (s serviceTPK) DeleteCustomer(id int) error {
	if err := s.repo.DeleteCustomer(id); err != nil {
		return err
	}
	return nil
}
