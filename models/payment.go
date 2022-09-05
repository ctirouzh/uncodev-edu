package models

import(
	"gorm.io/gorm"
)


type Payment struct{
	ID					int					`json:"id" query:"id" form:"id"`
	Uid					int					`json:"uid" query:"uid" form:"uid"`
	WalletId			int					`json:"wallet_id" query:"wallet_id" form:"wallet_id"`
	Amount				float64				`json:"amount" query:"amount" form:"amount"`
	Code				int					`json:"code" query:"code" form:"code"`
	RefId				int					`json:"ref_id" query:"ref_id" form:"ref_id"`
	Au					int					`json:"au" query:"au" form:"au"`
	Note				int					`json:"note" query:"note" form:"note"`
	Status				int					`json:"status" query:"status" form:"status"`
	CreatedAt			string				`json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt			string				`json:"update_at"`
	DeletedAt			gorm.DeletedAt		`json:"deleted_at"`
}


var PaymentStatusList = map[int]string{
	1 : "فعال",
	2 : "غیر‌فعال",
}
