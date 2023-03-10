package services

import (
	"car-rental-management-system/app/common/request"
	"car-rental-management-system/app/models"
	"car-rental-management-system/global"
	"car-rental-management-system/utils"
	"errors"
	"strconv"
)

type carService struct {
}

var CarService = new(carService)

// GetAllCarInfo 获取全部车辆信息
func (carService *carService) GetAllCarInfo(query *request.Pagination) (err error, total int64, car []models.Car) {
	db := global.App.DB.Model(&car)
	err = db.Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return err, 0, nil
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Order("created_at desc").Find(&car).Error
	return err, total, car

}

// GetCarInfoById 根据用户id获取车辆信息
func (carService *carService) GetCarInfoById(id string) (err error, Car models.Car) {
	err = global.App.DB.Where("id = ?", id).First(&Car).Error
	return

}

// AddCar 添加车辆
func (carService *carService) AddCar(query *models.Car) (err error, Car models.Car) {
	if err = global.App.DB.Create(&query).Error; err != nil {
		err = errors.New("添加车辆失败")
		return
	}
	err = global.App.DB.Where("id = ?", query.ID).First(&Car).Error
	return
}

// DeteleCar 删除车辆
func (carService *carService) DeteleCar(id string) (err error, car models.Car) {
	carId, err := strconv.Atoi(id)
	if err = global.App.DB.First(&car, carId).Error; err != nil {
		err = errors.New("车辆不存在")
		return
	}

	err = global.App.DB.Model(&car).Where("id=?", carId).Delete(&car).Error
	if err != nil {
		err = errors.New("删除失败")
	}
	return
}

// UpdateCarInfo 修改车辆信息
func (carService *carService) UpdateCarInfo(params *models.Car) (err error, car models.Car) {
	if err = global.App.DB.Model(&car).Where("id=?", params.ID).Updates(&params).Error; err != nil {
		err = errors.New("修改客户信息失败")
		return
	}
	err = global.App.DB.Where("id = ?", params.ID).First(&car).Error
	return
}

// GetCarInfoByNumber 根据车牌模糊查询车辆信息
func (carService *carService) GetCarInfoByNumber(query *request.GetCarInfoByNumber) (err error, total int64, car []models.Car) {
	db := global.App.DB.Model(&car)
	err = db.Where("deleted_at is null").Where("car_number like ? ", "%"+query.CarNumber+"%").Count(&total).Error
	if err != nil {
		return err, 0, nil
	}
	err = db.Scopes(utils.Paginate(query.Current, query.PageSize)).Where("car_number like ? ", "%"+query.CarNumber+"%").Order("created_at desc").Find(&car).Error
	return err, total, car

}

