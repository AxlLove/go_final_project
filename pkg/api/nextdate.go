package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is required")
	}
	d, err := time.Parse(db.DateFormat, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")

	switch parts[0] {
	case "y":
		for {
			d = d.AddDate(1, 0, 0)
			if d.After(now) {
				break
			}
		}
		return d.Format(db.DateFormat), nil
	case "d":
		if len(parts) < 2 {
			return "", errors.New("need at least two parts")
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		if n < 1 || n > 400 {
			return "", errors.New("interval must be between 1 and 400")
		}

		for {
			d = d.AddDate(0, 0, n)
			if d.After(now) {
				break
			}
		}

		return d.Format(db.DateFormat), nil
	case "w":
		if len(parts) < 2 {
			return "", errors.New("w: не указаны дни недели")
		}
		var weekdays [8]bool
		for _, s := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil || n < 1 || n > 7 {
				return "", errors.New("w: недопустимое значение дня недели")
			}
			weekdays[n] = true
		}
		for {
			d = d.AddDate(0, 0, 1)
			wd := int(d.Weekday())
			if wd == 0 {
				wd = 7
			}
			if weekdays[wd] && d.After(now) {
				break
			}
		}
		return d.Format(db.DateFormat), nil
	default:
		return "", errors.New("invalid repeat format")
	}
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(db.DateFormat, nowStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	result, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(result))
}
