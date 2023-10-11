package v1

import (
	oat "MS/config/auth"
	"MS/models"
	v "MS/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SVLOTLOGIN(c *gin.Context) {
	//
	var _req v.ReqBody
	if err := c.ShouldBindJSON(&_req); err != nil {
		v.LoginValidateStruct(&_req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := oat.Attemp(_req.Username, _req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"status": "success"})
	// return
	// resp, err := http.Get("http://10.0.32.99:8889/encode")
	// if err != nil {
	// 	log.Printf("Request Failed: %s", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(c.Request)
	// if err != nil {
	// 	log.Printf("Reading body failed: %s", err)
	// 	return
	// }
	// // Log the request body
	// bodyString := string(body)
	// log.Print(bodyString)
}

type RegisterInput struct {
	MobileNo string `json:"mobile_no" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.MobileNo = input.MobileNo
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
