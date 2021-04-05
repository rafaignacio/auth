package userInfo

type Password struct {
	encryptedPass []byte
}
type PasswordWriter interface {
	WritePassword(string) error
}

func (p Password) HasValue() bool {
	return p.encryptedPass != nil && len(p.encryptedPass) > 0
}
