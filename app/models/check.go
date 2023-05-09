package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Check struct {
	ID                int             `json:"id" gorm:"type:int(11);primaryKey"`
	Identity          string          `json:"identity" gorm:"type:varchar(255);not null;comment:检查单编号"`
	RentOrderIdentity string          `json:"rentorder_identity" gorm:"type:varchar(255);not null;comment:订单编号"`
	Problem           string          `json:"problem" gorm:"type:varchar(255);comment:车辆问题"`
	ProblemDesc       string          `json:"problem_desc" gorm:"type:varchar(255);comment:问题描述"`
	PayPrice          decimal.Decimal `json:"pay_price" gorm:"type:double(10,2);comment:赔付价格"`
	OperName          string          `json:"oper_name" gorm:"type:varchar(255);comment:记录员名称"`
	CheckTime         string          `json:"check_time" gorm:"type:varchar(255);comment:检查时间"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
