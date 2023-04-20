package users

type UserService struct {
	store Store
}

func NewService(store Store) *UserService {
	return &UserService{store}
}

func (u *UserService) Create(r UserRequest) (*BasicUser, error) {
	password, _ := HashPassword(r.Password)

	id, err := u.store.create(r.Name, r.Email, password)

	if err != nil {
		return nil, err
	}

	return &BasicUser{id, r.Email, r.Name}, nil
}

func (u *UserService) FindById(userId int) (*BasicUser, error) {
	return u.store.findById(userId)
}

func (u *UserService) FindPasswordByEmail(email string) (*UserIdAndPwd, error) {
	return u.store.findPasswordByEmail(email)
}

func (u *UserService) UserExists(email string) (bool, error) {
	return u.store.userExists(email)
}
