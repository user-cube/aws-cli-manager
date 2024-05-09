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

func CheckIfProfileExists(profile string) bool {
	// Check if the profile exists

	homeDirectory := GetHomeDirectory()

	dirExists := CheckIfAWSDirectoryExists(homeDirectory)

	awsDir := homeDirectory + "/.aws"

	if !dirExists {
		// Create the .aws directory
		err := os.Mkdir(awsDir, 0700)

		if err != nil {
			message := fmt.Errorf("error creating .aws directory: %v", err)
			fmt.Println(message)
			os.Exit(1)
		}

	}

	// Check if profile is in the credentials file
	credentialsFile := homeDirectory + "/.aws/credentials-" + profile
	_, err := os.Stat(credentialsFile)

	return !os.IsNotExist(err)
}

// CopyFile copies a file from source to destination.
func CopyFile(source string, destination string) error {
	// Open the source file
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Flushes any buffered data to the file
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
