package midware

import (
	"fmt"
	"regexp"

	"github.com/DDGRCF/DDGBlog/conf"
	"github.com/DDGRCF/DDGBlog/service"
	"github.com/kataras/iris/v12"
)

func UserLoginCheck(ctx iris.Context) {
	emailMatchPattern := "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	passwordMatchPattern := "^[A-Za-z0-9@]{9,16}$"
	// 表单发送数据不是这样的{"email": ...., "password": ....}; 而是 email=....&password=...
	emailAddr := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if emailOk, _ := regexp.MatchString(emailMatchPattern, emailAddr); emailOk {
		if passwordOk, _ := regexp.MatchString(passwordMatchPattern, password); passwordOk {
			userService := service.NewUserService()
			userSearchResult := userService.GetUserByEmail(emailAddr)
			if userSearchResult.Code != conf.DB_SUCCESS_CODE {
				ctx.JSON(iris.Map{
					"code": conf.DB_FAILURE_CODE,
					"msg":  conf.DB_FAILURE_STR,
				})
			} else {
				userModel, err := userSearchResult.GetUserModel()
				if err != nil {
					ctx.JSON(iris.Map{
						"code": conf.SYSTEM_ERROR_CODE,
						"msg": conf.SYSTEM_ERROR_STR,
					})
				} else {

				if userModel.Email == emailAddr && userModel.Password == password {
					jwtString := GenerateToken(ctx, userModel)
					// 这里在Header上Token传递jwtString
					ctx.Request().Header.Set("Token", jwtString)
					ctx.Next()
				} else {
					ctx.JSON(iris.Map{
						"code": conf.CHECK_ERROR_EMAIL_OR_PASSWORD_CODE,
						"msg":  conf.USER_LOGIN_FAILURE_STR,
					})
				}
				}
			}
		} else {
			ctx.Application().Logger().Debugf(fmt.Sprintf("[Login] email: %v, password: %v fail!", emailAddr, password))
			ctx.JSON(iris.Map{
				"code": conf.CHECK_BAD_PASSWORD_FROMAT_CODE,
				"msg":  conf.USER_LOGIN_FAILURE_STR,
			})
		}
	} else {
		ctx.Application().Logger().Debugf(fmt.Sprintf("[Login] email: %v, password: %v login fail!", emailAddr, password))
		ctx.JSON(iris.Map{
			"code": conf.CHECK_BAD_EMAIL_FROMAT_CODE,
			"msg":  conf.USER_LOGIN_FAILURE_STR,
		})
	}
}
