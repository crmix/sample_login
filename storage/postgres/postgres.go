package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"login/model"
	"login/query"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetUsersName(users []model.User) (interface{} , error) {
	row, err := u.db.Query(query.GET_NAME)
	if err != nil{
		log.Println(err)
		return nil, fmt.Errorf("error in GetUsersName: %w", err) 
	}
	var Users []model.User

	for row.Next() {
		var user model.User
		row.Scan(
			&user.Id,
			&user.Name,
			&user.Phone,
			&user.Username)

		Users= append(Users, user)
	}
	//fmt.Println(feruz)
	return Users, nil
}