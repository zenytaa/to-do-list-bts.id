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

	qFindAllItem = `
	SELECT id, item_name, checklist_id, is_done FROM items
	WHERE checklist_id = $1 AND deleted_at IS NULL;
	`

	qFindOneItemById = `
	SELECT id, item_name, checklist_id, is_done FROM items
	WHERE checklist_id = $2 AND id = $1 AND deleted_at IS NULL;
	`

	qUpdateItemStatus = `
	UPDATE items SET
	is_done = CASE WHEN id = $1 THEN NOT is_done
	ELSE is_done END WHERE id = $1 AND checklist_id = $2 AND deleted_at IS NULL;	
	`

	qDeleteOneItem = `
	UPDATE items SET
	deleted_at = NOW()
	WHERE id = $1 AND checklist_id = $2 AND deleted_at IS NULL
	`

	qUpdateItemName = `
	UPDATE items SET
	item_name = $3
	WHERE id = $1 AND checklist_id = $2 AND deleted_at IS NULL
	`
)
