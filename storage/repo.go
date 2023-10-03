package storage

import (
	"database/sql"
	"login/model"
	"login/storage/postgres"
)

type Post interface{
	CreateUser(post model.PostNewUser)([]model.PostNewUser, error)
	Login(login model.LogStruct)([]model.LogStruct, error)
}

func PostRepo(db *sql.DB) Post{
	return postgres.PostUserRepo(db)
}