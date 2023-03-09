package routers

import (
	"github.com/gin-gonic/gin"
	"vulnsite/controller"
	"vulnsite/middleware"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	loginGroup := r.Group("/v1/auth")
	{
		//登陆
		loginGroup.POST("/login", controller.Login)
	}

	adminGroup := r.Group("/v1")
	adminGroup.Use(middleware.JwtAuth(true))
	{
		//查询所有管理员信息
		adminGroup.GET("/Admin", controller.FindAllAdmin)
		//删除一个管理员
		adminGroup.DELETE("/Admin/:id", controller.DeteleAdmin)
		//添加一个管理员
		adminGroup.POST("/Admin", controller.CreateAdmin)
		//修改管理员信息
		adminGroup.PUT("/Admin", controller.UpdateAdmin)
	}

	userGroup := r.Group("/v1")
	userGroup.Use(middleware.JwtAuth(false))
	{
		//获取我的信息
		userGroup.GET("/Admin/My", controller.FindMyAdmin)
		//修改我的信息
		userGroup.PUT("/Admin/My", controller.UpdateMyAdmin)
	}

	machineGroup := r.Group("/v1/Machine")
	machineGroup.Use(middleware.JwtAuth(false))
	{
		//添加虚拟机
		machineGroup.POST("", controller.CreateMachine)
		//删除虚拟机
		machineGroup.DELETE("/:id", controller.DeleteMachine)
		//修改虚拟机
		machineGroup.PUT("", controller.UpdateMachine)
	}

	findMachineGroup := r.Group("/v1/Machine")
	{
		//查询虚拟机
		findMachineGroup.GET("/:pageNum/:pageSize", controller.FindMachine)
		//按难度查询虚拟机
		findMachineGroup.GET("/:pageNum/:pageSize/:difficulty", controller.FindMachineByDiff)
		//搜索虚拟机
		findMachineGroup.GET("/search/:pageNum/:pageSize/:keys", controller.SearchMachine)
	}

	return r
}
