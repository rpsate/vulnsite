package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	SUCCESS = 200
	FAIL = 400
	PERMISSION_DENIED = 401
	USERNAME_PASSWORD_ERR = 402
	LOGIN_SUCCESS = 403
	PARAMETER_ERROR = 404
	USERNAME_EXIST = 405
	USERNAME_NOT_EXIST = 406
	RECORD_NOT_EXIST = 407
	SYSTEM_ERR = 422
)

var MSG = map[int]string {
	SUCCESS: "成功",
	FAIL: "失败",
	PERMISSION_DENIED: "权限不足",
	USERNAME_PASSWORD_ERR: "账号或密码错误",
	LOGIN_SUCCESS: "登陆成功",
	PARAMETER_ERROR: "参数错误",
	USERNAME_EXIST: "用户已经存在",
	USERNAME_NOT_EXIST: "用户不存在",
	RECORD_NOT_EXIST: "该记录不存在",
	SYSTEM_ERR: "系统错误",
}

func Response(c *gin.Context, code int, msg string, data gin.H)  {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg": MSG[SUCCESS],
		"data": nil,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAIL,
		"msg": MSG[FAIL],
		"data": nil,
	})
}

func ParameterError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": PARAMETER_ERROR,
		"msg": MSG[PARAMETER_ERROR],
		"data": nil,
	})
}

func SystemError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SYSTEM_ERR,
		"msg": MSG[SYSTEM_ERR],
		"data": nil,
	})
}

func PermissionDenied(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": PERMISSION_DENIED,
		"msg": MSG[PERMISSION_DENIED],
		"data": nil,
	})
}