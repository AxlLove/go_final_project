package api

import (
	"encoding/json"
	"net/http"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func updateTaskHandle(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.UpdateTask(&task)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJson(w, map[string]any{})
}
