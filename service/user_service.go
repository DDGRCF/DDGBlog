package service

import (
	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/model"
	"github.com/DDGRCF/DDGBlog/repo"
)

type UserService interface {
	GetUserList() *model.Result
	GetUserByEmail(email string) (result model.Result)
	DelUserByEmail(email string) (result model.Result)

	PostSaveUser(user model.User) (result model.Result)
}

type userService struct {
	UserRepo repo.UserRepository
}

func NewUserService() UserService {
	return &userService{
		UserRepo: repo.NewUserRepository(),
	}
}

func (c *userService) GetUserList() *model.Result {
	userList := c.UserRepo.GetUserList()
	result := new(model.Result)
	result.Data = userList
	result.Msg = conf.DB_SUCCESS_STR
	result.Code = conf.DB_SUCCESS_CODE
	return result
}

func (c *userService) GetUserByEmail(email string) (result model.Result) {
	user, err := c.UserRepo.GetUserByEmail(email)
	if err != nil {
		result.Code = conf.DB_BAD_COMMAND_CODE
	} else {
		result.Data = user
		result.Code = conf.DB_SUCCESS_CODE
		result.Msg = conf.DB_SUCCESS_STR
	}

	return result
}

func (c *userService) DelUserByEmail(email string) (result model.Result) {
	err := c.UserRepo.DelUserByEmail(email)
	if err != nil {
		result.Code = conf.DB_BAD_COMMAND_CODE
	} else {
		result.Code = conf.DB_SUCCESS_CODE
		result.Msg = conf.DB_SUCCESS_STR
		list := c.UserRepo.GetUserList()
		result.Data = list
	}
	return result
}

func (c *userService) PostSaveUser(user model.User) (result model.Result) {
	err := c.UserRepo.SaveUser(user)
	if err != nil {
		result.Code = conf.DB_BAD_COMMAND_CODE
		result.Msg = conf.DB_FAILURE_STR
	} else {
		if user, err = c.UserRepo.GetUserByEmail(user.Email); err != nil {
			result.Code = conf.DB_NOT_FIND_CODE
			result.Msg = conf.DB_FAILURE_STR
		} else {
			result.Code = conf.DB_SUCCESS_CODE
			result.Msg = conf.DB_SUCCESS_STR
		}
	}
	return result
}
