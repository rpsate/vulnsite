package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"vulnsite/Response"
	"vulnsite/jwt"
	"vulnsite/model"
)

//Login 登陆
func Login(c *gin.Context) {
	//获取参数
	var adminParams = new(model.Admin)
	c.ShouldBindJSON(&adminParams)
	username := adminParams.Username
	password := adminParams.Password
	if username == "" || password == "" {
		Response.ParameterError(c)
		return
	}
	//获取账号密码
	admin, err := model.GetAdminByUsername(username)
	if admin == nil || err != nil {
		Response.Response(
			c,
			Response.USERNAME_PASSWORD_ERR,
			Response.MSG[Response.USERNAME_PASSWORD_ERR],
			nil,
		)
		return
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		Response.Response(
			c,
			Response.USERNAME_PASSWORD_ERR,
			Response.MSG[Response.USERNAME_PASSWORD_ERR],
			nil,
		)
		return
	}
	//登陆成功，获取token
	token, err := jwt.SetToken(admin.ID, admin.Username, admin.Name, admin.Grade)
	if err != nil {
		Response.SystemError(c)
		return
	}

	//发放token
	Response.Response(
		c,
		Response.LOGIN_SUCCESS,
		Response.MSG[Response.LOGIN_SUCCESS],
		gin.H{"token": token},
		)
}

//获取用户列表
func FindAllAdmin(c *gin.Context) {
	adminInfos, err := model.GetAllAdmin()
	if err != nil {
		Response.Fail(c)
	}
	Response.Response(
		c,
		Response.SUCCESS,
		Response.MSG[Response.SUCCESS],
		gin.H{
			"AdminInfoList": adminInfos,
		},
		)
}

//删除用户
func DeteleAdmin(c *gin.Context)  {
	//获取参数
	sId := c.Param("id")
	id, err := strconv.Atoi(sId)
	//判断参数是否为整数
	if sId == "" || err != nil {
		Response.ParameterError(c)
		return
	}
	//删除用户
	if err = model.DeleteAdmin(id); err != nil {
		Response.SystemError(c)
		return
	}
	//删除成功
	Response.Success(c)
}

// CreateAdmin 添加用户
func CreateAdmin(c *gin.Context) {
	//获取参数
	admin := new(model.Admin)
	c.ShouldBindJSON(&admin)
	//判断参数是否为空
	if admin.Name =="" || admin.Username == "" || admin.Password == "" {
		Response.ParameterError(c)
		return
	}
	//判断是否用户存在
	count, err := model.GetCountAdminByUsername(admin.Username)
	if err != nil {
		Response.SystemError(c)
		return
	}
	if count > 0 {
		Response.Response(
			c,
			Response.USERNAME_EXIST,
			Response.MSG[Response.USERNAME_EXIST],
			nil,
			)
		return
	}
	//加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		Response.SystemError(c)
	}
	admin.Password = string(password)
	//添加用户
	err = model.CreateAdmin(admin)
	if err != nil {
		Response.Fail(c)
		return
	}
	Response.Success(c)
}

// UpdateAdmin 修改用户信息
func UpdateAdmin(c *gin.Context) {
	//获取参数
	admin := new(model.Admin)
	c.ShouldBindJSON(&admin)
	//判断id是否为空,是否尝试修改用户名
	if admin.ID == 0 || admin.Username != ""{
		Response.ParameterError(c)
		return
	}
	//判断用户是否存在
	count, err := model.GetCountAdminById(admin.ID)
	if err != nil {
		Response.SystemError(c)
	}
	if count == 0 {
		Response.Response(
			c,
			Response.USERNAME_NOT_EXIST,
			Response.MSG[Response.USERNAME_NOT_EXIST],
			nil,
			)
		return
	}
	//如果密码存在，则加密密码
	if admin.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			Response.SystemError(c)
			return
		}
		admin.Password = string(password)
	}
	//修改用户
	err = model.UpdateAdmin(admin)
	if err != nil {
		Response.SystemError(c)
		return
	}
	Response.Success(c)
}

// FindMyAdmin 查询自己信息
func FindMyAdmin(c *gin.Context) {
	//获取信息
	authInfo, isExist := c.Get("AuthInfo")
	if !isExist {
		Response.SystemError(c)
		return
	}
	adminInfo := authInfo.(model.AdminInfo)
	//返回用户信息
	Response.Response(
		c,
		Response.SUCCESS,
		Response.MSG[Response.SUCCESS],
		gin.H{
			"AdminInfo": adminInfo,
		})
}

// UpdateMyAdmin 修改自己信息
func UpdateMyAdmin(c *gin.Context) {
	//获取参数
	admin := new(model.Admin)
	c.ShouldBindJSON(&admin)
	//是否尝试修改用户名
	if admin.Username != "" {
		Response.ParameterError(c)
		return
	}
	//判断是否越权修改用户权限
	if admin.Grade != 0 {
		Response.PermissionDenied(c)
		return
	}
	//设置自己的ID
	authInfo, isExist := c.Get("AuthInfo")
	if !isExist {
		Response.PermissionDenied(c)
		return
	}
	AdminInfo := authInfo.(model.AdminInfo)
	admin.ID = AdminInfo.ID

	//如果密码存在，则加密密码
	if admin.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			Response.SystemError(c)
			return
		}
		admin.Password = string(password)
	}

	//修改用户
	err := model.UpdateAdmin(admin)
	if err != nil {
		Response.SystemError(c)
		return
	}
	Response.Success(c)
}