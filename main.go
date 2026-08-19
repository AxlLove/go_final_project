package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AxlLove/go_final_project/pkg/db"
	"github.com/AxlLove/go_final_project/pkg/server"
)

const (
	BasePort  = 7540
	StaticDir = "./web"
)

func main() {
	port := os.Getenv("TODO_PORT")
	dbFile := os.Getenv("TODO_DBFILE")
	err := db.Init(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	if port == "" {
		port = strconv.Itoa(BasePort)
	}

	s := server.CreateServer(":"+port, StaticDir)

	log.Printf("Starting server on :%s port", port)
	if err = s.HTTP.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
}
