package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"strconv"
)

type checkService struct {
}

var CheckService = new(checkService)

// GetAllCheck 获取全部检查单
func (checkService *checkService) GetAllCheck(query *request.Pagination) (check []models.Check, total int64, err error) {
	db := global.App.DB.Model(&check)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&check).Error
	return check, total, err

}

// AddCheck 创建检查单
func (checkService *checkService) AddCheck(query *models.Check) (check models.Check, err error) {
	if err = global.App.DB.Create(&query).Error; err != nil {
		err = errors.New("添加检查单失败")
		return
	}
	err = global.App.DB.Where("id = ?", query.ID).First(&check).Error
	return
}

// DeteleCheck 删除检查单
func (checkService *checkService) DeteleCheck(id string) (check models.Check, err error) {
	CheckId, _ := strconv.Atoi(id)
	if err = global.App.DB.First(&check, CheckId).Error; err != nil {
		err = errors.New("检查不存在")
		return
	}
	err = global.App.DB.Model(&check).Where("id=?", CheckId).Delete(&check).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}
