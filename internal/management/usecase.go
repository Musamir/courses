package management

import "courses/internal/management/models"

// UserUseCase ...
type UserUseCase interface {
	SetUserInfo(userInfo *models.UsersInfo) error
	GetUserInfo(id int) (*models.UsersInfo, error)
	GetUserPhoto(id int) (*string, error)
	GetSexList() (*[]string, error)
}
