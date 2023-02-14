package users

type UserServiceStub struct {
}

func (u UserServiceStub) Create(r UserRequest) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceStub) FindById(userId int) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ Service = &UserServiceStub{}
