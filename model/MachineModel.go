package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"vulnsite/dao"
)

type Machine struct {
	gorm.Model
	Title string `gorm:"not null"`
	Author string
	Desc string `gorm:"type:varchar(1024)"`
	DownAddr string 
	DownPwd string
	Difficulty int `gorm:"default:1"`
}

func GetCountMachineById(id uint) (count int, err error) {
	err = dao.DB.Model(&Machine{}).Where("id=?", id).Count(&count).Error
	return
}

func FindMachine(pageNum int, pageSize int) (machines []*Machine, total int, err error) {
	err = dao.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&machines).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&Machine{}).Count(&total).Error
	return
}

func FindMachineByDiff(pageNum int, pageSize int, difficulty int) (machines []*Machine, total int, err error) {
	err = dao.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&machines, "difficulty=?", difficulty).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&Machine{}).Where("difficulty=?", difficulty).Count(&total).Error
	return
}

func CreateMachine(machine *Machine) error {
	return dao.DB.Create(&machine).Error
}

func DeleteMachine(id int) error {
	return dao.DB.Unscoped().Delete(&Machine{}, "id=?", id).Error
}

func UpdateMachine(machine *Machine) error {
	return dao.DB.Model(&Machine{}).Update(&machine).Error
}

func SearchMachine(pageNum int, pageSize int, keys string) (machines []*Machine, total int, err error) {
	keyArr := strings.Split(keys, " ")
	keyNum := len(keyArr)
	var whereStr = ""
	for i := 0; i < keyNum; i++ {
		whereStr = whereStr + "`title` LIKE '%"+ keyArr[i] +"%' OR `author` LIKE '%"+ keyArr[i] +"%' OR `desc` LIKE '%"+ keyArr[i] +"%'"
		if i < keyNum - 1 {
			whereStr = whereStr + " or "
		}
	}
	fmt.Println(whereStr)
	err = dao.DB.Debug().Where(whereStr).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&machines).Error
	if err != nil {
		return
	}
	err = dao.DB.Debug().Model(&Machine{}).Where(whereStr).Count(&total).Error
	return
}