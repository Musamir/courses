package postgres

import (
	"courses/internal/auth"
	"courses/internal/auth/models"
	"fmt"
	"testing"
)

func TestUserRepository(t *testing.T) {
	url := "postgres://postgres:sql@localhost:5432/courses"

	db := NewDatabase(url)
	defer db.Close()

	userRepo := NewUserRepository(db)

	user := &models.User{
		Login:    "TEST",
		Password: "b1b3773a05c0ed0176787a4f1574ff0075f7521e",
	}

	id, status := userRepo.SingIn(user)
	if status != auth.DbOk {
		fmt.Println("no id")
		fmt.Println("the login or password is incorrect")
		t.Error("Expected status = 0, got ", status)
	}
	if id != 1 {
		t.Error("Expected id = 1, got ", id)
	}

	userId := 1
	oldPass := "b1b3773a05c0ed0176787a4f1574ff0075f7521e"
	newPass := "b1b3773a05c0ed0176787a4f1574ff0075f7521e"

	stat := userRepo.ChangePassword(userId, oldPass, newPass)
	if stat != auth.DbOk {
		t.Error("Expected status = 0, got ", stat)
	}
}
