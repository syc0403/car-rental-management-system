package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"strconv"
)

type roleService struct {
}

var RoleService = new(roleService)

// GetAllRole 获取全部角色
func (roleService *roleService) GetAllRole(query *request.Pagination) (role []models.Role, total int64, err error) {
	db := global.App.DB.Model(&role)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&role).Error
	return role, total, err

}

// UpdateRole 修改角色信息
func (roleService *roleService) UpdateRole(params *models.Role) (role models.Role, err error) {
	if err = global.App.DB.Model(&role).Where("id=?", params.ID).Updates(&params).Error; err != nil {
		err = errors.New("修改角色信息失败")
		return
	}
	err = global.App.DB.Where("id = ?", params.ID).First(&role).Error
	return
}

// AddRole 创建角色
func (roleService *roleService) AddRole(query *models.Role) (role models.Role, err error) {
	if err = global.App.DB.Create(&query).Error; err != nil {
		err = errors.New("添加角色失败")
		return
	}
	err = global.App.DB.Where("id = ?", query.ID).First(&role).Error
	return
}

// DeteleCheck 删除角色
func (roleService *roleService) DeteleRole(id string) (role models.Role, err error) {
	roleId, _ := strconv.Atoi(id)
	if err = global.App.DB.First(&role, roleId).Error; err != nil {
		err = errors.New("角色不存在")
		return
	}
	err = global.App.DB.Model(&role).Where("id=?", roleId).Delete(&role).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}
