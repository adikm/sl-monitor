package users

type UserServiceStub struct {
}

func (u UserServiceStub) Create(r UserRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

func (u UserServiceStub) FindById(userId int) (*UserResponse, error) {
	return &UserResponse{}, nil
}

var _ Service = &UserServiceStub{}
