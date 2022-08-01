package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func CreateDBConnection(host, port, user, password, dbname string) (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB) error {
	query := createTableQ
	_, err := db.Exec(query)
	return err
}

func InsertTask(db *sql.DB, task Task) error {
	query := insertTaskQ
	_, err := db.Exec(query, task.Name, task.Price, task.Priority)
	return err
}

func ReturnTasks(db *sql.DB) ([]Task, error) {

	rows, err := db.Query(selectAllTasks)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.Id, &t.Name, &t.Price, &t.Priority)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, t)
	}

	return tasks, err
}
