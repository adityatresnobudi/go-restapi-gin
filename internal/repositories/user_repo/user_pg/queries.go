package user_pg

const CREATE_USER_QUERY = `
	INSERT INTO users (username, password) VALUES ($1, $2)
`

const GET_BY_USERNAME_QUERY = `
	SELECT id, username, password from users where username = $1
`

const GET_BY_ID_QUERY = `
	SELECT id, username, password from users where id = $1
`
