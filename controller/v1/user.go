package v1

import (
	u "MS/apiHelpers"
	vserv "MS/services/api"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var userService vserv.UserService

	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		u.Respond(c.Writer, u.Message(1, "Invalid request"))
		return
	}

	//call service
	resp := userService.UserList()

	//return response using api helper
	u.Respond(c.Writer, resp)
}
