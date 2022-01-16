package db

type stmtTyp int

const (
	InsertItem stmtTyp = iota
	ReadItem
	ReadAllItems
	Upsert
	UpdateItem
	DeleteItem
)

var stmts = map[stmtTyp]string{
	ReadItem:     `SELECT * FROM items WHERE id = ?`,
	ReadAllItems: `SELECT * FROM items`,
	InsertItem: `
		INSERT INTO items VALUES (
		?,
		?,
		?
		)`,
	// InsertItem: `
	// 	INSERT INTO items (
	// 		title,
	// 		created_at
	// 	) VALUES (?, ?)
	// 	RETURNING *`,
	UpdateItem: `UPDATE INTO items SET VALUES 
			title = ?,
			created_at = ?,
			WHERE id = ? RETURNING *`,
	DeleteItem: `DELETE FROM items WHERE id = ?`,
}

type migTyp int

const (
	CreateItems migTyp = iota
	DropItems
)

var migrate = map[migTyp]string{
	CreateItems: `
	CREATE TABLE IF NOT EXISTS items (
		"id"	BLOB,
		"title"	TEXT NOT NULL,
		"created_at"	DATETIME,
		PRIMARY KEY("id")
	);`,
	DropItems: `DROP IF EXISTS items`,
}
