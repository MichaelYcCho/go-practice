package member

import (
	"errors"
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
			assert.Equal(t, errors.New("already existed user_name"), err)
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
			assert.Equal(t, errors.New("need user_name"), err)
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
			assert.Equal(t, errors.New("need membership type"), err)
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
			assert.Equal(t, errors.New("choose membership type : naver, payco, toss"), err)
		}
	})
}
