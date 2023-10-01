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
	return Users, nil
}
type postRepo struct {
	db *sql.DB
}

func PostUserRepo(db *sql.DB) *postRepo {
	return &postRepo{
		db: db,
	}
}
func (r *postRepo) PostNewUser([]model.PostNewUser)(interface{}, error) {
	
	row, err :=r.db.Query(query.POST_NEWUSER)
	
	if err != nil{
		log.Println(err)
		return nil, fmt.Errorf("error in PostNewUser: %w", err)}
		
		var newpost model.PostNewUser
      row.Scan(
			&newpost.Name,
			&newpost.PhoneNumber,
			&newpost.UserName)
		//Posts = append(Posts,newpost )	
		return newpost, nil
	  }
