package account

import (
	"crud_api/dto"
	"crud_api/entities"
	"crud_api/utility"
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

func (c Controller) Create(body *CreateRequest) (*dto.AllActorResponse, error) {
	actors := entities.Actors{
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
	res := dto.ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}
	allres := &dto.AllActorResponse{
		Message: "Data berhasil dibuat",
	}
	allres.Data = append(allres.Data, res)

	return allres, nil
}

// Get All Data
func (c Controller) Read() (*dto.AllActorResponse, error) {
	actors, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}
	res := &dto.AllActorResponse{}
	for _, actor := range actors {
		c := dto.ActorDataResponse{
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

func (c Controller) ReadID(ID string) (*dto.AllActorResponse, error) {
	actors, err := c.useCase.ReadID(ID)
	if err != nil {
		return nil, err
	}

	if len(actors) == 0 {
		return nil, fmt.Errorf("entities.Actors not found")
	}

	res := dto.ActorDataResponse{
		ID:         uint8(actors[0].ID),
		Username:   actors[0].Username,
		Password:   actors[0].Password,
		RoleID:     actors[0].RoleID,
		IsVerified: actors[0].IsVerified,
		IsActive:   actors[0].IsActive,
	}
	allres := &dto.AllActorResponse{
		Message: "Data Admin berhasil diambil",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Update(body entities.Actors, ID string) (*dto.AllActorResponse, error) {
	actors, err := c.useCase.Update(body, ID)
	if err != nil {
		return nil, err
	}
	res := dto.ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}

	allres := &dto.AllActorResponse{
		Message: "Data Admin berhasil diupdate",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Delete(ID string) (*dto.AllActorResponse, error) {
	actors, err := c.useCase.Delete(ID)
	if err != nil {
		return nil, err
	}

	res := dto.ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}

	allres := &dto.AllActorResponse{
		Message: "Data Admin berhasil dihapus",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}
func (c Controller) Approval(body entities.Approval, ID string) (*entities.Approval, error) {
	approvalResponse, err := c.useCase.Approval(body, ID)
	if err != nil {
		return nil, err
	}

	return approvalResponse, nil
}
func (c Controller) Activate(body entities.Activate, ID string) (*entities.Activate, error) {
	activateResponse, err := c.useCase.Activate(body, ID)
	if err != nil {
		return nil, err
	}

	return activateResponse, nil
}
func (c Controller) Login(username string, password string) (*dto.AllActorResponse, error) {
	actors, err := c.useCase.Login(username, password)
	if err != nil {
		return nil, err
	}

	res := dto.ActorDataResponse{
		ID:         uint8(actors.ID),
		Username:   actors.Username,
		Password:   actors.Password,
		RoleID:     actors.RoleID,
		IsVerified: actors.IsVerified,
		IsActive:   actors.IsActive,
	}
	token, _ := utility.GenerateJWT(res, "rahasia")
	allres := &dto.AllActorResponse{
		Message: "Anda berhasil login",
		Token:   token,
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}
