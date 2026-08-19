package api

import "github.com/go-chi/chi/v5"

func Init(r *chi.Mux) {
	r.Get("/api/nextdate", nextDateHandler)
	r.Post("/api/signin", signInHandle)

	r.With(authMiddleware).Get("/api/tasks", getTasksHandle)
	r.With(authMiddleware).Post("/api/task/done", completeTaskHandle)
	r.With(authMiddleware).Get("/api/task", getTaskHandle)
	r.With(authMiddleware).Post("/api/task", addTaskHandle)
	r.With(authMiddleware).Put("/api/task", updateTaskHandle)
	r.With(authMiddleware).Delete("/api/task", deleteTaskHandle)
}
