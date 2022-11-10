package service

func (s serviceTPK) LogoutToken(token string) error {
	if err := s.repo.LogoutToken(token); err != nil {
		return err
	}
	return nil
}
