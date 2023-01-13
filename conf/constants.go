package conf

// STR
const SYSTEM_ERROR_STR = "System Error"
const SUCCESS_CODE_STR = "Success"
const FAILURE_CODE_STR = "Failure"


const DB_SUCCESS_STR = "DB Search Success"
const DB_FAILURE_STR = "DB Search Failure"
const USER_LOGIN_SUCCESS_STR = "User Login Success"
const USER_LOGIN_FAILURE_STR = "User Login Failure"
const JWT_AUTHORIZED_FAILURE_STR = "JWT Authorized Failure"
const JWT_AUTHORIZED_SUCCESS_STR = "JWT Authorized Success"

// CODE
const SYSTEM_ERROR_CODE = 001
const SUCCESS_CODE = 002
const FAILURE_CODE = 003

const DB_SUCCESS_CODE = 101 // 数据库匹配成功
const DB_FAILURE_CODE = 102
const DB_BAD_COMMAND_CODE = 103 // 查找数据库的命令行错误
const DB_NOT_FIND_CODE = 104    // 没有找到相关数据库

const CHECK_LOGIN_SUCCESS_CODE = 200 // 登录成功
const CHECK_BAD_EMAIL_FROMAT_CODE = 201 // 错误邮箱
const CHECK_BAD_PASSWORD_FROMAT_CODE = 202 // 错误密码
const CHECK_ERROR_EMAIL_OR_PASSWORD_CODE = 203 // 密码或者邮箱错误
const JWT_AUTHORIZED_SUCCESS_CODE = 204 // JWT校验成功
const JWT_AUTHORIZED_FAILURE_CODE = 205 // JWT校验失败
