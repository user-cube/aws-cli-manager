package configurator

import (
	"aws-cli-manager/sharedModules"
	"bufio"
	"fmt"
	"os"
)

func ConfigureAWSCLI() {
	configurationFile := setUpConfigurationFile()
	awsKeyID, awsKeySecret := askCredentials()
	storeCredentials(awsKeyID, awsKeySecret, configurationFile)
}

func setUpConfigurationFile() string {

	homeDirectory := sharedModules.GetHomeDirectory()
	dirExists := checkIfAWSDirectoryExists(homeDirectory)

	if !dirExists {
		// Create the .aws directory
		err := os.Mkdir(homeDirectory+"/.aws", 0700)

		if err != nil {
			panic(err)
		}

	}

	// Create a new scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter a string
	fmt.Print("Please select a name for your new profile: ")

	// Read the user input
	scanner.Scan()
	userInput := scanner.Text()

	configurationFile := homeDirectory + "/.aws/credentials-" + userInput

	// Create a new file in the home directory
	file, err := os.Create(configurationFile)

	if err != nil {
		panic(err)
	}

	// Close the file after the function ends
	defer file.Close()

	return configurationFile

}

func checkIfAWSDirectoryExists(homeDirectory string) bool {
	// Check if the .aws directory exists
	if _, err := os.Stat(homeDirectory + "/.aws"); os.IsNotExist(err) {
		return false
	}

	return true
}

func askCredentials() (awsKeyID string, awsKeySecret string) {
	// Create a new scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter a string
	fmt.Print("Please enter your AWS Access Key ID: ")

	// Read the user input
	scanner.Scan()
	awsKeyID = scanner.Text()

	// Prompt the user to enter a string
	fmt.Print("Please enter your AWS Secret Access Key: ")

	// Read the user input
	scanner.Scan()
	awsKeySecret = scanner.Text()

	return awsKeyID, awsKeySecret

}

func storeCredentials(awsKeyID string, awsKeySecret string, configurationFile string) {
	// Open the file in append mode
	file, err := os.OpenFile(configurationFile, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// Write the AWS credentials to the file
	_, err = file.WriteString("[default]\n")
	_, err = file.WriteString("aws_access_key_id = " + awsKeyID + "\n")
	_, err = file.WriteString("aws_secret_access_key = " + awsKeySecret + "\n")

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// Close the file after the function ends
	defer file.Close()

	fmt.Println("AWS CLI has been configured successfully")
}
