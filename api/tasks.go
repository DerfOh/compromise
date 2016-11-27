package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Task to build the properties of what youre working with
type Task struct {
	TaskId           int
	GroupId          string
	TaskName         string
	TaskDescription  string
	DateDue          string
	ApprovalStatus   string
	CompletionStatus string
	PointValue       int
}

// APIHandler Respond to URLs of the form /generic/...

// TaskAPIHandler responds to /tasks/
func TaskAPIHandler(response http.ResponseWriter, request *http.Request) {

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
		GroupId := strings.Replace(request.URL.Path, "/api/tasks/", "", -1)
		st, getErr := db.Prepare("select * from Tasks where GroupId=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(GroupId)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var TaskId int
			var GroupId string
			var TaskDescription string
			var DateDue string
			var TaskName string
			var ApprovalStatus string
			var CompletionStatus string
			var PointValue int

			getErr := rows.Scan(&TaskId, &GroupId, &TaskName, &TaskDescription, &DateDue, &ApprovalStatus, &CompletionStatus, &PointValue)
			task := &Task{TaskId: TaskId, GroupId: GroupId, TaskName: TaskName, TaskDescription: TaskDescription, DateDue: DateDue, ApprovalStatus: ApprovalStatus, CompletionStatus: CompletionStatus, PointValue: PointValue}
			b, getErr := json.Marshal(task)
			if getErr != nil {
				fmt.Println(getErr)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case POST:
		//TaskId :=
		GroupId := request.PostFormValue("GroupId")
		TaskName := request.PostFormValue("TaskName")
		TaskDescription := request.PostFormValue("TaskDescription")
		DateDue := request.PostFormValue("DateDue")
		ApprovalStatus := request.PostFormValue("ApprovalStatus")
		CompletionStatus := request.PostFormValue("CompletionStatus")
		PointValue := request.PostFormValue("PointValue")
		st, postErr := db.Prepare("INSERT INTO Tasks(`groupid`, `taskname`, `taskdescription`, `datedue`, `approvalstatus`, `completionstatus`, `pointvalue`) VALUES(?,?,?,?,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(GroupId, TaskName, TaskDescription, DateDue, ApprovalStatus, CompletionStatus, PointValue)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "Task Added"
		}
		result = result[:1]

	case PUT:
		TaskId := request.PostFormValue("TaskId")
		GroupId := request.PostFormValue("GroupId")
		TaskName := request.PostFormValue("TaskName")
		TaskDescription := request.PostFormValue("TaskDescription")
		DateDue := request.PostFormValue("DateDue")
		ApprovalStatus := request.PostFormValue("ApprovalStatus")
		CompletionStatus := request.PostFormValue("CompletionStatus")
		PointValue := request.PostFormValue("PointValue")

		st, putErr := db.Prepare("UPDATE Tasks SET GroupId=?, TaskName=?, TaskDescription=?, DateDue=?, ApprovalStatus=?, CompletionStatus=?, PointValue=? WHERE TaskId=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(GroupId, TaskName, TaskDescription, DateDue, ApprovalStatus, CompletionStatus, PointValue, TaskId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Task Modified"
		}
		result = result[:1]
	case DELETE:
		TaskId := strings.Replace(request.URL.Path, "/api/tasks/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM Tasks WHERE TaskId=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(TaskId)
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}

		if res != nil {
			result[0] = "Task Deleted"
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
