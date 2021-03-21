package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/modules/research/paper/model"
	"gorm.io/gorm"
)

func Create(u model.ResearchPaper) (err error, userInter model.ResearchPaper) {
	var user model.ResearchPaper
	if !errors.Is(global.GVA_DB.Where("title = ?", u.Title).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("论文名称已存在"), userInter	}
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func Update(reqUser model.ResearchPaper) (err error, user model.ResearchPaper) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

func Get(id uint) (err error, sysDictionaryDetail model.ResearchPaper) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

func Delete(sysDictionaryDetail model.ResearchPaper) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ResearchPaper{}).Unscoped()
	var userList []model.ResearchPaper
	err = db.Count(&total).Where("").Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}
