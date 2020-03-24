package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	userapi "github.com/themanciraptor/Backend-photagea/API/user"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
	userservice "github.com/themanciraptor/Backend-photagea/internal/user/service"
)

const (
	port = ":5577"
)

func main() {
	// Sign in to DB
	db, err := sql.Open("mysql", "ezdev:ForkmeMuthafukka@/photagea?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := userrepo.Initialize(db)
	userService := userservice.Initialize(userRepo)
	userAPI := userapi.Initialize(userService)

	RegisterRoutes(userAPI)

	http.ListenAndServe(":5577", nil)
}
