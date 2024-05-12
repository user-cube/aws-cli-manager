package config

import "fmt"

var (
	version = ""
	date    = ""
)

func LogVersion() {
	fmt.Println("AWS CLI Manager version " + version + ", release date: " + date) // Prints the version of the AWS CLI Manager
}
