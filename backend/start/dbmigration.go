package start

const (
	createTableUsers = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL
	);`

	createTablePayments = `CREATE TABLE IF NOT EXISTS payments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME NOT NULL,
		merchant TEXT NOT NULL UNIQUE,
		amount INTEGER NOT NULL,
		status TEXT NOT NULL
	);`
)
