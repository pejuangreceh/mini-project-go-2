package customers

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

type CustomerDataResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type AllResponse struct {
	Message string                 `json:"message"`
	Data    []CustomerDataResponse `json:"data"`
}

// Get All Data
func (c Controller) Read() (*AllResponse, error) {
	customers, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}
	res := &AllResponse{}
	for _, customer := range customers {
		c := CustomerDataResponse{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		}
		res.Data = append(res.Data, c)
	}
	res.Message = "Data Sukses dimuat"
	return res, nil
}

func (c Controller) Create(body *CreateRequest) (*AllResponse, error) {
	customers := Customers{
		Model:     gorm.Model{},
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Avatar:    body.Avatar,
	}
	err := c.useCase.Create(&customers)
	if err != nil {
		return nil, err
	}
	res := CustomerDataResponse{
		ID:        customers.ID,
		FirstName: customers.FirstName,
		LastName:  customers.LastName,
		Email:     customers.Email,
		Avatar:    customers.Avatar,
	}
	allres := &AllResponse{
		Message: "Data berhasil diambil",
	}
	allres.Data = append(allres.Data, res)

	return allres, nil
}
func (c Controller) ReadID(ID string) (*AllResponse, error) {
	customers, err := c.useCase.ReadID(ID)
	if err != nil {
		return nil, err
	}

	if len(customers) == 0 {
		return nil, fmt.Errorf("Customers not found")
	}

	res := CustomerDataResponse{
		ID:        customers[0].ID,
		FirstName: customers[0].FirstName,
		LastName:  customers[0].LastName,
		Email:     customers[0].Email,
		Avatar:    customers[0].Avatar,
	}
	allres := &AllResponse{
		Message: "Data berhasil diambil",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Update(body Customers, ID string) (*AllResponse, error) {
	customers, err := c.useCase.Update(body, ID)
	if err != nil {
		return nil, err
	}
	res := CustomerDataResponse{
		ID:        customers.ID,
		FirstName: customers.FirstName,
		LastName:  customers.LastName,
		Email:     customers.Email,
		Avatar:    customers.Avatar,
	}

	allres := &AllResponse{
		Message: "Data berhasil diupdate",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}

func (c Controller) Delete(ID string) (*AllResponse, error) {
	customers, err := c.useCase.Delete(ID)
	if err != nil {
		return nil, err
	}

	res := CustomerDataResponse{
		ID:        customers.ID,
		FirstName: customers.FirstName,
		LastName:  customers.LastName,
		Email:     customers.Email,
		Avatar:    customers.Avatar,
	}

	allres := &AllResponse{
		Message: "Data berhasil dihapus",
	}
	allres.Data = append(allres.Data, res)
	return allres, nil
}
