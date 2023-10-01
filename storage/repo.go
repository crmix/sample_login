package storage

import (
	"database/sql"
	"login/model"
	"login/storage/postgres"
)

type User1 interface {
	GetUsersName(users []model.User) (interface{}, error)
}

func NewUserRepo(db *sql.DB) User1 {
	return postgres.NewUserRepo(db)
}