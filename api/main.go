// main process
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run *.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "golang.org/x/crypto/bcrypt"
)

// dbAddress the api will connect to, is set by flag, defaults to localhost passed in via args
var dbAddress string

// dbUserName database username passed in via args
var dbUserName string

// dbPassword the password to be used to connect to the database passed in via args
var dbPassword string

// dbConnectionURL the url used to perform db connection passed in via args
var dbConnectionURL string

// connection message
var logMessage string

// GET string constant
var GET = "GET"

// POST string constant
var POST = "POST"

// PUT string constant
var PUT = "PUT"

// DELETE string constant
var DELETE = "DELETE"

// main function for app
func main() {
	port := flag.Int("port", 1234, "an int")
	dbAddress := flag.String("dbaddress", "localhost", "a string")
	dbUserName := flag.String("dbuser", "compromise", "a string")
	dbPassword := flag.String("dbpassword", "password", "a string")
	//dbConnectionURL = *dbUserName + ":" + *dbPassword + "@tcp(" + *dbAddress + ":3306)/compromise"

	// Execute the command-line parsing.
	flag.Parse()

	dbConnectionURL = *dbUserName + ":" + *dbPassword + "@tcp(" + *dbAddress + ":3306)/compromise"

	// Show flag trail in logs, do not show dbPassword intentionally, just check if the password is the default
	fmt.Println("api port:\t\t", *port)
	fmt.Println("database address:\t", *dbAddress)
	fmt.Println("database user:\t\t", *dbUserName)
	if *dbPassword == "password" {
		fmt.Println("default db password:\t true")
		fmt.Println("connection url:", dbConnectionURL)
	} else {
		fmt.Println("default db password:\t false")
	}
	fmt.Println("Comments, remarks, general thoughts:", flag.Args())

	var err string
	portstring := strconv.Itoa(*port)

	mux := http.NewServeMux()
	//mux.Handle("/api/", http.HandlerFunc(APIHandler))
	// Handler for User interactions
	mux.Handle("/api/users/", http.HandlerFunc(UserAPIHandler))
	// Handler for Task interactions
	mux.Handle("/api/tasks/", http.HandlerFunc(TaskAPIHandler))
	// Handler for TaskLeader interactions
	mux.Handle("/api/taskleaders/", http.HandlerFunc(TaskLeaderAPIHandler))
	// Handler for Points interactions
	mux.Handle("/api/points/", http.HandlerFunc(PointsAPIHandler))
	// Handler for Reward interactions
	mux.Handle("/api/rewards/", http.HandlerFunc(RewardAPIHandler))
	// Handler for Reward interactions
	mux.Handle("/api/purchasedrewards/", http.HandlerFunc(PurchasedRewardsAPIHandler))
	// Hanlder for Group interactions
	mux.Handle("/api/groups/", http.HandlerFunc(GroupAPIHandler))
	// Handler for Authentication of users
	mux.Handle("/api/auth/", http.HandlerFunc(AuthAPIHandler))
	// Handler for retrieve password
	mux.Handle("/api/retrievepassword/", http.HandlerFunc(RetrievePasswordAPIHandler))
	// Default path handler
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	logMessage = "Listening on port " + portstring + " ... "
	log.Print(logMessage)

	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}

// CleanJSON to remove extra slashes from returned json objects
func CleanJSON(s string) string {
	// fmt.Println(s)
	s = strings.Replace(s, "\\\"", "\"", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "}\"", "}", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "\"{", "{", -1)
	// fmt.Println(s)
	return s
}
