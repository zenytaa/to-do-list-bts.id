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
	INSERT INTO checklists (name) VALUES
	($1) RETURNING id
	`

	qFindAllChecklist = `
	SELECT id, name from checklists
	WHERE deleted_at IS NULL
	`

	qFindOneChecklistById = `
	SELECT id, name from checklists
	WHERE id = $1 AND deleted_at IS NULL
	`

	qDeleteOneChecklist = `
	UPDATE checklists SET
	deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;
	`

	qCreateItem = `
	INSERT INTO items(item_name, checklist_id) VALUES
	($1, $2) RETURNING id;
	`
)
