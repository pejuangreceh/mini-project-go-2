package customers

import (
	"crud_api/entities"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetAll() ([]entities.Customers, error) {
	var customers []entities.Customers
	err := r.db.Find(&customers).Error
	return customers, err
}
func (r Repository) Save(customers *entities.Customers) error {
	return r.db.Create(customers).Error
}
func (r Repository) FindByID(ID string) ([]entities.Customers, error) {
	var customers []entities.Customers
	err := r.db.First(&customers, ID).Error
	return customers, err
}
func (r Repository) UpdateByID(body entities.Customers, ID string) (*entities.Customers, error) {
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

func (r Repository) DeleteByID(ID string) (*entities.Customers, error) {
	var customers entities.Customers
	err := r.db.First(&customers, ID).Error

	delete_query := r.db.Delete(&customers).Error
	//delete_query := r.db.Raw("DELETE * FROM users where id = 12").Error
	if delete_query != nil {
		return nil, delete_query
	}
	return &customers, err
}
