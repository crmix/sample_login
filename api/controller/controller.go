package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	model"login/model"
	"login/service"
	"net/http"
    "time"
	jwt "github.com/dgrijalva/jwt-go"
)

type postApi struct {
	service service.PostService
}

func PostUserAPI(service service.PostService) PostAPI {
	return &postApi{
		service: service,
	}
}
type PostAPI interface{
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	
}

func(api *postApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	
	  body, _ :=ioutil.ReadAll(r.Body)

	var newUser model.PostNewUser

	json.Unmarshal(body, &newUser)
     fmt.Println(newUser)
	res, err :=api.service.CreateUser(newUser)
  if err != nil{
	log.Println(err)
	return
  }

  WriteSuccess(w, res)
}
   

type PostBody struct{
	Username, Password string
}

type  LogStruct struct{
	Username string
	Password string
	jwt.StandardClaims
}

var jwtKey =[]byte("hello_world")

func HomePage(w http.ResponseWriter, r *http.Request){
	
	token :=r.Header.Get("token")
    encoder :=json.NewEncoder(w)
	payload := &LogStruct{}

	tkn, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (
	   interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
  encoder.Encode([]int{1,2,3,4,5})
}



func Login(w http.ResponseWriter, r *http.Request){
	
    bStruct :=PostBody{}
	
	body, _ :=ioutil.ReadAll(r.Body)
  
	json.Unmarshal(body, &bStruct)
	
	if bStruct.Username == "feruz" && bStruct.Password== "feruz1" {
		
		expirationTime :=time.Now().Add(35 *time.Second)
    
    payload:=LogStruct{
	  Username: bStruct.Username,
	  Password: bStruct.Password,
	  StandardClaims: jwt.StandardClaims{
	  ExpiresAt: expirationTime.Unix(),
	 },
	}
  token :=jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

  tokenString, err := token.SignedString(jwtKey)

  if err!=nil{
	panic(err)
  }
    w.Write([]byte(tokenString))
 }else {
	w.Write([]byte("Wrong username or password"))
 }
}

