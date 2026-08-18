package db

import "time"

func isDate(s string) bool {
	_, err := time.Parse("02.01.2006", s)
	return err == nil
}
