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
		merchant_name TEXT NOT NULL UNIQUE,
		date DATETIME NOT NULL,
		amount INTEGER NOT NULL,
		status TEXT NOT NULL
	);`
)
