package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Repeat  string `json:"repeat"`
	Comment string `json:"comment"`
}

func GetTasks(limit int, search string) ([]*Task, error) {
	tasks := []*Task{}
	var rows *sql.Rows
	var err error

	switch {
	case search == "":
		rows, err = DB.Query(
			`SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit`,
			sql.Named("limit", limit),
		)
	case isDate(search):
		t, _ := time.Parse("02.01.2006", search)
		rows, err = DB.Query(
			`SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date LIMIT :limit`,
			sql.Named("date", t.Format(DateFormat)),
			sql.Named("limit", limit),
		)
	default:
		like := "%" + search + "%"
		rows, err = DB.Query(
			`SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit`,
			sql.Named("search", like),
			sql.Named("limit", limit),
		)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func AddTask(task *Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`

	res, err := DB.Exec(
		query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetTask(id string) (*Task, error) {
	var task Task
	query := `SELECT id, date, title, comment, repeat FROM scheduler where id = :id`

	err := DB.QueryRow(
		query, sql.Named("id", id),
	).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat  WHERE id = :id`

	res, err := DB.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`задача не найдена`)
	}
	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler where id = :id`

	res, err := DB.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Задача не найдена")
	}

	return nil
}

func UpdateTaskDate(id string, date string) error {
	_, err := DB.Exec(
		`UPDATE scheduler SET date = :date WHERE id = :id`,
		sql.Named("date", date),
		sql.Named("id", id),
	)
	return err
}
