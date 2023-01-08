package models

import (
	"encoding/json"
	"strings"

	"github.com/DDGRCF/DDGBlog/configure"
)

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func NewResult(data interface{}, c int, m ...string) *Result {
	r := &Result{Data: data, Code: c}

	if e, ok := data.(error); ok { // 接口断言
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = configure.DB_SUCCESS_STR
	}

	if len(m) > 0 {
		r.Msg = strings.Join(m, ", ")
	}

	return r
}

func (r *Result) GetUserModel() (user *User, err error) {
	userByte, err := json.Marshal(r.Data)
	if err != nil {
		return user, err
	} else {
		err = json.Unmarshal(userByte, &user)
	}
	return user, err
}
