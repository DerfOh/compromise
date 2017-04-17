package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Task to build the properties of what youre working with
type CompletedTask struct {
	TaskId           int
	GroupId          string
	TaskName         string
	TaskDescription  string
	CompletionStatus string
	CompletedBy      string
	PointValue       int
	TotalPoints      int
	EmailAddress     string
}

// APIHandler Respond to URLs of the form /generic/...

// CompleteTaskAPIHandler responds to /completetask/
func CompleteTaskAPIHandler(response http.ResponseWriter, request *http.Request) {
	t := time.Now()
	logRequest := t.Format("2006/01/02 15:04:05") + " | Request:" + request.Method + " | Endpoint: completetask | "
	fmt.Println(logRequest)
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
	var result = make([]string, 1000)

	switch request.Method {
	case POST:
		// Get information from db
		EmailAddress := request.PostFormValue("EmailAddress")
		TaskId := request.PostFormValue("TaskId")
		GroupId := request.PostFormValue("GroupId")

		// Retrieve point value from database
		var PointValue string
		pointValueQueryErr := db.QueryRow("SELECT PointValue from Tasks where TaskId=?", TaskId).Scan(&PointValue)
		switch {
		case pointValueQueryErr == sql.ErrNoRows:
			log.Printf(logRequest, "No Task with ID: \n", TaskId)
		case pointValueQueryErr != nil:
			log.Fatal(pointValueQueryErr)
		default:
		}

		// Retrieve User's TotalPoints from database
		var TotalPoints string
		pointTotalQueryErr := db.QueryRow("SELECT `TotalPoints` from `Points` where `EmailAddress`=? AND `GroupId`=?", EmailAddress, GroupId).Scan(&TotalPoints)
		switch {
		case pointTotalQueryErr == sql.ErrNoRows:
			log.Printf(logRequest, "No User with Email: \n", EmailAddress)
		case pointTotalQueryErr != nil:
			log.Fatal(pointTotalQueryErr)
		default:
			//fmt.Printf("Password is %s\n", Password)
		}

		// Add PointValue to PointTotal
		TotalPoints += PointValue

		// Perform update action on task table
		st, putErr := db.Prepare("UPDATE `Tasks` SET `CompletionStatus`=?, `CompletedBy`=? WHERE `TaskId`=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec("Complete", EmailAddress, TaskId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Task Modified"
		}
		result = result[:1]

		// Perform update action on points table
		st, putErr = db.Prepare("UPDATE Points SET TotalPoints=? WHERE EmailAddress=? AND GroupId=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr = st.Exec(TotalPoints, EmailAddress, GroupId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "User Modified"
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
	fmt.Fprintf(response, "%v", CleanJSON(string(json)))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}
