package member

func (service *Service) ValidateDelete(id string) error {
	if id == "" {
		return ErrUserIDIsRequired
	}
	if _, exist := service.repository.data[id]; !exist {
		return ErrUserIDNotFound
	}
	return nil
}
