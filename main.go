package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

func main() {

	db, err := sql.Open("mysql", "ezdev:ForkmeMuthafukka@/photagea")
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

	fmt.Println(s)

}
