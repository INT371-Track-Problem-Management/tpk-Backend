package service

func (s serviceTPK) DeleteEmployee(id int) error {
	if err := s.repo.DeleteEmployee(id); err != nil {
		return err
	}
	return nil
}
