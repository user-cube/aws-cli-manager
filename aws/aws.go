package aws

import (
	"aws-cli-manager/sharedModules"
	"bufio"
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
	fmt.Println("  select          Select a profile")
	fmt.Println("  credentials     Export credentials to environment variables, you should execute this command with eval, e.g. eval $(aws-cli-manager profile credentials)")
	fmt.Println("  help            Display help")
}

func Profiles() {
	if len(os.Args) < 3 {
		displayHelp()
		return
	}

	switch os.Args[2] {
	case "list":
		ListProfiles()
	case "select":
		selectProfile()
	case "credentials":
		exportCredentialsToEnvironmentVariables()
	default:
		fmt.Println("Invalid command")
		displayHelp()

	}
}

func selectProfile() {
	homeDirectory := sharedModules.GetHomeDirectory()
	userInput := ""

	// Check if arguments are provided
	if len(os.Args) < 4 {
		ListProfiles()

		// Create a new scanner to read user input
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt the user to enter a string
		fmt.Print("Please select a profile: ")

		// Read the user input
		scanner.Scan()
		userInput = scanner.Text()
	} else {
		userInput = os.Args[3]
	}

	// Check if the profile exists
	profileExists := sharedModules.CheckIfProfileExists(userInput)

	if !profileExists {
		fmt.Println("Profile does not exist")
		return
	} else {

		// Copy file .aws/credentials-profile to .aws/credentials
		err := sharedModules.CopyFile(homeDirectory+"/.aws/credentials-"+userInput, homeDirectory+"/.aws/credentials")
		if err != nil {
			fmt.Println("Error copying credentials file")
			return
		}

		// Copy file .aws/config-profile to .aws/config
		err = sharedModules.CopyFile(homeDirectory+"/.aws/config-"+userInput, homeDirectory+"/.aws/config")
		if err != nil {
			fmt.Println("Error copying config file")
			return
		}

		fmt.Println("Profile selected successfully, using profile: " + userInput)
	}
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

func exportCredentialsToEnvironmentVariables() {

	// We need to get variables from the credentials file
	// and export them to the environment variables

	// Get the home directory of the user
	homeDirectory := sharedModules.GetHomeDirectory()

	// Get the path to the credentials file
	credentialsFile := homeDirectory + "/.aws/credentials"

	// Open the credentials file

	file, err := os.Open(credentialsFile)

	if err != nil {
		fmt.Println("Error opening credentials file")
		return
	}

	// Close the file after the function ends
	defer file.Close()

	// Create a new scanner to read the file

	scanner := bufio.NewScanner(file)

	// Create a map to store the credentials

	credentials := make(map[string]string)

	// Read the file line by line, we need to ignore the first line [default]
	// and split the line by "=" to get the key and value

	// The credentials file has the following format:
	// [default]
	// aws_access_key_id = YOUR_ACCESS
	// aws_secret_access_key = YOUR_SECRET

	// We need to ignore the first line and split the line by "=" to get the key and valu

	for scanner.Scan() {

		line := scanner.Text()
		if line != "[default]" {
			parts := strings.Split(line, "=")
			credentials[strings.TrimSpace(parts[0])] = parts[1]
		}

	}

	fmt.Println("AWS_ACCESS_KEY_ID=\"" + credentials["aws_access_key_id"] + "\"")
	fmt.Println("AWS_SECRET_ACCESS_KEY=\"" + credentials["aws_secret_access_key"] + "\"")

}
