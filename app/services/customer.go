package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"strconv"
)

type customerService struct {
}

var CustomerService = new(customerService)

// GetAllCustomerInfo 获取全部客户信息
func (customerService *customerService) GetAllCustomerInfo(query *request.Pagination) (err error, total int64, customer []models.Customer) {
	db := global.App.DB.Model(&customer)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return err, 0, nil
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&customer).Error
	return err, total, customer

}

// GetCustomerInfoById 根据用户id获取客户信息
func (customerService *customerService) GetCustomerInfoById(id string) (err error, customer models.Customer) {
	err = global.App.DB.Where("id = ?", id).First(&customer).Error
	return

}

// AddCustomer 添加客户
func (customerService *customerService) AddCustomer(query *models.Customer) (err error, customer models.Customer) {
	if err = global.App.DB.Create(&query).Error; err != nil {
		err = errors.New("添加客户失败")
		return
	}
	err = global.App.DB.Where("id = ?", query.ID).First(&customer).Error
	return
}

// DeteleCustomer 删除客户
func (customerService *customerService) DeteleCustomer(id string) (err error, customer models.Customer) {
	customerId, err := strconv.Atoi(id)
	if err = global.App.DB.First(&customer, customerId).Error; err != nil {
		err = errors.New("用户不存在")
		return
	}

	err = global.App.DB.Model(&customer).Where("id=?", customerId).Delete(&customer).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}

// UpdateCustomerInfo 修改客户信息
func (customerService *customerService) UpdateCustomerInfo(params *models.Customer) (err error, customer models.Customer) {
	if err = global.App.DB.Model(&customer).Where("id=?", params.ID).Update("sex", params.Sex).Error; err != nil {
		err = errors.New("修改客户信息失败")
		return
	}
	if err = global.App.DB.Model(&customer).Where("id=?", params.ID).Updates(&params).Error; err != nil {
		err = errors.New("修改客户信息失败")
		return
	}
	err = global.App.DB.Where("id = ?", params.ID).First(&customer).Error
	return
}

// GetCustomerInfoByName 根据用户姓名模糊查询客户信息
func (customerService *customerService) GetCustomerInfoByName(query *request.GetCustomerInfoByName) (total int64, customer []models.Customer, err error) {
	db := global.App.DB.Model(&customer)
	err = db.Where("deleted_at is null").Where("customer_name like ? ", "%"+query.CustomerName+"%").Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Where("customer_name like ? ", "%"+query.CustomerName+"%").Order("created_at desc").Find(&customer).Error
	return total, customer, err

}

// GetCustomerInfoByIdentity 根据身份证获取客户信息
func (customerService *customerService) GetCustomerInfoByIdentity(identity string) (err error, customer models.Customer) {
	err = global.App.DB.Where("identity = ?", identity).First(&customer).Error
	return

}
