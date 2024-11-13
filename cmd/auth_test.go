package main

import (
	"bytes"
	"encoding/json"
	"httpServer/internal/auth"
	"httpServer/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "test@user.ru",
		Password: "$2a$10$fjo/.eP0sp7vIfrhwu5JH.vvQKVopvYb09vSh.qYprI5zbjbPBZOe",
		Name:     "Test",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().Where("email = ?", "test@user.ru").Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@user.ru",
		Password: "123",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token empty")
	}
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@ru3.ru",
		Password: "1234",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, response.StatusCode)
	}
	removeData(db)

}
