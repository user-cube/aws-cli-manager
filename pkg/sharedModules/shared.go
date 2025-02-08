// Package sharedModules provides utility functions for the application.
package sharedModules

import (
	"fmt"
	"log"
	"os/user"
)

// GetHomeDirectory returns the home directory of the current user.
func GetHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		message := fmt.Errorf("error getting home directory")
		log.Fatalf("%v: %v", message, err)
	}
	return currentUser.HomeDir
}
