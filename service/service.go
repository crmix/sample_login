package service

import (
	"fmt"
	"log"
	"login/model"
	"login/storage"
)

type postServiceImpl struct {
	storage storage.Post
	logger  log.Logger
}

func PostUserService(posts storage.Post, logger log.Logger) PostService {
	return &postServiceImpl{
		storage: posts,
		logger: logger,
	}
}
type PostService interface{
	CreateUser(posts model.PostNewUser)([]model.PostNewUser, error) 
	Login(login model.LogStruct)([]model.LogStruct, error)
}

func (s *postServiceImpl) CreateUser (posts model.PostNewUser)([]model.PostNewUser, error){
	res, err :=s.storage.CreateUser(posts)
	if err !=nil {
		log.Println(err)
		return nil, fmt.Errorf("error in posting %w", err)
	}
	return res, nil
}


func (p *postServiceImpl) LoginUser (login model.LogStruct)([]model.LogStruct, error){
	res, err :=p.storage.Login(login)
	if err !=nil {
		log.Println(err)
		return nil, fmt.Errorf("error in posting %w", err)
	}
	return res, nil
}
