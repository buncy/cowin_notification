package helpers

import (
	"log"
	"os/exec"
)

func WSLNnotify() {
	cmd := exec.Command("powershell.exe", "New-BurntToastNotification -Text 'Vaccine slots are available!'")

	err := cmd.Run()

	if err != nil {
		log.Fatal("wsl notify error : ====>", err)
	}
}
