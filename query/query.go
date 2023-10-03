package query

var POST_NEWUSER=`
  insert into users (name, username, phone) values ($1,$2, $3)
	returning
		name,
		username,
		phone
		
`

var LOGIN_QUERY =`
  select username, password from users
`