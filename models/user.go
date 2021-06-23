package models

type User struct {
	Model
	Firstname string `json:"firstname`
	Lastname  string `json: "lastname"`
}

func (u *User) TableName() string {
	return "user"
}
