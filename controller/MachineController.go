package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"vulnsite/Response"
	"vulnsite/model"
)

//查询虚拟机
func FindMachine(c *gin.Context) {
	//获取参数
	sPageNum := c.Param("pageNum")
	sPageSize := c.Param("pageSize")
	pageNum, errNum := strconv.Atoi(sPageNum)
	pageSize, errSize := strconv.Atoi(sPageSize)
	//判断参数是否有误
	if errSize != nil || errNum != nil {
		Response.ParameterError(c)
		return
	}
	//查询
	machines, total, err :=model.FindMachine(pageNum, pageSize)
	if err != nil {
		Response.SystemError(c)
		return
	}
	//返回数据
	Response.Response(
		c,
		Response.SUCCESS,
		Response.MSG[Response.SUCCESS],
		gin.H{
			"MachinesList": machines,
			"Total": total,
		})
}

//按照难度查询虚拟机
func FindMachineByDiff(c *gin.Context) {
	//获取参数
	sPageNum := c.Param("pageNum")
	sPageSize := c.Param("pageSize")
	sDifficulty := c.Param("difficulty")
	pageNum, errNum := strconv.Atoi(sPageNum)
	pageSize, errSize := strconv.Atoi(sPageSize)
	difficulty, errDiff := strconv.Atoi(sDifficulty)
	//判断参数是否有误
	if errSize != nil || errNum != nil || errDiff != nil{
		Response.ParameterError(c)
		return
	}
	//查询
	machines, total, err :=model.FindMachineByDiff(pageNum, pageSize, difficulty)
	if err != nil {
		Response.SystemError(c)
		return
	}
	//返回数据
	Response.Response(
		c,
		Response.SUCCESS,
		Response.MSG[Response.SUCCESS],
		gin.H{
			"MachinesList": machines,
			"total": total,
		})

}

//搜索虚拟机
func SearchMachine(c *gin.Context) {
	//获取参数
	sPageNum := c.Param("pageNum")
	sPageSize := c.Param("pageSize")
	keys := c.Param("keys")
	pageNum, errNum := strconv.Atoi(sPageNum)
	pageSize, errSize := strconv.Atoi(sPageSize)
	//判断参数是否有误
	if errSize != nil || errNum != nil {
		Response.ParameterError(c)
		return
	}
	//查询
	machines, total, err :=model.SearchMachine(pageNum, pageSize, keys)
	if err != nil {
		Response.SystemError(c)
		return
	}
	//返回数据
	Response.Response(
		c,
		Response.SUCCESS,
		Response.MSG[Response.SUCCESS],
		gin.H{
			"MachinesList": machines,
			"total": total,
		})
}

//添加虚拟机
func CreateMachine(c *gin.Context) {
	//获取参数
	var machine = new(model.Machine)
	err := c.ShouldBindJSON(&machine)
	//判断参数是否为空,参数是否错误
	if err != nil || machine.Title == "" {
		Response.ParameterError(c)
		return
	}
	//获取认证信息,并自动添加创建人姓名
	authInfo, isExist := c.Get("AuthInfo")
	if !isExist {
		Response.PermissionDenied(c)
		return
	}
	adminInfo := authInfo.(model.AdminInfo)
	machine.Author = adminInfo.Name
	//创建虚拟机
	err = model.CreateMachine(machine)
	if err != nil {
		Response.SystemError(c)
		return
	}
	//成功
	Response.Success(c)
}

//删除虚拟机
func DeleteMachine(c *gin.Context) {
	//获取id
	sId := c.Param("id")
	id, err := strconv.Atoi(sId)
	//判断参数是否有误
	if sId == "" || err != nil {
		Response.ParameterError(c)
		return
	}
	//删除虚拟机
	if err = model.DeleteMachine(id); err != nil {
		Response.SystemError(c)
		return
	}
	//成功
	Response.Success(c)
}

//修改虚拟机
func UpdateMachine(c *gin.Context) {
	//获取参数
	machine := new(model.Machine)
	c.ShouldBindJSON(&machine)
	//判断id是否为空
	if machine.ID == 0 {
		Response.ParameterError(c)
		return
	}
	//判断修改的记录是否存在
	count, err := model.GetCountMachineById(machine.ID)
	if err != nil {
		Response.SystemError(c)
		return
	}
	if count == 0 {
		Response.Response(
			c,
			Response.RECORD_NOT_EXIST,
			Response.MSG[Response.RECORD_NOT_EXIST],
			nil,
			)
		return
	}
	//修改记录
	err = model.UpdateMachine(machine)
	if err != nil {
		Response.SystemError(c)
		return
	}
	Response.Success(c)
}