package repos

import "rafaignacio.com/auth/src/pkg/userInfo"

type UserRepository struct{}

func (u UserRepository) Write(info userInfo.UserInfo) error {
	return nil
}
