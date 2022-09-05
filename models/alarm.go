package models

import(
	"gorm.io/gorm"
	"time"
)

type Alarm struct{
	ID							int					`gorm:"primaryKey;autoIncrement" json:"id" query:"close_time" form:"close_time"`
	Symbol						float64				`gorm:"type:varchar(50);" json:"symbol" query:"symbol" form:"symbol"`
	UerrID						int				    `json:"user_id" query:"user_id" form:"user_id"`
	CoinID						float64				`json:"coin_id" query:"coin_id" form:"coin_id"`
	High						float64				`gorm:"type:float;" json:"high" query:"high" form:"high"`
	Low							float64				`gorm:"type:float;" json:"low" query:"low" form:"low"`
	Status						int					`gorm:"type:int;" json:"status" query:"status" form:"status"`
	Type						int					`gorm:"type:float;" json:"type" query:"type" form:"type"`
	Priority					int					`json:"priority" query:"priority" form:"priority"`
	CreatedAt					time.Time			`gorm:"autocreatetime" json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt					time.Time		    `gorm:"autoupdatetime" json:"update_at"`
	DeletedAt					gorm.DeletedAt		`json:"deleted_at"`
}