package model

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type RegisterUserID struct {
	ID int `json:"id"`
}
