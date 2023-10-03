package postgres

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"log"
	"login/model"
	"login/query"
)

type postRepo struct {
	db *sql.DB
}

func PostUserRepo(db *sql.DB) *postRepo {
	return &postRepo{
		db: db,
	}
}
func (r *postRepo) CreateUser(newposts model.PostNewUser)([]model.PostNewUser, error) {
	
	row, err :=r.db.Query(query.POST_NEWUSER, newposts.Name, newposts.UserName, newposts.Password)
	
	if err != nil{
		log.Println(err)
		return nil, fmt.Errorf("error in PostNewUser: %w", err)}
		
		var Data []model.PostNewUser
		   var newpost model.PostNewUser
		row.Scan(
			&newpost.Name,
			&newpost.Password,
			&newpost.UserName)
		//Posts = append(Posts,newpost )	
		fmt.Println(&newpost)
		Data=append(Data, newpost)
		return Data, nil
	  }

	func (r *postRepo) LoginUser(login model.LogStruct)([]model.LogStruct, error) {
	
	row, err :=r.db.Query(query.LOGIN_QUERY, login.Username, login.Password)
	
	if err != nil{
		log.Println(err)
		return nil, fmt.Errorf("error in LoginUser: %w", err)}

		fmt.Println(row)
		
       var Info []model.LogStruct

		var checklogin model.LogStruct
		
		row.Scan(
			&checklogin.Password,
			&checklogin.Username)
		//Posts = append(Posts,newpost )	
		fmt.Println(&checklogin)
		Info=append(Info, checklogin)

		return Info, nil
	  }
