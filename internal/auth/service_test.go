package auth_test

import (
	"httpServer/internal/auth"
	"httpServer/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}
func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSucces(t *testing.T) {
	const initialEmail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Test")

	if err != nil {
		t.Fatal(err)
	}

	if email != initialEmail {
		t.Fatalf("Email %s do not match %s", email, initialEmail)
	}
}
