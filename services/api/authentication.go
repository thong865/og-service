package v1services

import (
	u "MS/apiHelpers"
	"MS/models"
	res "MS/resources/v1"
)

//UserService struct
type AuthenService struct {
	User models.User
}

//UserList function returns the list of users
func (us *AuthenService) UserList() map[string]interface{} {
	user := us.User

	userData := res.UserResponse{
		ID:    user.ID,
		Name:  "test",
		Email: "test@gmail.com",
	}
	response := u.Message(0, "This is from version 1 api")
	response["data"] = userData
	return response
}
