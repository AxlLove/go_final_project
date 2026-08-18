package api

import (
	"net/http"
	"time"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func completeTaskHandle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeError(w, "id is required")
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, err.Error())
			return
		}
		writeJson(w, map[string]any{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeError(w, err.Error())
		return
	}
	err = db.UpdateTaskDate(id, next)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, map[string]any{})
}
