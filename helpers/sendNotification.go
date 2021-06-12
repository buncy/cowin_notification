package helpers

import (
	"cowin/centerTypes"
	"log"
	"os/exec"
)

var (
	filePath string
	notify   = false
)

func SendNotification(centers centerTypes.CowinCenters) error {

	for _, center := range centers.Centers {

		for _, session := range center.Sessions {

			if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 && session.Vaccine == "COVAXIN" {

				filePath = "./slots/COVAXIN.txt"

				createOutput(filePath, center)

				notify = true
			}
			if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 && session.Vaccine == "COVISHIELD" {

				filePath = "./slots/COVISHIELD.txt"

				createOutput(filePath, center)

				notify = true
			} else if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 {

				filePath = "./slots/other.txt"

				createOutput(filePath, center)

				notify = true
			}

		}
	}

	//send a desktop notification
	if notify {
		cmd := exec.Command("notify-send", "Congratulations! we have hope", `'Vaccine slots are available!'`)

		err := cmd.Run()

		if err != nil {
			log.Fatal("notify==>error : ", err)
		}
		WSLNnotify()
	}

	return nil
}
