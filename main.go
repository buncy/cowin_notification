package main

import (
	"encoding/json"
	"fmt"
	"os"

	//"path/filepath"

	"net/http"
	"time"

	ctype "cowin/centerTypes"

	helpers "cowin/helpers"
)

func cowin() {
	fmt.Println("cowin started")
	date := time.Now().Format("02-01-2006")
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=363&date=" + date
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("this is the error for request :", err.Error())
	}
	req.Header.Add("Accept-Language", "en_US")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("this is the error for cowin :", err.Error())
	}
	defer res.Body.Close()

	var curentDaySessions ctype.CowinCenters
	jsonErr := json.NewDecoder(res.Body).Decode(&curentDaySessions)

	if jsonErr != nil {
		fmt.Println("this is the error for jsonErr :", jsonErr.Error())
	}
	//remove file if already present
	removeError := os.RemoveAll("slots")
	if removeError != nil {
		fmt.Println("this is the error for removeError :", removeError.Error())
	} else if removeError == nil {
		fmt.Println("folder removed==========================================>")
	}
	notificationErr := helpers.SendNotification(curentDaySessions)

	if notificationErr != nil {
		fmt.Println("this is the error for notificationErr :", notificationErr.Error())
	}
	// fmt.Println("this is the data: ", curentDaySessions.Centers[0])

	//spew.Dump("this is the data: ", curentDaySessions)

	defer fmt.Println("cowin exited ")
}

func main() {
	ticker := time.NewTicker(4 * time.Second)
	now := time.Now()
	for ; true; <-ticker.C {
		cowin()
		fmt.Println("time is =========================================>", time.Since(now))
	}
}
