package userInfoTests

import (
	"testing"

	"rafaignacio.com/auth/src/pkg/userInfo"
)

func TestNewUser(t *testing.T) {
	user, err := userInfo.NewUserInfo("email", "teste@a.com", "teste")

	if err != nil {
		t.Fatalf("error %v", err.Error())
	}

	t.Logf("user password: %+v", user.Password)

	t.Logf("%+v", user)
}

func TestPassComparison(t *testing.T) {
	p := userInfo.Password{}
	p.WritePassword("teste")

	if err := p.ComparePassword("teste"); err != nil {
		t.Fatalf("error comparing passwords: %v", err)
	}
}
