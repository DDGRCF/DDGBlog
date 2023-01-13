package model

type Response struct {
	Msg  string 	 `json:"msg"`
	Code int64 		 `json:"code"` 
	Data interface{} `json:"data"`
}