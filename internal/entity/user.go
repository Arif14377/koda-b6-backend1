package entity

import "time"

type Users struct {
	Id         int       `json:"id"`
	FullName   string    `json:"fullName"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Photo      string    `json:"photo"`
	Role       string    `json:"role"`
	Created_At time.Time `json:"created_at"`
	Created_By *string   `json:"created_by"`
	Updated_At time.Time `json:"updated_at"`
	Updated_By *string   `json:"updated_by"`
}

type RequestUserRegister struct {
	FullName string `json:"fullName" example:"John Doe"`
	Email    string `json:"email" example:"johndoe@mail.com"`
	Password string `json:"password" example:"Pass1234"`
}

type RequestUserLogin struct {
	Email    string `json:"email" example:"abc@mail.com"`
	Password string `json:"password" example:"password"`
}

type RequestUserEdit struct {
	FullName string `json:"fullName" example:"John Doe"`
	Email    string `json:"email" example:"johndoe@mail.com"`
	Password string `json:"password" example:"Pass1234"`
	Phone    string `json:"phone" example:"081234567890"`
	Address  string `json:"address" example:"Washington DC, United State"`
	Photo    string `json:"photo" example:"https://images.pexels.com/photos/9775483/pexels-photo-9775483.jpeg"`
}

type ResponseUserShow struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Photo    string `json:"photo"`
	Role     string `json:"role"`
}
