package models

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)


type User struct{
	ID						int					`gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName				string				`gorm:"type:varchar(50);default:null" json:"first_name" `
	LastName				string				`gorm:"type:varchar(50);default:null" json:"last_name" `
	Mobile					string				`gorm:"type:varchar(11);default:null;uniqueIndex"`
	Role					int					`gorm:"type:int" json:"role" query:"role"`
	OtpCode					int					`gorm:"type:int" json:"otp_code" query:"otp_code" `
	Status					int					`gorm:"type:int" json:"status"`
	ExpireOtpAt				time.Time			`gorm:"autocreatetime" json:"expire_otp_at"`
	CreatedAt				time.Time			`gorm:"autocreatetime" json:"created_at"`
	UpdatedAt				time.Time		    `gorm:"autoupdatetime" json:"update_at"`
	DeletedAt				gorm.DeletedAt		`json:"deleted_at"`
}


type Claims struct {
	ID int `json:"id"`
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}


var JwtKey = []byte("saeed")


