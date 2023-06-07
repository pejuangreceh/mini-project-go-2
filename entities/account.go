package entities

import "gorm.io/gorm"

type Actors struct {
	gorm.Model
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RoleID     uint8  `json:"role_id" binding:"required"`
	IsVerified string `json:"is_verified" binding:"required"`
	IsActive   string `json:"is_active" binding:"required"`
}

type Approval struct {
	gorm.Model
	IsVerified string `json:"is_verified" binding:"required"`
}

type Activate struct {
	gorm.Model
	IsActive string `json:"is_active" binding:"required"`
}

type Register struct {
	AdminID      uint8  `json:"admin_id" binding:"required"`
	SuperAdminID uint8  `json:"superadmin_id" binding:"required"`
	Status       string `json:"status"`
}
type RegisterStatus struct {
	AdminID uint8  `json:"admin_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

func (Approval) TableName() string {
	return "actors" // specify the actual table name here
}
func (Activate) TableName() string {
	return "actors" // specify the actual table name here
}

func (Register) TableName() string {
	return "register" // specify the actual table name here
}
func (RegisterStatus) TableName() string {
	return "register" // specify the actual table name here
}
