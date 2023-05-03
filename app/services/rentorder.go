package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
)

type rentOrderService struct {
}

var RentOrderService = new(rentOrderService)

// GetAllRentOrder 获取全部车辆信息
func (rentOrderService *rentOrderService) GetAllRentOrder(query *request.GetAllRentOrder) (total int64, rentOrder []models.RentOrder, err error) {
	db := global.App.DB.Model(&rentOrder)
	var customerId int
	db.Table("customers").Select("id").Where("identity = ?", query.CustomerIdentity).Scan(&customerId)
	db = global.App.DB.Model(&rentOrder)
	err = db.Where("deleted_at is null").Where("customer_id = ? AND identity like ? and car_number like ?", customerId, query.Identity, query.CarNumber).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&rentOrder).Error
	return total, rentOrder, err

}
