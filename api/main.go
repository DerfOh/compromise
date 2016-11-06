// server.go
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run server.go

package main

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := 80
	var err string
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	//mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/api/users", http.HandlerFunc(UserAPIHandler))     // Handler for User interactions
	mux.Handle("/api/tasks", http.HandlerFunc(TaskAPIHandler))     // Handler for Task interactions
	mux.Handle("/api/rewards", http.HandlerFunc(RewardAPIHandler)) // Handler for Reward interactions
	mux.Handle("/api/groups", http.HandlerFunc(GroupAPIHandler))   // Hanlder for Group interactions
	mux.Handle("/api/auth", http.HandlerFunc(AuthAPIHandler))      // Handler for Authentication of users
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
