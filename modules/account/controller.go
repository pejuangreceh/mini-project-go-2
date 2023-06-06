package account

import (
	"fmt"
	"gorm.io/gorm"
)

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

type ActorDataResponse struct {
	ID         uint8  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleID     uint8  `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type AllResponse struct {
	Message string              `json:"message"`
	Data    []ActorDataResponse `json:"data"`
	Token   string              `json:"token"`
}

func (c Controller) Create(body *CreateRequest) (*AllResponse, error) {
	actors := Actors{
		Model:      gorm.Model{},
		Username:   body.Username,
		Password:   body.Password,
		RoleID:     body.RoleID,
		IsVerified: body.IsVerified,
		IsActive:   body.IsActive,
	}
	err := c.useCase.Create(&actors)
	if err != nil {
		return nil, err
	}
	res := ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}
	allres := &AllResponse{
		Message: "Data berhasil dibuat",
	}
	allres.Data = append(allres.Data, res)

	return allres, nil
}

// Get All Data
func (c Controller) Read() (*AllResponse, error) {
	actors, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}
	res := &AllResponse{}
	for _, actor := range actors {
		c := ActorDataResponse{
			ID:         uint8(actor.ID),
			Username:   actor.Username,
			Password:   actor.Password,
			RoleID:     actor.RoleID,
			IsVerified: actor.IsVerified,
			IsActive:   actor.IsActive,
		}
		res.Data = append(res.Data, c)
	}
	res.Message = "Data Admin Sukses dimuat"
	return res, nil
}

func (c Controller) ReadID(ID string) (*AllResponse, error) {
	actors, err := c.useCase.ReadID(ID)
	if err != nil {
		return nil, err
	}

	if len(actors) == 0 {
		return nil, fmt.Errorf("Actors not found")
	}

	res := ActorDataResponse{
		ID:         uint8(actors[0].ID),
		Username:   actors[0].Username,
		Password:   actors[0].Password,
		RoleID:     actors[0].RoleID,
		IsVerified: actors[0].IsVerified,
		IsActive:   actors[0].IsActive,
	}
	allres := &AllResponse{
		Message: "Data Admin berhasil diambil",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Update(body Actors, ID string) (*AllResponse, error) {
	actors, err := c.useCase.Update(body, ID)
	if err != nil {
		return nil, err
	}
	res := ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}

	allres := &AllResponse{
		Message: "Data Admin berhasil diupdate",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Delete(ID string) (*AllResponse, error) {
	actors, err := c.useCase.Delete(ID)
	if err != nil {
		return nil, err
	}

	res := ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}

	allres := &AllResponse{
		Message: "Data Admin berhasil dihapus",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}
func (c Controller) Approval(body Approval, ID string) (*Approval, error) {
	approvalResponse, err := c.useCase.Approval(body, ID)
	if err != nil {
		return nil, err
	}

	return approvalResponse, nil
}
func (c Controller) Activate(body Activate, ID string) (*Activate, error) {
	activateResponse, err := c.useCase.Activate(body, ID)
	if err != nil {
		return nil, err
	}

	return activateResponse, nil
}
func (c Controller) Login(username string, password string) (*AllResponse, error) {
	actors, err := c.useCase.Login(username, password)
	if err != nil {
		return nil, err
	}

	res := ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}
	token, _ := GenerateJWT(res, "rahasia")
	allres := &AllResponse{
		Message: "Anda berhasil login",
		Token:   token,
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}
