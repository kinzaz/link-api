package jwt_test

import (
	"httpServer/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "test@test.ru"
	jwtService := jwt.NewJWT("7poRA+HPAzOmsoSI7iqA44LLe1EQpwyUOgV7BE5q4zQ=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
