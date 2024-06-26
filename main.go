// Package main is the entry point of the application.
package main

import (
	"aws-cli-manager/pkg/cmd"
)

// main is the entry point function of the application.
// It calls the Execute function from the cmd package to start the application.
func main() {
	cmd.Execute()
}
