// Simple auth endpoint. Needs to be swapped out for something more sophisticated eventually.

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Auth to build the properties of what youre working with
type Auth struct {
	EmailAddress string
	Password     string
}

// APIHandler Respond to URLs of the form /generic/...

// AuthAPIHandler responds to /auth/
func AuthAPIHandler(response http.ResponseWriter, request *http.Request) {

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
	var result string

	switch request.Method {

	/*
	 This really needs to be handled differently. It gets the job done for now
	 though, the password is sent as a response then sending the password is
	 handled within the application in future itorations it would be a better
	 idea to handle this through the use of a one-time login token that is put in
	 place of the password in the Users table then once the user logs in they
	 are prompted to reset their password to a new one.
	*/

	case GET:
		EmailAddress := strings.Replace(request.URL.Path, "/api/auth/", "", -1)
		// fmt.Printf("EmailAddress is %s\n", EmailAddress)
		var Password string
		queryErr := db.QueryRow("SELECT Password FROM Users WHERE EmailAddress=?", EmailAddress).Scan(&Password)
		switch {
		case queryErr == sql.ErrNoRows:
			log.Printf("No user with EmailAddress: %s\n", EmailAddress)
		case queryErr != nil:
			log.Fatal(queryErr)
		default:
			result = "No user with EmailAddress"
		}
		result = Password

	case POST:
		EmailAddress := request.PostFormValue("EmailAddress")
		// fmt.Printf("EmailAddress is %s\n", EmailAddress)
		ProvidedPassword := request.PostFormValue("Password")
		// fmt.Printf("Password provided is: %s\n", ProvidedPassword)
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

		// Compare variable returned from db query to provided Password
		if Password == ProvidedPassword {
			//return true if true
			//fmt.Println("Password Match")
			//result[0] = "Match"
			result = "true"
		} else {
			//fmt.Println("Password Mismatch")
			//result[0] = "Invalid"
			result = "false"
		}

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
