package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("phone = ?", params.Phone).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{LoginName: params.LoginName, Phone: params.Phone, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("login_name = ?", params.LoginName).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

// UpdateUserInfo 修改用户信息
func (userService *userService) UpdateUserInfo(params *models.User) (err error, user models.User) {
	err = global.App.DB.Model(&user).Where("id=?", params.ID).Updates(&params).Error
	if err != nil {
		err = errors.New("修改用户信息失败")
		return
	}
	err = global.App.DB.Where("id = ?", params.ID).First(&user).Error
	return

}

// DeteleUser 删除用户
func (userService *userService) DeteleUser(id string) (err error, user models.User) {
	userId, err := strconv.Atoi(id)
	if err = global.App.DB.First(&user, userId).Error; err != nil {
		err = errors.New("用户不存在")
	} else {
		if err = global.App.DB.Model(&user).Where("id=?", userId).Delete(&user).Error; err != nil {
			err = errors.New("删除失败")
		}
	}
	return
}

// GetAllUserInfo 获取全部用户信息

func (userService *userService) GetAllUserInfo(query *request.Pagination) (err error, total int64, user []models.User) {
	db := global.App.DB.Model(&user)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return err, 0, nil
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&user).Error
	return err, total, user

}
