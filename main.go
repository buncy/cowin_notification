package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	//"io"

	"log"
	"net/http"
	"os"
	"time"
)

type CowinCenters struct {
	Centers []Center `json:"centers"`
}

type Center struct {
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	Block_name   string        `json:"block_name"`
	Pincode      string        `json:"pincode"`
	Lat          string        `json:"lat"`
	Long         string        `json:"long"`
	From         string        `json:"from"`
	To           string        `json:"to"`
	Fee_type     string        `json:"fee_type"`
	Sessions     []Session     `json:"sessions"`
	Vaccine_fees []Vaccine_fee `json:"vaccine_fees"`
}

type Session struct {
	Available_capacity       string `json:"available_capacity"`
	Slots                    string `json:"slots"`
	Min_age_limit            string `json:"min_age_limit"`
	Vaccine                  string `json:"vaccine"`
	Available_capacity_dose1 int    `json:"available_capacity_dose1"`
	Available_capacity_dose2 int    `json:"available_capacity_dose2"`
}

type Vaccine_fee struct {
	Vaccine string `json:"vaccine"`
	Fee     string `json:"fee"`
}

var (
	filepath string
	notify   = false
)

func cowin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cowin! %s", time.Now())
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
	//content, err := ioutil.ReadAll(res.Body)
	var curentDaySessions CowinCenters
	jsonErr := json.NewDecoder(res.Body).Decode(&curentDaySessions)
	// jsonErr := json.Unmarshal(content, &curentDaySessions)
	if jsonErr != nil {
		fmt.Println("this is the error for jsonErr :", err.Error())
	}
	// if err != nil {
	// 	fmt.Println("this is the error for content :", err.Error())
	// }

	// notificationErr := sendNotification(curentDaySessions)

	// if notificationErr != nil {
	// 	fmt.Println("this is the error for notificationErr :", notificationErr.Error())
	// }
	fmt.Println("this is the data: ", curentDaySessions.Centers[0])

	//fmt.Println(string(content))
}

func sendNotification(centers CowinCenters) error {

	for _, center := range centers.Centers {

		for _, session := range center.Sessions {

			if session.Available_capacity_dose1 > 0 && session.Vaccine == "COVAXIN" {
				notify = true
				filepath = "./slots/COVAXIN.txt"
				removeError := removeFile(filepath)
				if removeError != nil {
					fmt.Println("this is the error for removeError :", removeError.Error())
				}
				jsonData, _ := json.MarshalIndent(session, "", " ")
				data := "//----------start center------------//\n" + string(jsonData) + "\n//----------End center---------//\n\n"
				err := writeToFlie(filepath, data)
				if err != nil {
					fmt.Println("this is the error for err :", err.Error())
				}
			}
			if session.Available_capacity_dose1 > 0 && session.Vaccine == "COVISHIELD" {
				notify = true
				filepath = "./slots/COVISHIELD.txt"
				removeError := removeFile(filepath)
				if removeError != nil {
					fmt.Println("this is the error for removeError :", removeError.Error())
				}
				jsonData, _ := json.MarshalIndent(session, "", " ")
				data := "//----------start center------------//\n" + string(jsonData) + "\n//----------End center---------//\n\n"
				err := writeToFlie(filepath, data)
				if err != nil {
					fmt.Println("this is the error for err :", err.Error())
				}
			} else if session.Available_capacity_dose1 > 0 {
				notify = true
				name := session.Vaccine
				filepath = "./slots/" + name + ".txt"
				removeError := removeFile(filepath)
				if removeError != nil {
					fmt.Println("this is the error for removeError :", removeError.Error())
				}
				jsonData, _ := json.MarshalIndent(session, "", " ")
				data := "//----------start center------------//\n" + string(jsonData) + "\n//----------End center---------//\n\n"
				err := writeToFlie(filepath, data)
				if err != nil {
					fmt.Println("this is the error for err :", err.Error())
				}
			}

		}
	}

	//send a desktop notification
	if notify {
		cmd := exec.Command("notify-send", "Congratulations! we have hope", `'Vaccine slots are available!'`)

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println(" -- error creating " + dir)
			return err
		}
	}
	return nil
}

func removeFile(filepath string) error {
	// Removing file
	// Using Remove() function

	if fileExists(filepath) {
		e := os.Remove(filepath)
		if e != nil {
			return e
		}
	}
	return nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeToFlie(filepath string, data string) error {
	// // check if file exists
	// var _, statErr = os.Stat(filepath)

	// // create file if not exists
	// if os.IsNotExist(statErr) {
	// 	CreateDir(filepath)
	// }

	// os.O_WRONLY tells the computer you are only going to writo to the file, not read
	// os.O_CREATE tells the computer to create the file if it doesn't exist
	// os.O_APPEND tells the computer to append to the end of the file instead of overwritting or truncating it
	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("this is the error for openfile :", err.Error())
	}
	// Create the file
	// out, err := os.Create(filepath)
	// if err != nil {
	// 	return err
	// }
	defer out.Close()

	// Write the body to file
	num, err := out.WriteString(data)

	fmt.Printf("wrote %d bytes\n", num)
	//Issue a Sync to flush writes to stable storage.

	out.Sync()
	return err
}

func main() {
	http.HandleFunc("/", cowin)
	http.ListenAndServe(":8080", nil)
}
