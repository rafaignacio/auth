package userInfo

type Email struct {
	email string
}

type UserInfo struct {
	ID       UserID
	Email    Email
	Password Password
}

func NewUserInfo(email, password string) (UserInfo, error) {
	p := Password{}
	p.WritePassword(password)

	userID, err := NewUserID()

	if err != nil {
		return UserInfo{}, err
	}

	output := UserInfo{
		ID:       userID,
		Password: p,
	}

	return output, nil
}
