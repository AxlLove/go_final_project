package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	var next string
	if task.Date == "" {
		task.Date = now.Format("20060102")
		return nil
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}

	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if today.After(t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}

	return nil
}

func addTaskHandle(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, err.Error())
		return
	}

	if task.Title == "" {
		writeError(w, "title is required")
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, err.Error())
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, map[string]any{"id": strconv.FormatInt(id, 10)})
}
