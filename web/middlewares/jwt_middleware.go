package middlewares

import (
	"log"
	"time"

	"github.com/DDGRCF/DDGBlog/configure"
	"github.com/DDGRCF/DDGBlog/models"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

const JwtKey = "My_Blog"

func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwtmiddleware.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},

		SigningMethod: jwtmiddleware.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err error) {
			if err != nil {
				log.Println(err)
			}
			ctx.JSON(iris.Map{
				"msg":  configure.JWT_AUTHORIZED_FAILURE_STR,
				"code": configure.JWT_AUTHORIZED_FAILURE_CODE,
			})
		},
	})
	return jwtHandler
}

func GenerateToken(ctx iris.Context, user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    user.Name,                                                //用户信息
		"session": user.Email + "_" + user.Name,                             //session
		"email":   user.Email,                                               //用户信息
		"iss":     configure.CommonConfig.GetString("appName"),              //签发者
		"iat":     time.Now().Unix(),                                        //签发时间
		"jti":     "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp":     time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})

	tokenString, _ := token.SignedString([]byte(JwtKey))
	ctx.Application().Logger().Debugf("Issue the Token, Message: %v", tokenString)
	return tokenString
}

func ParseToken(tokenStr string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims = token.Claims.(jwt.MapClaims)
	return
}

func DefaultJwtAuth(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"msg":  configure.JWT_AUTHORIZED_SUCCESS_STR,
		"code": configure.JWT_AUTHORIZED_SUCCESS_CODE,
	})

}
