package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	userapi "github.com/themancirapter/Backend-photagea/API/user"
	userrepo "github.com/themancirapter/Backend-photagea/Internal/user/repo"
	userservice "github.com/themancirapter/Backend-photagea/Internal/user/service"
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
}
