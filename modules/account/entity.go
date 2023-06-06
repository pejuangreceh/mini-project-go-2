package account

import "gorm.io/gorm"

type Actors struct {
	gorm.Model
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RoleID     uint8  `json:"role_id" binding:"required"`
	IsVerified string `json:"is_verified" binding:"required"`
	IsActive   string `json:"is_active" binding:"required"`
}

type Register struct {
	AdminID      uint8  `json:"admin_id" binding:"required"`
	SuperAdminID uint8  `json:"superadmin_id" binding:"required"`
	Status       string `json:"status"`
}

func (Register) TableName() string {
	return "register" // specify the actual table name here
}
