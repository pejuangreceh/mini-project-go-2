package customers

import (
	"crud_api/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]entities.Customers, error) {
	var customers []entities.Customers
	err := r.db.Find(&customers).Error
	return customers, err
}
func (r repository) Save(customers *entities.Customers) error {
	return r.db.Create(customers).Error
}
func (r repository) FindByID(ID string) ([]entities.Customers, error) {
	var customers []entities.Customers
	err := r.db.First(&customers, ID).Error
	return customers, err
}
func (r repository) UpdateByID(body entities.Customers, ID string) (*entities.Customers, error) {
	var customers entities.Customers
	err := r.db.First(&customers, ID).Error
	customers.FirstName = body.FirstName
	customers.LastName = body.LastName
	customers.Email = body.Email
	customers.Avatar = body.Avatar

	update_query := r.db.Save(&customers).Error
	if update_query != nil {
		return nil, update_query
	}
	return &customers, err
}

func (r repository) DeleteByID(ID string) (*entities.Customers, error) {
	var customers entities.Customers
	err := r.db.First(&customers, ID).Error

	delete_query := r.db.Delete(&customers).Error
	//delete_query := r.db.Raw("DELETE * FROM users where id = 12").Error
	if delete_query != nil {
		return nil, delete_query
	}
	return &customers, err
}

type Repository interface {
	GetAll() ([]entities.Customers, error)
	Save(customers *entities.Customers) error
	FindByID(ID string) ([]entities.Customers, error)
	UpdateByID(body entities.Customers, ID string) (*entities.Customers, error)
	DeleteByID(ID string) (*entities.Customers, error)
}
