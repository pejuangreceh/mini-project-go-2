package customers

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) Read() ([]Customers, error) {
	return u.repo.GetAll()
}
func (u UseCase) Create(customer *Customers) error {
	return u.repo.Save(customer)
}
func (u UseCase) ReadID(ID string) ([]Customers, error) {
	return u.repo.FindByID(ID)
}
func (u UseCase) Update(body Customers, ID string) (*Customers, error) {
	return u.repo.UpdateByID(body, ID)
}
func (u UseCase) Delete(ID string) (*Customers, error) {
	return u.repo.DeleteByID(ID)
}
