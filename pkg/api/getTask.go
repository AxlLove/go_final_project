package api

import (
	"net/http"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func getTaskHandle(w http.ResponseWriter, r *http.Request) {
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

	writeJson(w, task)
}
