package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	user "github.com/themanciraptor/Backend-photagea/internal/user/model"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

func main() {

	db, err := sql.Open("mysql", "ezdev:ForkmeMuthafukka@/photagea?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := userrepo.Initialize(db)

	ctx := context.Background()

	s, err := userRepository.Get(ctx, "13")
	if err != nil {
		log.Fatalf("Repository error: %s", err)
	}

	u := user.Model{
		UserID:    17,
		FirstName: "carll",
		LastName:  "Bennet",
		Alias:     "suckmaprick",
		Email:     "cocobean@cookie.net",
	}
	err = userRepository.Create(ctx, &u)
	if err != nil {
		log.Fatalf("Repository Error: %s", err)
	}

	fmt.Println(s)

}
