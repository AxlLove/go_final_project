package api

import (
	"net/http"

	"github.com/AxlLove/go_final_project/pkg/db"
)

func deleteTaskHandle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		writeError(w, "id is required")
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, map[string]any{})
}
