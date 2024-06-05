package repositories

const (
	qCreateOneUser = `
	INSERT INTO users (email, username ,password) VALUES
	($1, $2, $3) RETURNING id
	`

	qFindUserByUsername = `
	SELECT id, username, email, password FROM users
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
