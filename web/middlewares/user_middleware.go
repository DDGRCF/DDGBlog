package middlewares

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/models"
	"github.com/DDGRCF/DDGBlog/services"
	"github.com/kataras/iris/v12"
)

func BeforePostFormCheckFormat(ctx iris.Context) {
	emailMatchPattern := "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	passwordMatchPattern := "^[A-Za-z0-9@]{9,16}$"
	emailAddr := ctx.PostValue("emailInput")
	password := ctx.PostValue("passwordInput")

	if emailOk, _ := regexp.MatchString(emailMatchPattern, emailAddr); emailOk {
		if passwordOk, _ := regexp.MatchString(passwordMatchPattern, password); passwordOk {
			userService := services.NewUserService()
			userSearchResult := userService.GetUserByEmail(emailAddr)
			if userSearchResult.Code != configure.DB_SUCCESS_CODE {
				ctx.JSON(iris.Map{
					"code": "400",
					"msg":  configure.CHECK_BAD_PASSWORD_FROMAT_STR,
				})
			} else {
				var userModel models.User
				_userModel := userSearchResult.Data
				userByte, _ := json.Marshal(_userModel)
				json.Unmarshal(userByte, &userModel)
				if userModel.Email == emailAddr && userModel.Password == password {
					ctx.Application().Logger().Debugf("[Login][%v][%v] success!", userModel.Email, userModel.Password)
					GenerateToken(ctx, userModel)
					ctx.Next()
				} else {
					ctx.Application().Logger().Debugf("[Login][%v][%v] fail!", userModel.Email, userModel.Password)
					ctx.JSON(iris.Map{
						"code": "400",
						"msg":  configure.CHECK_ERROR_EMAIL_OR_PASSWORD_STR,
					})
				}
			}
		} else {
			ctx.Application().Logger().Debugf(fmt.Sprintf("[Login] email: %v, password: %v fail!", emailAddr, password))
			ctx.JSON(iris.Map{
				"code": "400",
				"msg":  configure.CHECK_BAD_PASSWORD_FROMAT_STR,
			})
		}
	} else {
		ctx.Application().Logger().Debugf(fmt.Sprintf("[Login] email: %v, password: %v login fail!", emailAddr, password))
		ctx.JSON(iris.Map{
			"code": "400",
			"msg":  configure.CHECK_BAD_EMAIL_FROMAT_STR,
		})
	}
}
