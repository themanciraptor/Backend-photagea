package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	userrepo "github.com/themanciraptor/Backend-photagea/internal/user/repo"
)

func main() {

	db, err := sql.Open("mysql", "ezdev:ForkmeMuthaFukka@/photagea")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := userrepo.Initialize(db)

	s, _ := userRepository.Get("me")

	fmt.Printf("%s", s)

}
