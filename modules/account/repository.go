package account

import (
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetAll() ([]Actors, error) {
	var actors []Actors
	err := r.db.Find(&actors).Error
	return actors, err
}

func (r Repository) Save(actors *Actors) error {

	err := r.db.Create(actors).Error
	fmt.Println("id admin : ", actors.ID)
	var registerData = Register{
		AdminID:      uint8(actors.ID),
		SuperAdminID: uint8(1),
		Status:       "Not Verified",
	}
	r.db.Create(registerData)

	return err

}
func (r Repository) FindByID(ID string) ([]Actors, error) {
	var actors []Actors
	err := r.db.First(&actors, ID).Error
	return actors, err
}

func (r Repository) UpdateByID(body Actors, ID string) (*Actors, error) {
	var actors Actors
	err := r.db.First(&actors, ID).Error
	actors.Username = body.Username
	actors.Password = body.Password
	actors.RoleID = body.RoleID
	actors.IsVerified = body.IsVerified
	actors.IsActive = body.IsActive

	update_query := r.db.Save(&actors).Error
	if update_query != nil {
		return nil, update_query
	}
	return &actors, err
}

func (r Repository) DeleteByID(ID string) (*Actors, error) {
	var actors Actors
	err := r.db.First(&actors, ID).Error

	delete_query := r.db.Delete(&actors).Error
	//delete_query := r.db.Raw("DELETE * FROM users where id = 12").Error
	if delete_query != nil {
		return nil, delete_query
	}
	return &actors, err
}

func (r Repository) Approval(body Approval, ID string) (*Approval, error) {
	var approve Approval

	err := r.db.First(&approve, ID).Error
	approve.IsVerified = body.IsVerified

	update_query := r.db.Save(&approve).Error
	if update_query != nil {
		return nil, update_query
	}
	return &approve, err
}
