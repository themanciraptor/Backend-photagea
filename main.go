package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

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
	db, err := sql.Open("mysql", "ezdev:ForkmeMuthafukka@/photagea?parseTime=true") //TODO: this should be a compile time secret
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	RegisterRoutes(userAPI, accountAPI, imageAPI, imageDataAPI)

	http.ListenAndServe(":8001", nil)
}
