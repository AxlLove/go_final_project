package db

import "time"

const DateFormat = "20060102"

func isDate(s string) bool {
	_, err := time.Parse("02.01.2006", s)
	return err == nil
}
