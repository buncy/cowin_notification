package helpers

import (
	"cowin/centerTypes"
	"log"
	"os/exec"
)

var (
	filepath string
	notify   = false
)

func SendNotification(centers centerTypes.CowinCenters) error {

	for _, center := range centers.Centers {

		for _, session := range center.Sessions {

			if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 && session.Vaccine == "COVAXIN" {

			}
			if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 && session.Vaccine == "COVISHIELD" {

			} else if session.Available_capacity_dose1 > 0 && session.Min_age_limit == 18 {

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
