package users

import (
	"testing"
)

func TestUsersService_Create(t *testing.T) {
	s := UserService{&UserStoreStub{}}

	r := UserRequest{
		"1@1.com",
		"Mr. Adrian",
		"pwd123",
	}

	got, _ := s.Create(r)

	if got.Id < 1 {
		t.Errorf("Create() got = %v, want >= 1", got.Id)
	}

	if got.Name != r.Name {
		t.Errorf("Create() got = %v, want %v", got.Name, r.Name)
	}

	if got.Email != r.Email {
		t.Errorf("Create() got = %v, want %v", got.Email, r.Email)
	}

}

func TestUsersService_FindById(t *testing.T) {
	s := UserService{&UserStoreStub{}}

	r := UserRequest{
		"1@1.com",
		"Mr. Adrian",
		"pwd123",
	}
	u, _ := s.Create(r)

	got, _ := s.FindById(u.Id)
	if got.Id < 1 {

		t.Errorf("FindById() got = %v, want = %v", got.Id, u.Id)
	}

	if got.Name != r.Name {
		t.Errorf("FindById() got = %v, want %v", got.Name, r.Name)
	}

	if got.Email != r.Email {
		t.Errorf("FindById() got = %v, want %v", got.Email, r.Email)
	}

}

func TestUsersService_FindPasswordByEmail(t *testing.T) {
	s := UserService{&UserStoreStub{}}

	r := UserRequest{
		"1@1.com",
		"Mr. Adrian",
		"pwd123",
	}
	u, _ := s.Create(r)

	got, _ := s.FindPasswordByEmail(r.Email)
	if got.Id < 1 {

		t.Errorf("FindPasswordByEmail() got = %v, want = %v", got.Id, u.Id)
	}

	if got.Pwd == r.Password {
		t.Errorf("FindPasswordByEmail() got = %v, want %v", got.Pwd, "hashed password")
	}
}

func TestUsersService_UserExists(t *testing.T) {
	s := UserService{&UserStoreStub{}}

	r := UserRequest{
		"1@1.com",
		"Mr. Adrian",
		"pwd123",
	}
	_, _ = s.Create(r)

	exists, _ := s.UserExists(r.Email)
	if !exists {

		t.Errorf("UserExists() got false")
	}

}
