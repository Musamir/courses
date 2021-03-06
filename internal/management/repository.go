package management

import "courses/internal/management/models"

// UserRepository ...
type UserRepository interface {
	SetUserInfo(request *models.UsersInfo) error
	GetUserInfo(id int) (*models.UsersInfo, error)
	GetUserPhoto(id int) (*string, error)
	GetSexList() (*[]string, error)
}
