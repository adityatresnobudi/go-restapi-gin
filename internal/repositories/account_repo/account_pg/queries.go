package account_pg

const GET_ALL_ACCOUNTS = `
	SELECT id, account_number, account_holder, balance, created_at, updated_at
	FROM accounts
`

const GET_ACCOUNT_BY_ID = `
	SELECT id, account_number, account_holder, balance, created_at, updated_at
	FROM accounts WHERE id = $1
`

const GET_ACCOUNT_BY_ACCOUNTNUM = `
	SELECT id, account_number, account_holder, balance, created_at, updated_at
	FROM accounts WHERE account_number = $1
`

const INSERT_ACCOUNT = `
	INSERT INTO accounts (account_number, account_holder, balance) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, account_number, account_holder, balance, created_at, updated_at
`

const UPDATE_ACCOUNT = `
	UPDATE accounts
	SET account_holder = $1, balance = $2
	WHERE id = $3
	RETURNING id, account_number, account_holder, balance, created_at, updated_at
`

const DELETE_ACCOUNT = `
	DELETE FROM accounts
	WHERE id = $1
`

const UPDATE_BALANCE = `
	UPDATE accounts
	SET balance = $1
	WHERE id = $2
`

const INSERT_TRANSACTION = `
	INSERT INTO transactions (account_id_from, account_id_to, amount) 
	VALUES ($1, $2, $3)
	RETURNING id, account_id_from, account_id_to, amount, created_at, updated_at
`