package customers

import "crud_api/entities"

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) Read() ([]entities.Customers, error) {
	return u.repo.GetAll()
}
func (u UseCase) Create(customer *entities.Customers) error {
	return u.repo.Save(customer)
}
func (u UseCase) ReadID(ID string) ([]entities.Customers, error) {
	return u.repo.FindByID(ID)
}
func (u UseCase) Update(body entities.Customers, ID string) (*entities.Customers, error) {
	return u.repo.UpdateByID(body, ID)
}
func (u UseCase) Delete(ID string) (*entities.Customers, error) {
	return u.repo.DeleteByID(ID)
}
