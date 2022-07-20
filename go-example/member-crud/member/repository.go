package member

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) CreateMembership(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.UserName == m.UserName {
			return Membership{}, ErrUserAlreadyExists
		}
	}
	r.data[m.ID] = m
	return m, nil
}

func (r *Repository) UpdateMembership(m Membership) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == m.ID {
			continue
		}
		if membership.UserName == m.UserName {
			return Membership{}, ErrUserAlreadyExists
		}
	}
	r.data[m.ID] = m
	return m, nil
}
