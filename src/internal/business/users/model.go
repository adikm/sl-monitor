package users

type Store interface {
	create(name string, mail string, password string) (int, error)
	findById(id int) (*BasicUser, error)

	findPasswordByEmail(email string) (*UserIdAndPwd, error)

	userExists(email string) (bool, error)
}

type Service interface {
	Create(r UserRequest) (*BasicUser, error)
	FindById(userId int) (*BasicUser, error)

	FindPasswordByEmail(email string) (*UserIdAndPwd, error)

	UserExists(email string) (bool, error)
}

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type BasicUser struct {
	Id          int
	Email, Name string
}

type UserIdAndPwd struct {
	Id  int
	Pwd string
}

var _ Store = &UserStore{}
var _ Service = &UserService{}
