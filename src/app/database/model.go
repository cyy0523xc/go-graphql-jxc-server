package database

import (
	//"github.com/jinzhu/gorm"
	"time"
)

// 用户
type User struct {
	ID        int32
	Name      string `gorm:"size:100;index:idx_name"`       // 姓名
	Phone     string `gorm:"not null;type:char(11);unique"` //手机号，用于登陆
	Password  string `gorm:"not null;type:varchar(150)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 商品, 以最小SKU为标准
type Goods struct {
	ID        int32
	Code      string `gorm:"type:varchar(100)"`       // 商品码
	Name      string `gorm:"size:255;index:idx_name"` // 名称
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 商品基本统计信息
type GoodsBaseStat struct {
	GoodsID             int32     `gorm:"unique"`
	CostPrice           float32   // 成本价
	MSRP                float32   `gorm:"default 0"` // 厂商建议售价
	Volume              int32     // 库存数量
	TotalPurchaseVolume int32     // 总进货数量
	TotalPurchaseTimes  int32     // 总进货次数
	TotalAmount         float32   // 总金额
	LastPurchsedAt      time.Time // 最后进货时间
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}

// 商品属性配置表
// TODO 可以增加商品分类, 不同的分类可以配置不同的属性
type GoodsAttrConfig struct {
	ID        int32
	Attr      string `gorm:"type:varchar(31);unique"` // 属性
	ValueType string `gorm:"type:varchar(50)"`        // 属性值的类型
	Remark    string `gorm:"type:varchar(255)"`       // 备注
}

// 商品属性
type GoodsAttr struct {
	ID        int32
	GoodsID   int32
	AttrID    int32
	Value     string `gorm:"type:varchar(255)"` // 属性值的类型
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 进货
type Purchase struct {
	ID           int32
	GoodsID      int32
	MSRP         float32 `gorm:"default 0"` // 厂商建议售价
	CostPrice    float32 // 进货价
	Volume       int32   // 进货数量
	StoreVolume  int32   // 库存数量, 该批次的商品销售完了，库存
	TotalAmount  float32 // 总金额
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Remark       string `gorm:"type:varchar(255)"` // 进货备注
	UpdateRemark string `gorm:"type:varchar(255)"` // 修改备注
}

// 订单记录
type Order struct {
	ID           int32
	GoodsID      int32
	PurchaseID   int32   // 进货批次
	MSRP         float32 `gorm:"default 0"` // 厂商建议售价
	CostPrice    float32 // 进货价
	Volume       int32   // 销售数量
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Remark       string `gorm:"type:varchar(255)"` // 进货备注
	UpdateRemark string `gorm:"type:varchar(255)"` // 修改备注
}
