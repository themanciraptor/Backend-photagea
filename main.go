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

	s, err := userRepository.Get(ctx, 12)
	if err != nil {
		log.Fatalf("Repository Get error: %s", err)
	}

	u := user.Model{
		UserID:    12,
		FirstName: "carll3",
		LastName:  "Ben",
		Alias:     "suckick",
		AccountID: 12,
	}

	err = userRepository.Update(ctx, &u)
	if err != nil {
		log.Fatalf("Repository Update Error: %s", err)
	}

	u = user.Model{
		UserID:    17,
		FirstName: "carll",
		LastName:  "Bennet",
		Alias:     "suckmaprick",
		AccountID: 14,
	}
	err = userRepository.Create(ctx, &u)
	if err != nil {
		log.Fatalf("Repository Create Error: %s", err)
	}

	fmt.Println(s)
}
