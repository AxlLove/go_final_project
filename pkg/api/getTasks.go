package api

import (
	"net/http"

	"github.com/AxlLove/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandle(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := db.GetTasks(50, search)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
