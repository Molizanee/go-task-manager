package main

import (
	"TaskManager/db"
	"TaskManager/routes"
	"context"

	"log"
	"net/http"
)

func main() {

	conn, err := db.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer conn.Close(context.Background())

	r := routes.Routes(conn)

	err = http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}
