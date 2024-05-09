package aws

import (
	"aws-cli-manager/sharedModules"
	"container/list"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

func displayHelp() {
	// This function will display the help menu
	fmt.Println("Usage: aws-cli-manager profile [command]")
	fmt.Println("Commands:")
	fmt.Println("  list            List all available profiles")
}

func Profiles() {
	if len(os.Args) < 3 {
		displayHelp()
		return
	}

	switch os.Args[2] {
	case "list":
		ListProfiles()
	default:
		fmt.Println("Invalid command")
		displayHelp()

	}
}

func SelectProfile() {
	// This function will allow the user to select an AWS CLI profile
}

func ListProfiles() *list.List {

	// Get the home directory of the user
	homeDirectory := sharedModules.GetHomeDirectory()

	// Get the path to the .aws directory
	awsDirectory := homeDirectory + "/.aws"

	// Check if the .aws directory exists
	dirExists := sharedModules.CheckIfAWSDirectoryExists(homeDirectory)

	if !dirExists {
		fmt.Println("No profiles found")
		return nil
	}

	// List all files that start with "credentials" in the .aws directory
	files := sharedModules.ListFiles(awsDirectory, "credentials-")

	// Create a new table
	t := table.NewWriter()
	t.SetOutputMirror(nil)
	t.SetStyle(table.StyleLight)

	// Define table headers
	t.AppendHeader(table.Row{"Name", "File Name", "File Location"})

	// Populate the table
	for e := files.Front(); e != nil; e = e.Next() {
		profile := strings.Split(e.Value.(string), "credentials-")[1]
		fileName := e.Value.(string)
		fileLocation := awsDirectory + "/" + fileName
		t.AppendRow([]interface{}{profile, fileName, fileLocation})
	}

	// Render the table
	renderedTable := t.Render()

	// Print the rendered table
	fmt.Println(renderedTable)

	return files
}
