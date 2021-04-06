package repos

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	utils "rafaignacio.com/auth/src/internal/utils"
	"rafaignacio.com/auth/src/pkg/apis/models"
	"rafaignacio.com/auth/src/pkg/userInfo"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() (*UserRepository, error) {

	cfg, err := utils.LoadConfig("../configs")

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.DBServer))
	db.AutoMigrate(&models.User{})

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db,
	}, nil
}

func (u UserRepository) Write(info userInfo.UserInfo) error {
	if u.db == nil {
		return errors.New("database is not connected")
	}

	var user models.User
	result := u.db.Where("id = ? OR provider_value = ?", info.ID.String(), info.Provider.Value).Find(&user)

	if result.RowsAffected > 0 {
		return errors.New("user already exists in our database")
	}

	user.NewUserModel = new(models.NewUserModel)

	user.ID = info.ID.String()
	user.ProviderType = string(info.Provider.Type)
	user.ProviderValue = info.Provider.Value
	user.Password = info.ReadPassword()

	result = u.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
