package models

import(
	"gorm.io/gorm"
	"time"
)


type Order struct{
	ID					int					`json:"id" query:"id" form:"id" validation:"required"`
	Uid					int					`json:"uid" query:"uid" form:"uid" validation:"required"`
	CoinId		    	int					`json:"coin_id" query:"coin_id" form:"coin_id"`
	WalletId		    int					`json:"wallet_id" query:"wallet_id" form:"wallet_id"`
	Amount		    	int					`json:"amount" query:"amount" form:"amount"`
	Price				int					`json:"price" query:"price" form:"price" validation:"required"`
	Priority			int					`json:"priority" query:"priority" form:"priority" validation:"required"`
	Volume				float64				`json:"volume" query:"volume" form:"volume"`
	Fill				float64				`json:"fill" query:"fill" form:"fill"`
	Status				int					`json:"status" query:"status" form:"status" validation:"required"`
	Type				int					`json:"type" query:"type" form:"type" validation:"required"`
	CreatedAt			time.Time			`gorm:"autocreatetime" json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt			time.Time		    `gorm:"autoupdatetime" json:"update_at"`
	DeletedAt			gorm.DeletedAt		`json:"deleted_at"`
}

	



var OrderStatusList = map[int]string{
	1 : "ایجاد‌شده",
	2 : "تکمیل‌شده",
	3 : "کنسل‌شده",
}
