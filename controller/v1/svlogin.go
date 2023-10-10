package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SVLOTLOGIN(c *gin.Context) {
	v := validator.New()
	v.RegisterValidation()
	// var json Login
	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"status": "success"})
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
