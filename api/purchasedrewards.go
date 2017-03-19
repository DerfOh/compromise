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

// Reward to build the properties of what youre working with
type PurchasedReward struct {
	RequestId         int
	GroupId           int
	RewardName        string
	PointCost         int
	RewardDescription string
	RewardedUser      string
}

// APIHandler Respond to URLs of the form /generic/...

// PurchasedRewardsAPIHandler responds to /purchasedrewards/
func PurchasedRewardsAPIHandler(response http.ResponseWriter, request *http.Request) {
	t := time.Now()
	logRequest := t.Format("2006/01/02 15:04:05") + " | Request:" + request.Method + " | Endpoint: purchasedrewards | "	//Connect to database
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
		GroupId := strings.Replace(request.URL.Path, "/api/purchasedrewards/", "", -1)

		//fmt.Println(GroupId)
		st, getErr := db.Prepare("select * from PurchasedRewards where GroupId=?")
		if err != nil {
			fmt.Print(getErr)
		}
		rows, getErr := st.Query(GroupId)
		if getErr != nil {
			fmt.Print(getErr)
		}
		i := 0
		for rows.Next() {
			var RequestId int
			var GroupId int
			var RewardName string
			var PointCost int
			var RewardDescription string
			var RewardedUser string

			getErr := rows.Scan(&RequestId, &GroupId, &RewardName, &PointCost, &RewardDescription, &RewardedUser)
			reward := &PurchasedReward{RequestId: RequestId, GroupId: GroupId, RewardName: RewardName, PointCost: PointCost, RewardDescription: RewardDescription, RewardedUser: RewardedUser}
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
		//RequestId := request.PostFormValue("RequestId")
		GroupId := request.PostFormValue("GroupId")
		RewardName := request.PostFormValue("RewardName")
		PointCost := request.PostFormValue("PointCost")
		RewardDescription := request.PostFormValue("RewardDescription")
		RewardedUser := request.PostFormValue("RewardedUser")
		st, postErr := db.Prepare("INSERT INTO PurchasedRewards(`requestid`, `groupid`, `rewardname`, `pointcost`, `rewarddescription`, `rewardeduser`) VALUES(NULL,?,?,?,?,?)")
		if err != nil {
			fmt.Print(err)
		}
		res, postErr := st.Exec(GroupId, RewardName, PointCost, RewardDescription, RewardedUser)
		if postErr != nil {
			fmt.Print(postErr)
		}

		if res != nil {
			result[0] = "Purchase Added"
		}
		result = result[:1]

	case PUT:
		RequestId := request.PostFormValue("RequestId")
		GroupId := request.PostFormValue("GroupId")
		RewardName := request.PostFormValue("RewardName")
		PointCost := request.PostFormValue("PointCost")
		RewardDescription := request.PostFormValue("RewardDescription")
		RewardedUser := request.PostFormValue("RewardedUser")

		st, putErr := db.Prepare("UPDATE PurchasedRewards SET GroupId=?, RewardName=?, PointCost=?, RewardDescription=?, RewardedUser=? WHERE RequestId=?")
		if err != nil {
			fmt.Print(putErr)
		}
		res, putErr := st.Exec(GroupId, RewardName, PointCost, RewardDescription, RewardedUser, RequestId)
		if putErr != nil {
			fmt.Print(putErr)
		}

		if res != nil {
			result[0] = "Reward Modified"
		}
		result = result[:1]

	case DELETE:
		RequestId := strings.Replace(request.URL.Path, "/api/purchasedrewards/", "", -1)
		st, deleteErr := db.Prepare("DELETE FROM PurchasedRewards where RequestId=?")
		if deleteErr != nil {
			fmt.Print(deleteErr)
		}
		res, deleteErr := st.Exec(RequestId)
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
