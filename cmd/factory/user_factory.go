package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/banisys/user-service/pkg/database"
	"github.com/bxcodec/faker/v3"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func main() {
	count := flag.Int("n", 5, "Number of fake users to create")
	flag.Parse()

	DB := database.DB()
	defer DB.Close()

	for i := 0; i < *count; i++ {
		name := faker.Name()
		email := faker.Email()
		password := faker.Password()

		_, err := DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
		if err != nil {
			log.Println("Insert error:", err)
		} else {
			fmt.Printf("Inserted user %d: %s\n", i+1, email)
		}
	}
}
