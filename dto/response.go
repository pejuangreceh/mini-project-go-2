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
