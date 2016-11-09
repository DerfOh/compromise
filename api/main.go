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
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	port := flag.Int("port", 1234, "an int")

	// Execute the command-line parsing.
	flag.Parse()

	// Show port and trail in logs
	fmt.Println("port:", *port)
	fmt.Println("tail:", flag.Args())

	var err string
	portstring := strconv.Itoa(*port)

	mux := http.NewServeMux()
	//mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/api/users/", http.HandlerFunc(UserAPIHandler))     // Handler for User interactions
	mux.Handle("/api/tasks/", http.HandlerFunc(TaskAPIHandler))     // Handler for Task interactions
	mux.Handle("/api/rewards/", http.HandlerFunc(RewardAPIHandler)) // Handler for Reward interactions
	mux.Handle("/api/groups/", http.HandlerFunc(GroupAPIHandler))   // Hanlder for Group interactions
	mux.Handle("/api/auth/", http.HandlerFunc(AuthAPIHandler))      // Handler for Authentication of users
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func cleanJSON(s string) string {
	// fmt.Println(s)
	s = strings.Replace(s, "\\\"", "\"", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "}\"", "}", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "\"{", "{", -1)
	// fmt.Println(s)
	return s
}
