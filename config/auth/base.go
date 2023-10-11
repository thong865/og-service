package auth

import (
	hs "MS/config/hash"
	"MS/models"
	"fmt"
)

func VerifyPassword(password string, hashedPassword string) error {
	p, err := hs.Encrypt(password)
	if err != nil {
		return err
	}

	if string(hashedPassword) == string(p) {
		fmt.Println("%v", p)
	}

	return err
}

func Attemp(username string, password string) (string, error) {

	var err error

	u := models.User{}
	err = models.GetDB().Model(models.User{}).Where("mobile_no = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}
	token, err := GenarateAccessToken(string(u.ID))
	if err != nil {
		return "", err
	}

	return token, nil

}
