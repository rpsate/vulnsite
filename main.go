package main

import (
	"golang.org/x/crypto/bcrypt"
	"vulnsite/dao"
	"vulnsite/model"
	"vulnsite/routers"
	"vulnsite/utils"
)

func main() {
	dao.InitDb()
	defer dao.DB.Close()
	//如果数据库不存在则创建数据库
	dao.DB.AutoMigrate(&model.Admin{})
	dao.DB.AutoMigrate(&model.Machine{})
	//如果没有用户则创建初管理员
	if utils.AutoCreateAdmin {
		adminUsername := utils.Admin
		count, err := model.GetCountAdminByUsername(adminUsername)
		if err == nil && count == 0 {
			adminPassword, _ := bcrypt.GenerateFromPassword([]byte(adminUsername), bcrypt.DefaultCost)
			var admin = new(model.Admin)
			admin.Username = adminUsername
			admin.Name = adminUsername
			admin.Grade = 1
			admin.Password = string(adminPassword)
			model.CreateAdmin(admin)
		}
	}
	//启动路由
	r := routers.SetupRouter()
	//启动服务
	r.Run(utils.HostPort)
}
