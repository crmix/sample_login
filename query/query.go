package query

var GET_NAME= `
  SELECT *FROM users
`

var POST_NEWUSER=`
  insert into users (name, username, phone) values ($1,$2, $3)
	returning
		name,
		username,
		phone
		
`