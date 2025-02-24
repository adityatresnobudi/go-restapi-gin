package transaction_pg

const INSERT_TRANSACTION = `
	INSERT INTO transactions (account_id_from, account_id_to, amount) 
	VALUES ($1, $2, $3)
	RETURNING id, account_id_from, account_id_to, amount, created_at, updated_at
`

const GET_ALL_TRANSACTIONS_BY_ID = `
	SELECT id, account_id_from, account_id_to, amount, created_at, updated_at
	FROM transactions WHERE account_id_from = $1 OR account_id_to = $2
`