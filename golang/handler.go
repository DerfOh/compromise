package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User to build the properties of what youre working with
type User struct {
	EmailAddress string
	FirstName    string
	LastName     string
	NickName     string
	Password     string
}

// Handler for all requests
func Handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

// APIHandler Respond to URLs of the form /generic/...
func APIHandler(response http.ResponseWriter, request *http.Request) {

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
	case "GET":
		st, err := db.Prepare("select * from Users limit 10")
		if err != nil {
			fmt.Print(err)
		}
		rows, err := st.Query()
		if err != nil {
			fmt.Print(err)
		}
		i := 0
		for rows.Next() {
			//var Id int
			var EmailAddress string
			var FirstName string
			var LastName string
			var Nickname string
			var Password string
			err := rows.Scan(&EmailAddress, &FirstName, &LastName, &Nickname, &Password)
			user := &User{EmailAddress: EmailAddress, FirstName: FirstName, LastName: LastName, NickName: Nickname, Password: Password}
			b, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case "POST":
		EmailAddress := request.PostFormValue("EmailAddress")
		FirstName := request.PostFormValue("FirstName")
		LastName := request.PostFormValue("LastName")
		Nickname := request.PostFormValue("Nickname")
		Password := request.PostFormValue("Password")
		st, err := db.Prepare("INSERT INTO Users VALUES(?,?,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(EmailAddress, FirstName, LastName, Nickname, Password)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]

	case "PUT":
		FirstName := request.PostFormValue("FirstName")
		LastName := request.PostFormValue("LastName")
		Nickname := request.PostFormValue("Nickname")
		Password := request.PostFormValue("Password")
		EmailAddress := request.PostFormValue("EmailAddress")

		st, err := db.Prepare("UPDATE Users SET FirstName=?, LastName=?, Nickname=?, Password=? WHERE EmailAddress=?")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(FirstName, LastName, Nickname, Password, EmailAddress)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]
	case "DELETE":
		id := strings.Replace(request.URL.Path, "/api/", "", -1)
		st, err := db.Prepare("DELETE FROM Users WHERE EmailAddress=?")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(id)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]

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
