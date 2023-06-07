package dto

type AllActorResponse struct {
	Message string              `json:"message"`
	Data    []ActorDataResponse `json:"data"`
	Token   string              `json:"token"`
}

type ActorDataResponse struct {
	ID         uint8  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleID     uint8  `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type CustomerDataResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type AllCustomerResponse struct {
	Message string                 `json:"message"`
	Data    []CustomerDataResponse `json:"data"`
}
type ErrorResponse struct {
	Error string
}
type MyClaimsResponse struct {
	ID         uint8  `mapstructure:"id"`
	Username   string `mapstructure:"username"`
	RoleID     uint8  `mapstructure:"role_id"`
	IsVerified string `mapstructure:"is_verified"`
	IsActive   string `mapstructure:"is_active"`
	Iat        int64  `mapstructure:"iat"`
	Exp        int64  `mapstructure:"exp"`
}
