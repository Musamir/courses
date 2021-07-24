package http

import (
	"courses/internal/management"
	"courses/internal/management/models"

	"github.com/gin-gonic/gin"
)

// Handler-used for data interaction over the http protocol
type UserHandler struct {
	// userImp - user repository implementation
	Impl management.UserUseCase
}

// NewHandler creates a new Handler structure
func NewUserHandler(Impl management.UserUseCase) *UserHandler {

	return &UserHandler{
		Impl: Impl,
	}
}

// signRequest - structure that is sent to sign-in/up url
type userInfo struct {
	Login              string `json:"login"`              // login - user login
	FullName           string `json:"fullName"`           // fullName - user fullName
	Sex                string `json:"sex"`                // sex - user sex
	DateOfBirth        string `json:"dateOfBirth"`        // dateOfBirth - user date of birth
	ChangeProfilePhoto string `json:"changeProfilePhoto"` // changeProfilePhoto = {"yes" - change, "no" - pass}
	ProfilePhoto       string `json:"profilephoto"`       // profilephoto - user profile photo file. If you don't want to change profile photo send current profilephoto filename
}

// RegisterUser registers userinfo handlers
func (UserHandler *UserHandler) RegisterUser(groupname string, router *gin.RouterGroup) {

	userEndpoints := router.Group(groupname)
	{
		userEndpoints.GET("/test", test)
	}
}

const (
	internalServerError string = "Oops something went wrong"
	userInfoWasSet      string = "The user information was set"
)

func toSetUserInfo(u *models.UsersInfo) *userInfo {
	return &userInfo{}
}

func toModelsUser(u *userInfo, id int) *models.UsersInfo {
	return &models.UsersInfo{
		ID:           id,
		Login:        u.Login,
		FullName:     u.FullName,
		Sex:          u.Sex,
		DateOfBirth:  u.DateOfBirth,
		ProfilePhoto: u.ProfilePhoto,
	}
}

func geId(c *gin.Context) (id int, exist bool) {
	idstr, exist := c.Get("user_id")
	id = 0
	if exist {
		id = idstr.(int)
	}
	return
}

// @Summary test
// @Security ApiKeyAuth
// @Tags safe
// @Description test is used to test your jwt token. It return "hello" if your token is suitable
// @Produce  json
// @Success 200 {string} message "hello"
// @Failure 400 {string} error "error"
// @Failure 401 {string} error "Unauthorized!!"
// @Failure 500 {string} error "something went wrong"
// @Router /safe/business/test [get]
func test(c *gin.Context) {
	c.JSON(200, "hello")
}
