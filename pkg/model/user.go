package model

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type RegisterUserID struct {
	ID int `json:"id"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenLogin struct {
	TokenJWT string `json:"tokenJWT"`
}

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
