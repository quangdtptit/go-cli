package user

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	GetAll() ([]User, error)
}

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) GetAll() ([]User, error) {
	return []User{}, nil
}
