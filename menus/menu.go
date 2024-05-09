package menus

import (
	"aws-cli-manager/aws"
	"aws-cli-manager/configurator"
	"fmt"
	"os"
)

var version = "0.1"

func displayHelp() {
	fmt.Println("Usage: aws-cli-manager [command]")
	fmt.Println("Commands:")
	fmt.Println("  configure       Configure AWS CLI")
	fmt.Println("  help            Display help")
	fmt.Println("  version         Display version")
	fmt.Println("  profile         Select AWS CLI profile")
}

func Menu() {
	if len(os.Args) < 2 {
		displayHelp()
		return
	}

	switch os.Args[1] {
	case "configure":
		configurator.ConfigureAWSCLI()
	case "profile":
		aws.Profiles()
	case "help":
		displayHelp()
	case "version":
		fmt.Println("AWS CLI Manager version " + version)
	default:
		fmt.Println("Invalid command")
		displayHelp()
	}
}
