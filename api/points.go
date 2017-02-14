package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Points to build the properties of what youre working with
type Points struct {
	PointId      int
	TotalPoints  int
	EmailAddress string
	GroupId      int
}

// APIHandler Respond to URLs of the form /generic/...

// PointsAPIHandler responds to /points/
func PointsAPIHandler(response http.ResponseWriter, request *http.Request) {

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
		EmailAddress := strings.Replace(request.URL.Path, "/api/points/", "", -1)
		st, getErr := db.Prepare("select * from Points where EmailAddress=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(EmailAddress)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var PointId int
			var TotalPoints int
			var EmailAddress string
			var GroupId int

			getErr := rows.Scan(&PointId, &TotalPoints, &EmailAddress, &GroupId)
			points := &Points{PointId: PointId, TotalPoints: TotalPoints, EmailAddress: EmailAddress, GroupId: GroupId}
			b, getErr := json.Marshal(points)
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
		TotalPoints := request.PostFormValue("TotalPoints")
		GroupId := request.PostFormValue("GroupId")
		st, postErr := db.Prepare("INSERT INTO Points(`pointid`,`totalpoints`,`emailaddress`,`groupid`) VALUES(NULL,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(TotalPoints, EmailAddress, GroupId)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "Points Initialized"
		}
		result = result[:1]

	case PUT:
		PointId := request.PostFormValue("PointId")
		EmailAddress := request.PostFormValue("EmailAddress")
		TotalPoints := request.PostFormValue("TotalPoints")
		GroupId := request.PostFormValue("GroupId")

		st, putErr := db.Prepare("UPDATE Points SET TotalPoints=?, GroupId=? WHERE PointId=? AND EmailAddress=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(TotalPoints, GroupId, PointId, EmailAddress)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Points Modified"
		}
		result = result[:1]
	case DELETE:
		PointId := strings.Replace(request.URL.Path, "/api/points/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM Points WHERE PointId=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(PointId)
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}

		if res != nil {
			result[0] = "Points Entry Deleted"
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
