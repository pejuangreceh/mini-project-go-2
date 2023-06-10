package account

import "crud_api/entities"

type UseCase struct {
	repo Repository
}

func NewUseCase(repo *repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) Read() ([]entities.Actors, error) {
	return u.repo.GetAll()
}

func (u UseCase) Create(actor *entities.Actors) error {
	return u.repo.Save(actor)
}
func (u UseCase) ReadID(ID string) ([]entities.Actors, error) {
	return u.repo.FindByID(ID)
}

func (u UseCase) Update(body entities.Actors, ID string) (*entities.Actors, error) {
	return u.repo.UpdateByID(body, ID)
}

func (u UseCase) Delete(ID string) (*entities.Actors, error) {
	return u.repo.DeleteByID(ID)
}

func (u UseCase) Approval(body entities.Approval, ID string) (*entities.Approval, error) {
	return u.repo.Approval(body, ID)
}
func (u UseCase) Activate(body entities.Activate, ID string) (*entities.Activate, error) {
	return u.repo.Activate(body, ID)
}
func (u UseCase) Login(username string, password string) (*entities.Actors, error) {
	return u.repo.Login(username, password)
}
