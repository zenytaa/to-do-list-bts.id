package repositories

const (
	qCreateOneUser = `
	INSERT INTO users (username, email, password) VALUES
	($1, $2, $3) RETURNING id
	`

	qFindUserByUsername = `
	SELECT id, username, password FROM users
	WHERE username = $1 AND deleted_at IS NULL;
	`

	qCreateOneChecklist = `
	INSERT INTO checklists (item) VALUES
	($1) RETURNING id
	`

	qFindAllChecklist = `
	SELECT id, name from checklists
	WHERE deleted_at IS NULL
	`
)
