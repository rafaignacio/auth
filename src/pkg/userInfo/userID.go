package userInfo

import "github.com/google/uuid"

type UserID struct {
	id string
}

func (u UserID) HasValue() bool {
	return u.id != ""
}

func NewUserID() (UserID, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return UserID{}, err
	}

	output := UserID{
		id: id.String(),
	}

	return output, nil
}
