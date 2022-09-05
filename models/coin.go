package models

import(
	"gorm.io/gorm"
)

type Coin struct{
	ID							int					`json:"id" query:"close_time" form:"close_time"`
	Symbol						string				`json:"symbol" query:"symbol" form:"symbol"`
	CategoryID					int					`json:"category_id" query:"category_id" form:"category_id"`
	Close						float64				`json:"close" query:"close" form:"close"`
	Change						float64				`json:"change" query:"change" form:"change"`
	Height						float64				`json:"height" query:"height" form:"height"`
	Moving50					float64				`json:"moving50" query:"moving50" form:"moving50"`
	Moving100					float64				`json:"Moving100" query:"Moving100" form:"Moving100"`
	Moving200					float64				`json:"moving200" query:"moving200" form:"moving200"`
	UpperShadowPercent			float64				`json:"upper_shadow_percent" query:"upper_shadow_percent" form:"upper_shadow_percent"`
	LowerShadowPercent			float64				`json:"lower_shadow_percent" query:"lower_shadow_percent" form:"lower_shadow_percent"`
	OpenTime					int					`json:"open_time" query:"open_time" form:"open_time"`
	CloseTime					int					`json:"close_time" query:"close_time" form:"close_time"`
	Status						int					`json:"status" query:"status" form:"status"`
	CreatedAt					string				`json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt					string				`json:"update_at"`
	DeletedAt					gorm.DeletedAt		`json:"deleted_at"`
}