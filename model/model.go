package model

type User struct {
	Id int           `json: "id"`
	Name string      `json: "name"`
	Username string  `json: "username"`
	Phone string        `json: "phone"`
}

type PostNewUser struct {
	Name        string  `json:"name"`
	UserName    string  `json:"username"`
	PhoneNumber string  `json:"phonenumber"`
    
}