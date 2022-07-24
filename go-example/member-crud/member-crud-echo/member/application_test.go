package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMembership(t *testing.T) {
	t.Run("CreateMember", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("Fail when Duplicate MemberName", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		app.Create(req)

		existedNameReq := CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		}

		_, err := app.Create(existedNameReq)
		if assert.Error(t, err) {
			assert.Equal(t, ErrUserAlreadyExists, err)
		}
	})
	t.Run("Fail When MemberName is nil", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "",
			MembershipType: "naver",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, ErrUserNameIsRequired, err)
		}
	})

	t.Run("Fail When MembershipType is nil", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "jenny",
			MembershipType: "",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, ErrMembershipTypeIsRequired, err)
		}
	})

	t.Run("Fail When MembershipType is not naver/toss/payco", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "jenny",
			MembershipType: "kakao",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, ErrInvalidMembershipType, err)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Membership Update", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		res, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "payco",
		})

		req := UpdateRequest{
			ID:             res.ID,
			UserName:       "jenny",
			MembershipType: "naver",
		}

		_, err = app.Update(req)
		membershipFromData, _ := app.repository.data[res.ID]

		assert.Equal(t, "naver", membershipFromData.MembershipType)
		assert.Nil(t, err)
	})

	t.Run("Duplicate Error When Exist Name", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "payco",
		})

		_, err := app.Update(UpdateRequest{
			ID:             "update-1",
			UserName:       "jenny",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, ErrUserAlreadyExists, err)
		}
	})

	t.Run("Exception When Membership ID is None", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "",
			UserName:       "jenny",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, ErrUserIDIsRequired, err)
		}
	})

	t.Run("Exception When Membership UserName is None", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "update-2",
			UserName:       "",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, ErrUserNameIsRequired, err)
		}
	})

	t.Run("Exception MembershipType is None", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "update-3",
			UserName:       "jenny",
			MembershipType: "",
		})

		if assert.Error(t, err) {
			assert.Equal(t, ErrMembershipTypeIsRequired, err)
		}
	})

	t.Run("Exception MembershipType is Not naver/payco/toss ", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "update-4",
			UserName:       "jenny",
			MembershipType: "kakao",
		})

		if assert.Error(t, err) {
			assert.Equal(t, ErrInvalidMembershipType, err)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete Membership", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		res, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})
		assert.Nil(t, err)

		deleteReq := DeleteRequest{ID: res.ID}

		deleteRes, _ := app.Delete(deleteReq)
		assert.Equal(t, res.ID, deleteRes.ID)

	})

	t.Run("Exception When not write Id", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})
		assert.Nil(t, err)

		deleteReq := DeleteRequest{ID: ""}
		_, err = app.Delete(deleteReq)

		if assert.Error(t, err) {
			assert.Equal(t, ErrUserIDIsRequired, err)
		}

	})

	t.Run("Exception When Id is None", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})
		assert.Nil(t, err)

		req := DeleteRequest{ID: "uuid"}

		_, err = app.Delete(req)

		if assert.Error(t, err) {
			assert.Equal(t, ErrUserIDNotFound, err)
		}
	})
}
