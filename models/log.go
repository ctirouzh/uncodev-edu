package models

import(
	"time"
)


type Log struct{
	ID						int					`json:"id" query:"id" form:"id"`
	UserID					int					`json:"user_id" query:"user_id" form:"user_id"`
	Type					string				`gorm:"type:varchar(30);" json:"type" query:"type" form:"type"`
	Error					string				`json:"error" query:"error" form:"error"`
	Route					string				`gorm:"type:varchar(50);" json:"route" query:"route" form:"route"`
	CreatedAt				time.Time			`gorm:"autocreatetime" json:"created_at" query:"created_at" form:"created_at"`
}
