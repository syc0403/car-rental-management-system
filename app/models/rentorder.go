package models

import "time"

type RentOrder struct {
	ID
	CustomerId int32     `json:"customer_id" gorm:"comment:客户id"`
	OrderPrice string    `json:"order_price" gorm:"type:double(10,2);comment:租赁价格"`
	BeginDate  time.Time `json:"begin_date" gorm:"comment:开始时间"`
	ReturnDate time.Time `json:"return_date" gorm:"comment:归还时间"`
	Identity   string    `json:"identity" gorm:"varchar(255);comment:订单编号"`
	CarNumber  string    `json:"car_number" gorm:"type:varchar(255);not null;comment:车牌号"`
	Timestamps
	SoftDeletes
}
