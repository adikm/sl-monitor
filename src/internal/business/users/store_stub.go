package users

import "errors"

var users []stubUser
var id = 0

type stubUser struct {
	*BasicUser
	password string
}
type UserStoreStub struct {
}

func (u UserStoreStub) create(name string, email string, password string) (int, error) {
	id++
	user := stubUser{&BasicUser{id, email, name}, password}
	users = append(users, user)
	return user.Id, nil
}

func (u UserStoreStub) findById(id int) (*BasicUser, error) {
	if id == 0 {
		return &BasicUser{0, "1@1.com", "Mister User"}, nil
	}
	for _, user := range users {
		if user.Id == id {
			return user.BasicUser, nil
		}
	}
	return nil, errors.New("no user")
}

func (u UserStoreStub) findPasswordByEmail(email string) (*UserIdAndPwd, error) {
	for _, user := range users {
		if user.Email == email {
			return &UserIdAndPwd{user.Id, user.password}, nil
		}
	}
	return nil, errors.New("no user")
}

func (u UserStoreStub) userExists(email string) (bool, error) {
	for _, user := range users {
		if user.Email == email {
			return true, nil
		}
	}
	return false, errors.New("no user")
}

var _ Store = &UserStoreStub{}
