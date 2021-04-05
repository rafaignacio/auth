package userInfoTests

import (
	"testing"

	"rafaignacio.com/auth/src/pkg/userInfo"
)

type UserInfoTest struct{}

func (u UserInfoTest) Write(userInfo userInfo.UserInfo) error {
	return nil
}

func TestNewUser(t *testing.T) {
	_, err := userInfo.NewUserInfo("email", "teste@a.com", "teste", UserInfoTest{})

	if err != nil {
		t.Fatalf("error %v", err.Error())
	}
}

func TestPassComparison(t *testing.T) {
	u := userInfo.UserInfo{}
	u.WritePassword("teste")

	if err := u.ComparePassword("teste"); err != nil {
		t.Fatalf("error comparing passwords: %v", err)
	}
}
