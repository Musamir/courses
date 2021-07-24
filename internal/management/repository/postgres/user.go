package postgres

import (
	"courses/internal/management/models"
	"fmt"
)

type user struct {
	ID       int
	Login    string
	Password string
}

// func toPostgresUser(u *models.User) *user {
// 	return &user{
// 		ID:       u.ID,
// 		Login:    u.Login,
// 		Password: u.Password,
// 	}
// }

// func toModel(u *user) *models.User {
// 	return &models.User{
// 		ID:       u.ID,
// 		Login:    u.Login,
// 		Password: u.Password,
// 	}
// }

// UserRepository ...
type UserRepository struct {
	db *Database
}

// NewUserRepository ...
func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SetUserInfo sets the user information
func (user *UserRepository) SetUserInfo(request *models.UsersInfo) error {
	fmt.Println("(user *UserRepository) SetUserInfo start")
	defer fmt.Println("(user *UserRepository) SetUserInfo end")

	return nil
}

//GetUserInfo gets the user information
func (user *UserRepository) GetUserInfo(id int) (*models.UsersInfo, error) {
	fmt.Println(" (user *UserRepository) GetUserInfo start")
	defer fmt.Println(" (user *UserRepository) GetUserInfo end")

	var userInfo models.UsersInfo

	return &userInfo, nil
}

//GetUserPhoto gets the user profile photo
func (user *UserRepository) GetUserPhoto(id int) (*string, error) {
	fmt.Println(" (user *UserRepository) GetUserPhoto start")
	defer fmt.Println(" (user *UserRepository) GetUserPhoto end")
	var profilephoto string

	return &profilephoto, nil
}

//GetSexList gets sex list
func (user *UserRepository) GetSexList() (*[]string, error) {
	fmt.Println(" (user *UserRepository) GetSexList start")
	defer fmt.Println("(user *UserRepository) GetSexList() end")

	// Using tmp variable for reading
	var sexList []string

	return &sexList, nil
}
