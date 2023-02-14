package users

type Store interface {
	create(r UserRequest) (int, error)
	findById(id int) (*UserResponse, error)
}

type Service interface {
	Create(r UserRequest) (*UserResponse, error)
	FindById(userId int) (*UserResponse, error)
}

type UserRequest struct {
	Email, Name, Password string `json:""`
}

type UserResponse struct {
	Id          int
	Email, Name string
}

var _ Store = &UserStore{}
var _ Service = &UserService{}
