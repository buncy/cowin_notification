package helpers

import (
	"fmt"
	"log"
	"os"
)

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
