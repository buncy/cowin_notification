package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell.exe", "New-BurntToastNotification -Text", `Vaccine slots are available!`)

	err := cmd.Run()

	if err != nil {
		log.Fatal("notify", err)
	}
}
