package config

import (
	"fmt"
)

var (
	version = ""
)

func LogVersion() {
	fmt.Println("AWS CLI Manager version " + version) // Prints the version of the AWS CLI Manager
}
