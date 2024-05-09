package sharedModules

import (
	"container/list"
	"os"
	"os/user"
	"strings"
)

func GetHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return currentUser.HomeDir
}

func ListFiles(directory string, prefix string) *list.List {
	dir, err := os.Open(directory)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	matchingFiles := list.New()

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), prefix) {
			matchingFiles.PushFront(file.Name())
		}
	}

	return matchingFiles
}
