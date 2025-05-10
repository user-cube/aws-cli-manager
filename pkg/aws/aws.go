// Package aws provides functions to manage AWS profiles.
package aws

import (
	"fmt" // Importing fmt for output formatting
	"os"  // Importing os for file and directory operations
	"os/exec"

	"github.com/jedib0t/go-pretty/v6/table" // Importing table for creating tables
	"github.com/user-cube/aws-cli-manager/pkg/profile"
)

// ListProfiles lists all available AWS profiles and returns a list of them.
func ListProfiles() {

	_, awsProfiles := profile.GetProfiles()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Profile Name", "Region", "SSO Enabled"})

	// Fill the table with the profiles
	for awsProfile, details := range awsProfiles.Profiles {
		t.AppendRow(table.Row{
			awsProfile,
			details.Region,
			fmt.Sprintf("%v", details.SSOEnabled),
		})
	}

	t.SetStyle(table.StyleLight)
	t.Render()

}

// AddProfile adds a new AWS profile to the configuration file.
func AddProfile() {

	profileName := profile.PromptProfileName()

	// Check if the profile already exists
	_, awsProfiles := profile.GetProfiles()
	if _, ok := awsProfiles.Profiles[profileName]; ok {
		fmt.Println("Profile already exists")
		os.Exit(1)
	}

	// Prompt the user for profile details
	profileDetails := profile.PromptProfileDetails()

	// Add the profile to the configuration file
	profile.AddProfile(profileName, profileDetails)

	fmt.Println("Profile added successfully")
}

// TestConnection tests the connection to AWS by executing the 'aws sts get-caller-identity' command.
// This command returns details about the IAM user or role whose credentials are used to call the operation.
func TestConnection() {

	// The command to be executed
	// 'aws sts get-caller-identity' returns details about the IAM user or role whose credentials are used to call the operation
	out, err := exec.Command("aws", "sts", "get-caller-identity").Output()

	// Check if there was an error executing the command
	if err != nil {
		fmt.Println("Error executing command")
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the output of the command
	fmt.Println(string(out))
}
