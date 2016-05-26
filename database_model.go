package main

import (
	"github.com/jinzhu/gorm"
)

// 用户
type User struct {
	ID        int
	Name      string `gorm:"size:100;index:idx_name"` // 姓名
	Phone     string `gorm:"not_null;size:11;unique"` //手机号，用于登陆
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 商品
type Goods string {
	ID        int
	Name      string `gorm:"size:255;index:idx_name"` // 名称
	CostPrice float32                                 // 成本价
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

}
