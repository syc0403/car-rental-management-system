package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type rentOrderService struct {
}

var RentOrderService = new(rentOrderService)

// GetAllRentOrder 获取全部订单信息
func (rentOrderService *rentOrderService) GetAllRentOrder(query *request.Pagination) (rentOrder []models.RentOrder, total int64, err error) {
	db := global.App.DB.Model(&rentOrder)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&rentOrder).Error
	return rentOrder, total, err

}

// GetRentOrderBy 根据条件获取订单信息
func (rentOrderService *rentOrderService) GetRentOrderBy(query *request.GetAllRentOrder) (total int64, rentOrder []models.RentOrder, err error) {
	db := global.App.DB.Model(&rentOrder)
	var customerId int
	db.Table("customers").Select("id").Where("identity = ?", query.CustomerIdentity).Scan(&customerId)
	db = global.App.DB.Model(&rentOrder)
	err = db.Where("deleted_at is null").Where("customer_id = ? or identity = ? or car_number = ? or status = ?", customerId, query.Identity, query.CarNumber, query.Status).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&rentOrder).Error
	return total, rentOrder, err

}

// CreateRentIdentity 根据条件生成订单号
func (rentOrderService *rentOrderService) CreateRentIdentity(query *request.CreateRentIdentity) (rentOrderIdentity string, err error) {
	timeUnix := time.Now().Unix()
	rentOrderIdentity = fmt.Sprintf("CZ_%s%s_%d", query.CustomerIdentity[len(query.CustomerIdentity)-4:], query.CarNumber[len(query.CarNumber)-4:], timeUnix)
	return rentOrderIdentity, err

}

// AddRentOrder 创建订单
func (rentOrderService *rentOrderService) AddRentOrder(query *models.RentOrder) (rentOrde models.RentOrder, err error) {
	if err = global.App.DB.Create(&query).Error; err != nil {
		err = errors.New("添加订单失败")
		return
	}
	err = global.App.DB.Where("id = ?", query.ID).First(&rentOrde).Error
	return
}

// DeteleRentOrder 删除订单
func (rentOrderService *rentOrderService) DeteleRentOrder(id string) (rentOrde models.RentOrder, err error) {
	RentOrderId, _ := strconv.Atoi(id)
	if err = global.App.DB.First(&rentOrde, RentOrderId).Error; err != nil {
		err = errors.New("订单不存在")
		return
	}
	err = global.App.DB.Model(&rentOrde).Where("id=?", RentOrderId).Delete(&rentOrde).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}

// UpdateRentOrderStatus 修改出租单状态
func (rentOrderService *rentOrderService) UpdateRentOrder(params *request.UpdateRentOrderStatus) (rentOrde models.RentOrder, err error) {
	if err = global.App.DB.Model(&rentOrde).Where("identity=?", params.Identity).Update("status", params.Status).Error; err != nil {
		err = errors.New("修改出租单状态失败")
		return
	}
	err = global.App.DB.Where("identity = ?", params.Identity).First(&rentOrde).Error
	return
}
