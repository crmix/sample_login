package service

import (
	"fmt"
	"log"
	"login/model"
	"login/storage"
)

type userServiceImpl struct {
	storage storage.User1
	logger  log.Logger

}

func NewUserService(user1 storage.User1, logger log.Logger) UserService {
	return &userServiceImpl{
		storage: user1,
		logger: logger,
	}
}

type UserService interface{
	GetUsersName (users []model.User) (interface{}, error)
}


func (c *userServiceImpl) GetUsersName (user []model.User)(interface{}, error) {
	res, err :=c.storage.GetUsersName(user)
	if err !=nil {
		log.Println(err)
		return nil, fmt.Errorf("error in getusername %w", err) 
	}
  return res, nil

}

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
	PostNewUser(posts []model.PostNewUser)(interface{}, error) 
}

func (s *postServiceImpl) PostNewUser (posts []model.PostNewUser)(interface{}, error){
	res, err :=s.storage.PostNewUser(posts)
	if err !=nil {
		log.Println(err)
		return nil, fmt.Errorf("error in posting %w", err)
	}
	return res, nil
}