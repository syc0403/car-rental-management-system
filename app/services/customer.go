package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
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
