package repos

import (
	"github.com/DDGRCF/DDGBlog/database"
	"github.com/DDGRCF/DDGBlog/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserList() *[]models.User
	GetUserByEmail(email string) (user models.User, err error)
	GetUserByName(name string) *[]models.User

	SaveUser(user models.User) (err error)
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

func (c *userRepository) GetUserList() *[]models.User {
	user := new([]models.User)
	err := c.ReDB.Raw(`select * FROM user`).Scan(user).RowsAffected

	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (c *userRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = c.ReDB.First(&user, "email = ?", email).Error
	return user, err
}

func (c *userRepository) GetUserByName(name string) *[]models.User {
	user := new([]models.User)
	err := c.ReDB.Raw(`select * FROM user when user.name = ?`, name).Scan(user).RowsAffected
	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (c *userRepository) SaveUser(user models.User) (err error) {
	err = c.ReDB.Create(&user).Error
	return err
}

func (c *userRepository) DelUserByEmail(email string) (err error) {
	user := new(models.User)
	user.Email = email
	err = c.ReDB.Unscoped().Delete(&user).Error
	return err
}
