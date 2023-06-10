package account

import (
	"crud_api/entities"
	"crud_api/utility"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]entities.Actors, error) {
	var actors []entities.Actors
	err := r.db.Find(&actors).Error
	return actors, err
}

func (r repository) Save(actors *entities.Actors) error {
	pass, _ := utility.HashPassword(actors.Password)
	actors.Password = pass
	err := r.db.Create(actors).Error
	fmt.Println("id admin : ", actors.ID)
	var registerData = entities.Register{
		AdminID:      uint8(actors.ID),
		SuperAdminID: uint8(1),
		Status:       "Not Verified",
	}
	r.db.Create(registerData)

	return err

}
func (r repository) FindByID(ID string) ([]entities.Actors, error) {
	var actors []entities.Actors
	err := r.db.First(&actors, ID).Error
	return actors, err
}

func (r repository) UpdateByID(body entities.Actors, ID string) (*entities.Actors, error) {
	var actors entities.Actors
	err := r.db.First(&actors, ID).Error
	pass, _ := utility.HashPassword(body.Password)
	actors.Username = body.Username
	actors.Password = pass
	actors.RoleID = body.RoleID
	actors.IsVerified = body.IsVerified
	actors.IsActive = body.IsActive

	update_query := r.db.Save(&actors).Error
	if update_query != nil {
		return nil, update_query
	}
	return &actors, err
}

func (r repository) DeleteByID(ID string) (*entities.Actors, error) {
	var actors entities.Actors
	err := r.db.First(&actors, ID).Error

	delete_query := r.db.Delete(&actors).Error
	//delete_query := r.db.Raw("DELETE * FROM users where id = 12").Error
	if delete_query != nil {
		return nil, delete_query
	}
	return &actors, err
}

func (r repository) Approval(body entities.Approval, ID string) (*entities.Approval, error) {
	var approve entities.Approval

	err := r.db.First(&approve, ID).Error
	approve.IsVerified = body.IsVerified

	updateQuery := r.db.Save(&approve).Error
	if updateQuery != nil {
		return nil, updateQuery
	}

	// Update the register table
	var registerData entities.RegisterStatus
	registerData.AdminID = uint8(body.ID)
	if body.IsVerified == "true" {
		registerData.Status = "Verified"
	} else {
		registerData.Status = "Not Verified"
	}
	fmt.Println(registerData)
	updateRegisterQuery := r.db.Model(registerData).Where("admin_id = ?", ID).Update("status", registerData.Status).Error
	if updateRegisterQuery != nil {
		return nil, updateRegisterQuery
	}

	return &approve, err
}
func (r repository) Activate(body entities.Activate, ID string) (*entities.Activate, error) {
	var activate entities.Activate

	err := r.db.First(&activate, ID).Error
	activate.IsActive = body.IsActive

	update_query := r.db.Save(&activate).Error
	if update_query != nil {
		return nil, update_query
	}
	// Update the register table
	var registerData entities.RegisterStatus
	registerData.AdminID = uint8(body.ID)
	if body.IsActive == "true" {
		registerData.Status = "Active"
	} else {
		registerData.Status = "Inactive"
	}
	fmt.Println(registerData)
	updateRegisterQuery := r.db.Model(registerData).Where("admin_id = ?", ID).Update("status", registerData.Status).Error
	if updateRegisterQuery != nil {
		return nil, updateRegisterQuery
	}
	return &activate, err
}
func (r repository) Login(username, password string) (*entities.Actors, error) {
	var actor entities.Actors

	// Retrieve the actor based on the username
	err := r.db.Where("username = ?", username).First(&actor).Error
	if err != nil {
		return &actor, err
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(actor.Password), []byte(password))
	if err != nil {
		return &actor, err
	}

	// Password is correct, return the actor
	return &actor, nil
}

type Repository interface {
	GetAll() ([]entities.Actors, error)
	Save(actors *entities.Actors) error
	FindByID(ID string) ([]entities.Actors, error)
	UpdateByID(body entities.Actors, ID string) (*entities.Actors, error)
	DeleteByID(ID string) (*entities.Actors, error)
	Approval(body entities.Approval, ID string) (*entities.Approval, error)
	Activate(body entities.Activate, ID string) (*entities.Activate, error)
	Login(username, password string) (*entities.Actors, error)
}
