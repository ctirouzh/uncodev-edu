package models

import(
	"gorm.io/gorm"
	"time"
)


type GoldWallet struct{
	ID					int					`json:"id" query:"id" form:"id" validation:"required"`
	Uid					int					`json:"uid" query:"uid" form:"uid" validation:"required"`
	OrderId				int					`json:"order_id" query:"order_id" form:"order_id"`
	PaymentId			int					`json:"payment_id" query:"payment_id" form:"payment_id"`
	Amount				float64				`json:"amount" query:"amount" form:"amount"`
	Inventory			float64				`json:"inventory" query:"inventory" form:"inventory"`
	PreviousInventory	float64				`json:"previous_inventory" query:"previous_inventory" form:"previous_inventory"`
	Lock				int					`json:"lock" query:"lock" form:"lock"`
	Note				int					`json:"note" query:"note" form:"note"`
	Status				int					`json:"status" query:"status" form:"status"`
	CreatedAt			time.Time			`gorm:"autocreatetime" json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt			time.Time		    `gorm:"autoupdatetime" json:"update_at"`
	DeletedAt			gorm.DeletedAt		`json:"deleted_at"`
}


var GoldWalletStatusList = map[int]string{
	1 : "فعال",
	2 : "غیر‌فعال",
}
