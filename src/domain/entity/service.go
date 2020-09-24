package entity

type Service interface {
	GetByID(id int64) (*Entity, error)
	CreateEntity(e Entity) error
}

type Repository interface {
	GetByID(id int64) (*Entity, error)
	CreateEntity(e Entity) error
}

func NewService(repo Repository) Service {
	return &service{
		db: repo,
	}
}

type service struct {
	db Repository
}

func (s *service) GetByID(id int64) (*Entity, error) {
	return s.db.GetByID(id)
}

func (s *service) CreateEntity(e Entity) error {
	return s.db.CreateEntity(e)
}
