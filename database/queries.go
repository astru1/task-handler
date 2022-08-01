package database

const (
	createTableQ = `CREATE TABLE IF NOT EXISTS tasks (
		id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL,
		price VARCHAR ( 50 ) NOT NULL,
		priority VARCHAR ( 255 ) NOT NULL
	);`
	insertTaskQ    = `INSERT INTO tasks (name, price, priority) VALUES ($1, $2, $3)`
	selectAllTasks = `SELECT * FROM tasks`
)
