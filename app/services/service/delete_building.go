package service

func (s serviceTPK) BuildingDelete(buildingId int) error {
	if err := s.repo.BuildingDelete(buildingId); err != nil {
		return err
	}
	return nil
}
