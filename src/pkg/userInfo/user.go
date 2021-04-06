package userInfo

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type ProviderType string

const (
	EMAIL        ProviderType = "email"
	SOCIAL_MEDIA ProviderType = "social_media"
)

type UserProvider struct {
	Type  ProviderType `json:"type"`
	Value string       `json:"value"`
}

type UserInfo struct {
	ID       UserID       `json:"id"`
	Provider UserProvider `json:"provider"`
	password Password
}

type UserInfoWriter interface {
	Write(info UserInfo) error
}

type UserInfoReader interface {
	Read(id string) UserInfo
}

func NewUserInfo(providerType ProviderType, providerValue, password string, writer UserInfoWriter) (UserInfo, error) {

	if writer == nil {
		return UserInfo{}, errors.New("writer is not defined")
	}

	userID, err := NewUserID()

	if err != nil {
		return UserInfo{}, err
	}

	provider, err := setUserProvider(providerType, providerValue)

	if err != nil {
		return UserInfo{}, err
	}

	info := UserInfo{
		ID:       userID,
		Provider: provider,
	}

	info.WritePassword(password)

	err = writer.Write(info)

	if err != nil {
		return UserInfo{}, err
	}

	return info, nil
}

func CreateUserInfo(id string, provider UserProvider, password []byte) UserInfo {
	output := UserInfo{}

	output.WriteEncryptedPassword(password)
	output.ID = createUserID(id)
	output.Provider = provider

	return output
}

func (u UserInfo) ReadPassword() []byte {
	return u.password.encryptedPass
}

func (u *UserInfo) WritePassword(password string) error {

	c, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	u.password.encryptedPass = c

	return nil
}

func (u *UserInfo) WriteEncryptedPassword(enc []byte) error {
	if enc != nil && len(enc) == 0 {
		return errors.New("password cannot be empty")
	}

	u.password.encryptedPass = enc

	return nil
}

func (u UserInfo) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.password.encryptedPass, []byte(password))
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
