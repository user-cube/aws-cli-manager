package sharedModules

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

func GetHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		message := fmt.Errorf("error getting home directory")
		fmt.Println(message)
		os.Exit(1)
	}
	return currentUser.HomeDir
}

func ListFiles(directory string, prefix string) *list.List {
	dir, err := os.Open(directory)
	if err != nil {
		message := fmt.Errorf("error opening directory")
		fmt.Println(message)
		os.Exit(1)
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		message := fmt.Errorf("error reading files")
		fmt.Println(message)
		os.Exit(1)
	}

	matchingFiles := list.New()

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), prefix) {
			matchingFiles.PushFront(file.Name())
		}
	}

	return matchingFiles
}

func CheckIfAWSDirectoryExists(homeDirectory string) bool {
	// Check if the .aws directory exists
	if _, err := os.Stat(homeDirectory + "/.aws"); os.IsNotExist(err) {
		return false
	}

	return true
}
