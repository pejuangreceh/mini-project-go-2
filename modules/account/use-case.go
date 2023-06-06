package account

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) Read() ([]Actors, error) {
	return u.repo.GetAll()
}

func (u UseCase) Create(actor *Actors) error {
	return u.repo.Save(actor)
}
func (u UseCase) ReadID(ID string) ([]Actors, error) {
	return u.repo.FindByID(ID)
}

func (u UseCase) Update(body Actors, ID string) (*Actors, error) {
	return u.repo.UpdateByID(body, ID)
}

func (u UseCase) Delete(ID string) (*Actors, error) {
	return u.repo.DeleteByID(ID)
}