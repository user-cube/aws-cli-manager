package configurator

import (
	"aws-cli-manager/pkg/sharedModules"
	"bufio"
	"fmt"
	"os"
)

// ConfigureAWSCLI sets up the credentials file, asks for credentials, and stores them.
func ConfigureAWSCLI() {
	credentialsFile, configurationFile := setUpCredentialsFile()
	awsKeyID, awsKeySecret, region := askCredentials()
	storeCredentials(awsKeyID, awsKeySecret, region, credentialsFile, configurationFile)
}

// setUpCredentialsFile sets up the credentials file and returns its path.
// It creates the .aws directory if it doesn't exist.
func setUpCredentialsFile() (credentialsFile string, configurationFile string) {

	homeDirectory := sharedModules.GetHomeDirectory()
	awsDir := homeDirectory + "/.aws"
	dirExists := sharedModules.CheckIfAWSDirectoryExists(homeDirectory)

	if !dirExists {
		// Create the .aws directory
		err := os.Mkdir(awsDir, 0700)

		if err != nil {
			message := fmt.Errorf("error creating .aws directory: %v", err)
			fmt.Println(message)
			os.Exit(1)
		}

	}

	// Create a new scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter a string
	fmt.Print("Please select a name for your new profile: ")

	// Read the user input
	scanner.Scan()
	userInput := scanner.Text()

	credentialsFile = homeDirectory + "/.aws/credentials-" + userInput
	configurationFile = homeDirectory + "/.aws/config-" + userInput

	// Create a new file in the home directory
	file, err := os.Create(credentialsFile)
	if err != nil {
		message := fmt.Errorf("error creating credentials file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	defer file.Close()

	file, err = os.Create(configurationFile)
	if err != nil {
		message := fmt.Errorf("error creating configuration file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Close the file after the function ends
	defer file.Close()

	return credentialsFile, configurationFile

}

// askCredentials prompts the user to enter their AWS credentials and returns them.
func askCredentials() (awsKeyID string, awsKeySecret string, region string) {
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

	// Prompt the user to enter a string
	fmt.Print("Please enter your AWS Region: ")

	// Read the user input
	scanner.Scan()
	region = scanner.Text()

	return awsKeyID, awsKeySecret, region

}

// storeCredentials stores the provided AWS credentials in the specified credentials file.
func storeCredentials(awsKeyID string, awsKeySecret string, region string, credentialsFile string, configurationFile string) {
	// Open the file in append mode
	file, err := os.OpenFile(credentialsFile, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		message := fmt.Errorf("error opening credentials file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Write the AWS credentials to the file
	_, err = file.WriteString("[default]\n")
	if err != nil {
		message := fmt.Errorf("error writing to credentials file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	_, err = file.WriteString("aws_access_key_id = " + awsKeyID + "\n")
	if err != nil {
		message := fmt.Errorf("error writing to credentials file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	_, err = file.WriteString("aws_secret_access_key = " + awsKeySecret + "\n")
	if err != nil {
		message := fmt.Errorf("error writing to credentials file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	defer file.Close()

	file, err = os.OpenFile(configurationFile, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		message := fmt.Errorf("error opening configuration file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Write the AWS configuration to the file
	_, err = file.WriteString("[default]\n")
	if err != nil {
		message := fmt.Errorf("error writing to configuration file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	_, err = file.WriteString("region = " + region + "\n")
	if err != nil {
		message := fmt.Errorf("error writing to configuration file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Close the file after the function ends
	defer file.Close()

	fmt.Println("AWS CLI has been configured successfully")
}
