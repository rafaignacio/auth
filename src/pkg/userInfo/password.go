package userInfo

import "golang.org/x/crypto/bcrypt"

type Password struct {
	encryptedPass []byte
}

type PasswordWriter interface {
	WritePassword(string) error
}

func (p Password) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(p.encryptedPass, []byte(password))
}

func (p Password) HasValue() bool {
	return len(p.encryptedPass) > 0
}

func (p *Password) WritePassword(password string) error {

	c, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	p.encryptedPass = c

	return nil
}
