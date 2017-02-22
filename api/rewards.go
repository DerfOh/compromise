package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Reward to build the properties of what youre working with
type Reward struct {
	RewardId          int
	GroupId           int
	RewardName        string
	PointCost         int
	RewardDescription string
}

// APIHandler Respond to URLs of the form /generic/...

// RewardAPIHandler responds to /Rewards/
func RewardAPIHandler(response http.ResponseWriter, request *http.Request) {

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
		GroupId := strings.Replace(request.URL.Path, "/api/rewards/", "", -1)
		//fmt.Println(GroupId)
		st, getErr := db.Prepare("SELECT * FROM Rewards WHERE GroupId=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(GroupId)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var RewardId int
			var GroupId int
			var RewardName string
			var PointCost int
			var RewardDescription string

			getErr := rows.Scan(&RewardId, &GroupId, &RewardName, &PointCost, &RewardDescription)
			reward := &Reward{RewardId: RewardId, GroupId: GroupId, RewardName: RewardName, PointCost: PointCost, RewardDescription: RewardDescription}
			b, getErr := json.Marshal(reward)
			if getErr != nil {
				fmt.Println(getErr)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case POST:
		//RewardId := request.PostFormValue("RewardId")
		GroupId := request.PostFormValue("GroupId")
		RewardName := request.PostFormValue("RewardName")
		RewardDescription := request.PostFormValue("RewardDescription")
		PointCost := request.FormValue("PointCost")
		st, postErr := db.Prepare("INSERT INTO Rewards(`rewardid`, `groupid`, `rewardname`, `rewarddescription`, `pointcost`) VALUES(NULL,?,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(GroupId, RewardName, RewardDescription, PointCost)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "Reward Added"
		}
		result = result[:1]

	case PUT:
		GroupId := request.PostFormValue("GroupId")
		RewardName := request.PostFormValue("RewardName")
		RewardDescription := request.PostFormValue("RewardDescription")
		PointCost := request.PostFormValue("PointCost")
		RewardId := request.PostFormValue("RewardId")

		st, putErr := db.Prepare("UPDATE Rewards SET GroupId=?, RewardName=?, RewardDescription=?, PointCost=? WHERE RewardId=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(GroupId, RewardName, RewardDescription, PointCost, RewardId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Reward Modified"
		}
		result = result[:1]

	case DELETE:
		RewardId := strings.Replace(request.URL.Path, "/api/rewards/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM Rewards WHERE RewardId=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(RewardId)
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}

		if res != nil {
			result[0] = "Reward Deleted"
		}
		result = result[:1]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client. Clean backslashes from json
	fmt.Fprintf(response, "%v", CleanJSON(string(json)))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}
