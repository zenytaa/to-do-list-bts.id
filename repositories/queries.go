package repositories

const (
	qCreateOneUser = `
	INSERT INTO users (email, username ,password) VALUES
	($1, $2, $3) RETURNING id
	`

	qFindUserByUsername = `
		SELECT id, name, email, password FROM users
		WHERE name = $1 AND deleted_at IS NULL;
	`
)
