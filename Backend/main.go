//go version
// go mod init domain/project-name go module initialise domainname/project

package main

import (
	"fmt"
	"net/http"

	"tie.com/project1/db"
)

const PORT = ":9090"

func main() {
	db := db.NewDB()

	fmt.Println("Starting Server on port: %v\n http://localhost%v\n", PORT, PORT)
	http.HandleFunc("/search", db.HttpHandler)

	http.ListenAndServe(PORT, nil)

}
