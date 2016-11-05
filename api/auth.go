// Simple auth endpoint. This probably needs to be swapped out for something more sophisticated at some point

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Auth to build the properties of what youre working with
type Auth struct {
	Nickname string
	Password string
}

// APIHandler Respond to URLs of the form /generic/...

// AuthAPIHandler responds to /auth/
func AuthAPIHandler(response http.ResponseWriter, request *http.Request) {

	//Connect to database
	db, e := sql.Open("mysql", "compromise:password@tcp(localhost:3306)/compromise")
	if e != nil {
		fmt.Print(e)
	}

	//set mime type to JSON
	response.Header().Set("Content-type", "application/json")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	//can't define dynamic slice in golang
	var result = make([]string, 1000)

	switch request.Method {
	case "POST":
		Nickname := request.PostFormValue("Nickname")
		fmt.Printf("Nickname is %s\n", Nickname)
		//Nickname := "Coolguy1234"
		//Password := request.PostFormValue("Password")

		// SELECT Password FROM Users WHERE Nickname=?
		var Password string
		queryErr := db.QueryRow("SELECT Password FROM Users WHERE Nickname=?", Nickname).Scan(&Password)
		switch {
		case queryErr == sql.ErrNoRows:
			log.Printf("No user with Nickname: %s\n", Nickname)
		case queryErr != nil:
			log.Fatal(queryErr)
		default:
			fmt.Printf("Password is %s\n", Password)
		}

		// Compare variable returned from db query to provided Password

		//return true if true

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}
