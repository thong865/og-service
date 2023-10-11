package validation

import (
	"github.com/go-playground/validator/v10"
)

// type ReqBody struct {
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// }
type ReqBody struct {
	Username string `form:"Username" json:"Username" xml:"Username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginValidateStruct(req *ReqBody) error {
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return err
	}
	return nil
}
