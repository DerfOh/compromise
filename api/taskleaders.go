package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// TaskLeader to build the properties of what youre working with
type TaskLeader struct {
	TaskLeaderId int
	EmailAddress string
	GroupId      int
}

// APIHandler Respond to URLs of the form /generic/...

// TaskLeaderAPIHandler responds to /taskleaders/
func TaskLeaderAPIHandler(response http.ResponseWriter, request *http.Request) {

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
	case GET:
		GroupId := strings.Replace(request.URL.Path, "/api/taskleaders/", "", -1)
		st, getErr := db.Prepare("select * from TaskLeaders where GroupId=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(GroupId)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var TaskLeaderId int
			var EmailAddress string
			var GroupId int

			getErr := rows.Scan(&TaskLeaderId, &EmailAddress, &GroupId)
			taskleader := &TaskLeader{TaskLeaderId: TaskLeaderId, EmailAddress: EmailAddress, GroupId: GroupId}
			b, getErr := json.Marshal(taskleader)
			if getErr != nil {
				fmt.Println(getErr)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	// case POST:
	// 	//TaskId := request.PostFormValue("TaskId")
	// 	GroupId := request.PostFormValue("GroupId")
	// 	TaskName := request.PostFormValue("TaskName")
	// 	TaskDescription := request.PostFormValue("TaskDescription")
	// 	CompletionStatus := request.PostFormValue("CompletionStatus")
	// 	//CompletedBy := request.PostFormValue("CompletedBy")
	// 	PointValue := request.PostFormValue("PointValue")
	// 	st, postErr := db.Prepare("INSERT INTO Tasks(`taskid`, `groupid`, `taskname`, `taskdescription`, `completionstatus`, `completedby`,`pointvalue`) VALUES(NULL,?,?,?,?,NULL,?)")
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}
	// 	res, postErr := st.Exec(GroupId, TaskName, TaskDescription, CompletionStatus, PointValue)
	// 	if postErr != nil {
	// 		fmt.Print(postErr)
	// 	}

	// 	if res != nil {
	// 		result[0] = "Task Added"
	// 	}
	// 	result = result[:1]

	// case PUT:
	// 	GroupId := request.PostFormValue("GroupId")
	// 	TaskName := request.PostFormValue("TaskName")
	// 	TaskDescription := request.PostFormValue("TaskDescription")
	// 	CompletionStatus := request.PostFormValue("CompletionStatus")
	// 	CompletedBy := request.PostFormValue("CompletedBy")
	// 	PointValue := request.PostFormValue("PointValue")
	// 	TaskId := request.PostFormValue("TaskId")

	// 	st, putErr := db.Prepare("UPDATE Tasks SET GroupId=?, TaskName=?, TaskDescription=?, CompletionStatus=?, CompletedBy=?, PointValue=? WHERE TaskId=?")
	// 	if err != nil {
	// 		fmt.Print(putErr)
	// 	}
	// 	res, putErr := st.Exec(GroupId, TaskName, TaskDescription, CompletionStatus, CompletedBy, PointValue, TaskId)
	// 	if putErr != nil {
	// 		fmt.Print(putErr)
	// 	}

	// 	if res != nil {
	// 		result[0] = "Task Modified"
	// 	}
	// 	result = result[:1]
	// case DELETE:
	// 	TaskId := strings.Replace(request.URL.Path, "/api/tasks/", "", -1)
	// 	st, deleteErr := db.Prepare("DELETE FROM Tasks WHERE TaskId=?")
	// 	if deleteErr != nil {
	// 		fmt.Print(deleteErr)
	// 	}
	// 	res, deleteErr := st.Exec(TaskId)
	// 	if deleteErr != nil {
	// 		fmt.Print(deleteErr)
	// 	}

	// 	if res != nil {
	// 		result[0] = "Task Deleted"
	// 	}
	// 	result = result[:1]

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
