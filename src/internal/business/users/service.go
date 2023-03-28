package users

type UserService struct {
	store Store
}

func NewService(store Store) *UserService {
	return &UserService{store}
}

func (u *UserService) Create(r UserRequest) (*UserResponse, error) {
	id, err := u.store.create(r)

	if err != nil {
		return nil, err
	}

	return &UserResponse{id, r.Email, r.Name}, nil
}

func (u *UserService) FindById(userId int) (*UserResponse, error) {
	return u.store.findById(userId)
}
