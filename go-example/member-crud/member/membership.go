package member

import "errors"

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

type MembershipGenerator struct {
	Membership *Membership
}

func NewMembershipGenerator() *MembershipGenerator {
	return &MembershipGenerator{Membership: &Membership{}}
}

func (m *MembershipGenerator) SetID(id string) *MembershipGenerator {
	m.Membership.ID = id
	return m
}

func (m *MembershipGenerator) SetUserName(userName string) *MembershipGenerator {
	m.Membership.UserName = userName
	return m
}

func (m *MembershipGenerator) SetMembershipType(membershipType string) *MembershipGenerator {
	m.Membership.MembershipType = membershipType
	return m
}

func (m *MembershipGenerator) GetMembership() (*Membership, error) {
	err := m.validateMembership()
	if err != nil {
		return nil, err
	}
	return m.Membership, nil
}

func (m *MembershipGenerator) validateMembership() error {
	if m.Membership.ID == "" {
		return errors.New("need id")
	}
	if m.Membership.UserName == "" {
		return errors.New("need user_name")
	}
	if m.Membership.MembershipType == "" {
		return errors.New("need membership type")
	}
	if !(m.Membership.MembershipType == "naver" || m.Membership.MembershipType == "payco" || m.Membership.MembershipType == "toss") {
		return errors.New("choose membership type : naver, payco, toss")
	}
	return nil
}