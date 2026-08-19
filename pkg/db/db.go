package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const (
	defaultDBFile  = "scheduler.db"
	createTableSQL = `
CREATE TABLE scheduler(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT ""
);
CREATE INDEX idx_scheduler_date ON scheduler (date);
`
)

var DB *sql.DB

func Init(dbFile string) error {
	var fileName = dbFile
	if fileName == "" {
		fileName = defaultDBFile
	}
	var install bool
	var err error

	_, err = os.Stat(fileName)
	install = os.IsNotExist(err)

	DB, err = sql.Open("sqlite", fileName)

	if err != nil {
		return err
	}

	if install {
		_, err = DB.Exec(createTableSQL)
		if err != nil {
			DB.Close()
			return err
		}
	}

	return nil
}
