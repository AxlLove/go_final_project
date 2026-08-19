package api

import (
	"net/http"

	"github.com/AxlLove/go_final_project/pkg/db"
)

const tasksLimit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandle(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := db.GetTasks(tasksLimit, search)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
