package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"vulnsite/Response"
	"vulnsite/jwt"
	"vulnsite/model"
)

//认证中间件
func JwtAuth(isRoot bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.GetHeader("Authorization")
		checkToken := strings.SplitN(tokenHerder, " ", 2)

		//判断token格式是否正确
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			Response.PermissionDenied(c)
			c.Abort()
			return
		}

		//解析token，并判断token是否合法
		token, claims, err := jwt.ParseToken(checkToken[1])
		if err != nil || !token.Valid {
			fmt.Println(err)
			Response.PermissionDenied(c)
			c.Abort()
			return
		}

		//解析token成功后判断数据库中是否存在该用户
		user, err := model.GetAdminById(claims.ID)
		if err != nil || user.ID == 0 || user.Grade != claims.Grade{
			Response.PermissionDenied(c)
			c.Abort()
			return
		}

		//超级管理员认证
		if isRoot && claims.Grade != 1 {
			Response.PermissionDenied(c)
			c.Abort()
			return
		}
		//设置认证信息
		c.Set("AuthInfo", *claims.GetUserInfo())
	}
}


