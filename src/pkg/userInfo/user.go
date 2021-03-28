package userInfo

import (
	"errors"
	"regexp"
)

type ProviderType string

const (
	EMAIL        ProviderType = "email"
	SOCIAL_MEDIA ProviderType = "social_media"
)

type UserProvider struct {
	Type  ProviderType
	Value string
}

type UserInfo struct {
	ID       UserID
	Provider UserProvider
	Password Password
}

type UserInfoWriter interface {
	Write(info UserInfo) error
}

type UserInfoReader interface {
	Read(id string) UserInfo
}

func NewUserInfo(providerType ProviderType, providerValue, password string, writer UserInfoWriter) error {

	if writer == nil {
		return errors.New("writer is not defined")
	}

	userID, err := NewUserID()

	if err != nil {
		return err
	}

	provider, err := setUserProvider(providerType, providerValue)

	if err != nil {
		return err
	}

	info := UserInfo{
		ID:       userID,
		Provider: provider,
	}

	info.Password.WritePassword(password)

	return writer.Write(info)
}

func setUserProvider(providerType ProviderType, value string) (UserProvider, error) {
	output := UserProvider{}

	switch providerType {
	case EMAIL:
		if err := validateEmail(value); err != nil {
			return UserProvider{}, err
		}
	case SOCIAL_MEDIA:
		break
	default:
		return UserProvider{}, errors.New("invalid provider type")
	}

	output.Type = providerType
	output.Value = value

	return output, nil
}

func validateEmail(email string) error {
	r, err := regexp.Compile(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`)

	if err != nil {
		return err
	}

	if !r.MatchString(email) {
		return errors.New("invalid e-mail")
	}

	return nil
}
