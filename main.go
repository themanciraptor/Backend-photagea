package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	userapi "github.com/themanciraptor/Backend-photagea/API/user"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
	userservice "github.com/themanciraptor/Backend-photagea/internal/user/service"

	accountapi "github.com/themanciraptor/Backend-photagea/API/account"
	accountrepo "github.com/themanciraptor/Backend-photagea/internal/account/repo"
	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"

	imageapi "github.com/themanciraptor/Backend-photagea/API/image"
	imagerepo "github.com/themanciraptor/Backend-photagea/internal/image/repo"
	imageservice "github.com/themanciraptor/Backend-photagea/internal/image/service"

	imagedataapi "github.com/themanciraptor/Backend-photagea/API/imagedata"
	// imagedatarepo "github.com/themanciraptor/Backend-photagea/internal/imagedata/repo"
	// imagedataservice "github.com/themanciraptor/Backend-photagea/internal/imagedata/service"
)

const (
	port = ":5577"
)

func main() {
	// Sign in to DB
	db, err := sql.Open("mysql", "ezdev:developmentpassword@/photagea?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	accountRepo := accountrepo.Initialize(db)
	accountService := accountservice.Initialize(accountRepo)
	accountAPI := accountapi.Initialize(accountService)

	userRepo := userrepo.Initialize(db)
	userService := userservice.Initialize(userRepo)
	userAPI := userapi.Initialize(userService, accountService)

	imageRepo := imagerepo.Initialize(db)
	imageService := imageservice.Initialize(imageRepo)
	imageAPI := imageapi.Initialize(imageService, accountService)

	// imageRepo := imagerepo.Initialize(db)
	// imageService := imageservice.Initialize(imageRepo)
	imageDataAPI := imagedataapi.Initialize(accountService)

	RegisterRoutes(r, userAPI, accountAPI, imageAPI, imageDataAPI)

	http.Handle("/", r)
	log.Fatal(srv.ListenAndServe())
}
