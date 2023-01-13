package repo

import (
	"github.com/DDGRCF/DDGBlog/database"
	"github.com/DDGRCF/DDGBlog/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserList() *[]model.User
	GetUserByEmail(email string) (user model.User, err error)
	GetUserByName(name string) *[]model.User

	SaveUser(user model.User) (err error)
	DelUserByEmail(email string) (err error)
}

type userRepository struct {
	ReDB *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		ReDB: database.GetReDB(),
	}
}

func (c *userRepository) GetUserList() *[]model.User {
	user := new([]model.User)
	err := c.ReDB.Raw(`select * FROM user`).Scan(user).RowsAffected

	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (c *userRepository) GetUserByEmail(email string) (user model.User, err error) {
	err = c.ReDB.First(&user, "email = ?", email).Error
	return user, err
}

func (c *userRepository) GetUserByName(name string) *[]model.User {
	user := new([]model.User)
	err := c.ReDB.Raw(`select * FROM user when user.name = ?`, name).Scan(user).RowsAffected
	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (c *userRepository) SaveUser(user model.User) (err error) {
	err = c.ReDB.Create(&user).Error
	return err
}

func (c *userRepository) DelUserByEmail(email string) (err error) {
	user := new(model.User)
	user.Email = email
	err = c.ReDB.Unscoped().Delete(&user).Error
	return err
}
