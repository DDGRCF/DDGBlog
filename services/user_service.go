package services

import (
	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/models"
	"github.com/DDGRCF/DDGBlog/repos"
)

type UserService interface {
	GetUserList() *models.Result
	GetUserByEmail(email string) (result models.Result)
	DelUserByEmail(email string) (result models.Result)

	PostSaveUser(user models.User) (result models.Result)
}

type userService struct {
	UserRepo repos.UserRepository
}

func NewUserService() UserService {
	return &userService{
		UserRepo: repos.NewUserRepository(),
	}
}

func (c *userService) GetUserList() *models.Result {
	userList := c.UserRepo.GetUserList()
	result := new(models.Result)
	result.Data = userList
	result.Msg = configure.DB_SUCCESS_STR
	result.Code = configure.DB_SUCCESS_CODE
	return result
}

func (c *userService) GetUserByEmail(email string) (result models.Result) {
	user, err := c.UserRepo.GetUserByEmail(email)
	if err != nil {
		result.Code = configure.DB_BAD_COMMAND_CODE
	} else {
		result.Data = user
		result.Code = configure.DB_SUCCESS_CODE
		result.Msg = configure.DB_SUCCESS_STR
	}

	return result
}

func (c *userService) DelUserByEmail(email string) (result models.Result) {
	err := c.UserRepo.DelUserByEmail(email)
	if err != nil {
		result.Code = configure.DB_BAD_COMMAND_CODE
	} else {
		result.Code = configure.DB_SUCCESS_CODE
		result.Msg = configure.DB_SUCCESS_STR
		list := c.UserRepo.GetUserList()
		result.Data = list
	}
	return result
}

func (c *userService) PostSaveUser(user models.User) (result models.Result) {
	err := c.UserRepo.SaveUser(user)
	if err != nil {
		result.Code = configure.DB_BAD_COMMAND_CODE
		result.Msg = configure.DB_FAILURE_STR
	} else {
		if user, err = c.UserRepo.GetUserByEmail(user.Email); err != nil {
			result.Code = configure.DB_NOT_FIND_CODE
			result.Msg = configure.DB_FAILURE_STR
		} else {
			result.Code = configure.DB_SUCCESS_CODE
			result.Msg = configure.DB_SUCCESS_STR
		}
	}
	return result
}
