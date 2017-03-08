// Simple auth endpoint. Needs to be swapped out for something more sophisticated eventually.

/*
 This really needs to be handled differently. It gets the job done for now
 though, the password is sent as a response then sending the password is
 handled within the application in future itorations it would be a better
 idea to handle this through the use of a one-time login token that is put in
 place of the password in the Users table then once the user logs in they
 are prompted to reset their password to a new one.
*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Password to build the properties of what youre working with
type Password struct {
	Password string
}

var result string

// APIHandler Respond to URLs of the form /generic/...

// RetrievePasswordAPIHandler responds to /retrievepassword/
func RetrievePasswordAPIHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint request: /retrievepassword/ ")
	//Connect to database
	db, e := sql.Open("mysql", dbConnectionURL)
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
	// var result = make([]string, 1)

	switch request.Method {

	case POST:
		EmailAddress := request.PostFormValue("EmailAddress")
		// fmt.Printf("EmailAddress is %s\n", EmailAddress)
		var Password string
		queryErr := db.QueryRow("SELECT Password FROM Users WHERE EmailAddress=?", EmailAddress).Scan(&Password)
		switch {
		case queryErr == sql.ErrNoRows:
			log.Printf("No user with EmailAddress: %s\n", EmailAddress)
		case queryErr != nil:
			log.Fatal(queryErr)
		default:
			//fmt.Printf("Password is %s\n", Password)
		}
		result = Password

	default:
		result = "No response..."
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
