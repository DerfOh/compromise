package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Group properties
type Group struct {
	GroupId     int
	GroupName   string
	TotalPoints int
}

// APIHandler Respond to URLs of the form /generic/...

// GroupAPIHandler responds to /groups/
func GroupAPIHandler(response http.ResponseWriter, request *http.Request) {
	t := time.Now()
	logRequest := t.Format("2006/01/02 15:04:05") + " | Request:" + request.Method + " | Endpoint: groups | " //Connect to database
	fmt.Println(logRequest)
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
	case GET:
		EmailAddress := strings.Replace(request.URL.Path, "/api/groups/", "", -1)
		//fmt.Println(EmailAddress)

		st, getErr := db.Prepare("SELECT Points.GroupId, Groups.GroupName, Points.TotalPoints FROM Points JOIN Groups ON Points.GroupId = Groups.GroupId WHERE EmailAddress=?")
		if err != nil {
			fmt.Print(getErr)
		}

		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(EmailAddress)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var GroupId int
			var GroupName string
			var TotalPoints int
			getErr := rows.Scan(&GroupId, &GroupName, &TotalPoints)
			group := &Group{GroupId: GroupId, GroupName: GroupName, TotalPoints: TotalPoints}
			b, getErr := json.Marshal(group)
			if getErr != nil {
				fmt.Println(getErr)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case POST:
		GroupName := request.PostFormValue("GroupName")
		st, postErr := db.Prepare("INSERT INTO Groups(`groupid`, `groupname`) VALUES(NULL,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(GroupName)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "Group Added"
		}
		result = result[:1]

	case PUT:
		GroupId := request.PostFormValue("GroupId")
		GroupName := request.PostFormValue("GroupName")

		st, putErr := db.Prepare("UPDATE Groups SET GroupName=? WHERE GroupId=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(GroupName, GroupId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Group Modified"
		}
		result = result[:1]

	case DELETE:
		GroupId := strings.Replace(request.URL.Path, "/api/groups/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM Groups WHERE GroupId=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(GroupId)
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}

		if res != nil {
			result[0] = "Group Deleted"
		}
		result = result[:1]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client and remove backslashes
	fmt.Fprintf(response, "%v", CleanJSON(string(json)))
	db.Close()
}
