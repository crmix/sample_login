package model


type PostNewUser struct {
	Name        string  `json: "name"`
	Password       string  `json: "password"`
	UserName    string  `json: "username"`
}

//type Users struct {
//	Id       int      `json: "id"`
//	Name     string   `json: "name"`
//	Phone    string   `json: "phone"`
//    Username string   `json: "username"`}

type LogStruct struct {
    Username string `json: "username"`
    Password string `json: "password"`
}