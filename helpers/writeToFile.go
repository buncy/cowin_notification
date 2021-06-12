package helpers

import (
	"cowin/centerTypes"
	"encoding/json"
	"fmt"
)

func createOutput(filePath string, center centerTypes.Center) {

	jsonData, _ := json.MarshalIndent(center, "", " ")
	data := "//----------start center------------//\n" + string(jsonData) + "\n//----------End center---------//\n\n"
	err := writeToFlie(filePath, data)
	if err != nil {
		fmt.Println("this is the error for err :", err.Error())
	}
}

/*
**this is for future vaccines sputnic v etc
 */

// notify = true
// name := session.Vaccine
// filepath = "./slots/" + name + ".txt"
// removeError := removeFile(filepath)
// if removeError != nil {
// 	fmt.Println("this is the error for removeError :", removeError.Error())
// }
// jsonData, _ := json.MarshalIndent(session, "", " ")
// data := "//----------start center------------//\n" + string(jsonData) + "\n//----------End center---------//\n\n"
// err := writeToFlie(filepath, data)
// if err != nil {
// 	fmt.Println("this is the error for err :", err.Error())
// }
