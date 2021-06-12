package helpers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func writeToFlie(filePath string, data string) error {

	absPath, err := filepath.Abs(filePath)
	fmt.Println(absPath)
	if err != nil {
		fmt.Println("error while absPath", err.Error())
	}
	DirErr := EnsureBaseDir(absPath)
	if err != nil {
		fmt.Println("dir error", DirErr.Error())
	}

	out, err := os.OpenFile(absPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("err creating file", err.Error())
	}

	defer out.Close()

	// Write the body to file
	num, err := out.WriteString(data)

	fmt.Printf("wrote %d bytes\n", num)
	//Issue a Sync to flush writes to stable storage.

	out.Sync()
	return err
}
func EnsureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
