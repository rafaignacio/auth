package models

import (
	"rafaignacio.com/auth/src/pkg/userInfo"
)

type NewUserModel struct {
	Password      string `json:"password"`
	ProviderType  string `json:"provider_type"`
	ProviderValue string `json:"provider_value"`
}

func (u NewUserModel) Validate() (errs []string) {

	if u.Password == "" {
		errs = append(errs, "password cannot be empty")
	}

	if userInfo.ProviderType(u.ProviderType) == "" {
		errs = append(errs, "provider cannot be empyt")
	}

	if u.ProviderValue == "" {
		errs = append(errs, "provider value cannot be empty")
	}

	return
}
