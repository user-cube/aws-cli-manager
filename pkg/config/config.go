package config

import "fmt"

var (
	Version string
	Date    string
)

func LogVersion() {
	fmt.Println("AWS CLI Manager version " + Version + ", release date: " + Date) // Prints the version of the AWS CLI Manager
}
