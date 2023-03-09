package model

import (
	"vulnsite/dao"
)

type AdminInfo struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `gorm:"unique;not null"`
	Name string `gorm:"not null"`
	Grade int `gorm:"default:2"`
}

type Admin struct {
	AdminInfo
	Password string `gorm:"not null"`
}

func GetAdminByUsername(username string) (admin *Admin, err error) {
	admin = new(Admin)
	if err = dao.DB.Where("username=?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return
}

func GetCountAdminByUsername(username string) (count int, err error) {
	err = dao.DB.Model(&Admin{}).Where("username=?", username).Count(&count).Error
	return
}

func GetAdminById(id int) (admin *Admin, err error) {
	admin = new(Admin)
	if err = dao.DB.Where("id=?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return
}

func GetCountAdminById(id int) (count int, err error) {
	err = dao.DB.Model(&Admin{}).Where("id=?", id).Count(&count).Error
	return
}

func GetAllAdmin() (AdminInfos []*AdminInfo, err error) {
	if err =  dao.DB.Table("admins").Find(&AdminInfos).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteAdmin(id int) error {
	return dao.DB.Delete(&Admin{}, "id=?", id).Error
}

func CreateAdmin(admin *Admin) error {
	return dao.DB.Create(&admin).Error
}

func UpdateAdmin(admin *Admin) error {
	return dao.DB.Model(&Admin{}).Update(&admin).Error
}