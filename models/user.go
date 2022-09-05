package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)


type User struct{
	ID						int					`gorm:"primaryKey;autoIncrement" json:"id" query:"id" form:"id"`
	TelegramUsername		string				`gorm:"type:varchar(50);default:null" json:"telegram_username" query:"telegram_username" form:"telegram_username"`
	TelegramID				string				`gorm:"type:varchar(50);default:null;uniqueIndex" json:"telegram_id" query:"telegram_id" form:"telegram_id"`
	FirstName				string				`gorm:"type:varchar(50);default:null" json:"first_name" query:"first_name" form:"first_name" validation:"required,alphbet_fa"`
	LastName				string				`gorm:"type:varchar(50);default:null" json:"last_name" query:"last_name" form:"last_name" validation:"required,alphbet_fa"`
	Mobile					string				`gorm:"type:varchar(11);default:null;uniqueIndex" json:"mobile" query:"mobile" form:"mobile" validation:"required,mobile,unique"`
	Role					int					`gorm:"type:int" json:"role" query:"role" form:"role"`
	OtpCode					int					`gorm:"type:int" json:"otp_code" query:"otp_code" form:"otp_code"`
	Status					int					`gorm:"type:int" json:"status" query:"status" form:"status"`
	ExpireOtpAt				time.Time			`gorm:"autocreatetime" json:"expire_otp_at" query:"expire_otp_at" form:"expire_otp_at"`
	CreatedAt				time.Time			`gorm:"autocreatetime" json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt				time.Time		    `gorm:"autoupdatetime" json:"update_at"`
	DeletedAt				gorm.DeletedAt		`json:"deleted_at"`
}


type Claims struct {
	ID int `json:"id"`
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}


var JwtKey = []byte("saeed")


var UserRoleList = map[int]string{
	0 : "کاربرعادی",
	1 : "فروشگاه‌دار",
	100 : "مدیر‌کل",
}

var UserGenderList = map[int]string{
	1 : "آقا",
	2 : "خانم",
}

var UserStatusList = map[int]string{
	1 : "فعال",
	2 : "مسدود",
}