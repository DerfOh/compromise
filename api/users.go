package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
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
	TotalPoints  int
	//Password     string
}

// APIHandler Respond to URLs of the form /generic/...

// UserAPIHandler responds to /user/
func UserAPIHandler(response http.ResponseWriter, request *http.Request) {
	t := time.Now()
	logRequest := t.Format("2006/01/02 15:04:05") + " | Request:" + request.Method + " | Endpoint: users | "
	fmt.Println(logRequest)
	//Connect to database
	db, e := sql.Open("mysql", dbConnectionURL)
	if e != nil {
		fmt.Print(e)
	}

	// set mime type to JSON
	response.Header().Set("Content-type", "application/json")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	// can't define dynamic slice in golang
	var result = make([]string, 1000)

	switch request.Method {
	case GET:
		GroupId := strings.Replace(request.URL.Path, "/api/users/", "", -1)

		st, getErr := db.Prepare("SELECT Users.EmailAddress, Users.FirstName, Users.LastName, Users.Nickname, Points.TotalPoints FROM Users JOIN Points ON Users.EmailAddress = Points.EmailAddress JOIN Groups ON Groups.GroupId = Points.GroupId WHERE Groups.GroupId=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(GroupId)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var EmailAddress string
			var FirstName string
			var LastName string
			var Nickname string
			var TotalPoints int
			getErr := rows.Scan(&EmailAddress, &FirstName, &LastName, &Nickname, &TotalPoints)
			user := &User{EmailAddress: EmailAddress, FirstName: FirstName, LastName: LastName, NickName: Nickname, TotalPoints: TotalPoints}
			b, getErr := json.Marshal(user)
			if getErr != nil {
				fmt.Println(getErr)
				return
			}

			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case POST:
		EmailAddress := request.PostFormValue("EmailAddress")
		FirstName := request.PostFormValue("FirstName")
		LastName := request.PostFormValue("LastName")
		Nickname := request.PostFormValue("Nickname")
		Password, _ := HashPassword(request.PostFormValue("Password"))
		//GroupId := request.PostFormValue("GroupId")
		st, postErr := db.Prepare("INSERT INTO Users VALUES(?,?,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(EmailAddress, FirstName, LastName, Nickname, Password)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "User Added"
		}
		result = result[:1]

	case PUT:
		FirstName := request.PostFormValue("FirstName")
		LastName := request.PostFormValue("LastName")
		Nickname := request.PostFormValue("Nickname")
		Password, _ := HashPassword(request.PostFormValue("Password"))
		EmailAddress := request.PostFormValue("EmailAddress")

		st, putErr := db.Prepare("UPDATE Users SET FirstName=?, LastName=?, Nickname=?, Password=? WHERE EmailAddress=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(FirstName, LastName, Nickname, Password, EmailAddress)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "User Modified"
		}
		result = result[:1]
	case DELETE:
		EmailAddress := strings.Replace(request.URL.Path, "/api/users/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM Users WHERE EmailAddress=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(EmailAddress)
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}

		if res != nil {
			result[0] = EmailAddress + " removed"
		}
		result = result[:1]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Clean up JSON before returning
	// Send the text diagnostics to the client.
	//fmt.Fprintf(response, "%v", CleanJSON(string(json)))
	fmt.Fprintf(response, "%v", CleanJSON(string(json)))

	db.Close()
}
