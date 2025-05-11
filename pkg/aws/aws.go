// Package aws provides functions to manage AWS profiles.
package aws

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/user-cube/aws-cli-manager/v2/pkg/profile"
	"github.com/user-cube/aws-cli-manager/v2/pkg/ui"
)

// ListProfiles lists all available AWS profiles and returns a list of them.
func ListProfiles() {
	_, awsProfiles := profile.GetProfiles()

	// Prepare table data
	headers := []string{"Profile Name", "Region", "SSO Enabled"}
	rows := []ui.TableData{}

	// Fill the table with the profiles
	for awsProfile, details := range awsProfiles.Profiles {
		ssoStatus := "No"
		if details.SSOEnabled {
			ssoStatus = color.GreenString("Yes")
		} else {
			ssoStatus = color.RedString("No")
		}

		rows = append(rows, ui.TableData{
			Columns: []interface{}{
				color.HiWhiteString(awsProfile),
				color.YellowString(details.Region),
				ssoStatus,
			},
		})
	}

	// Render the table
	ui.RenderTable(ui.TableOptions{
		Headers: headers,
		Rows:    rows,
	})
}

// AddProfile adds a new AWS profile to the configuration file.
func AddProfile() {
	profileName := profile.PromptProfileName()

	// Check if the profile already exists
	_, awsProfiles := profile.GetProfiles()
	if _, ok := awsProfiles.Profiles[profileName]; ok {
		ui.PrintError("Profile %s already exists", profileName)
		os.Exit(1)
	}

	// Prompt the user for profile details
	profileDetails := profile.PromptProfileDetails()

	// Add the profile to the configuration file
	profile.AddProfile(profileName, profileDetails)

	ui.PrintSuccess("Profile %s added successfully", profileName)
}

// TestConnection tests the connection to AWS by executing the 'aws sts get-caller-identity' command.
// This command returns details about the IAM user or role whose credentials are used to call the operation.
func TestConnection() {
	// The command to be executed
	// 'aws sts get-caller-identity' returns details about the IAM user or role whose credentials are used to call the operation
	ui.PrintInfo("Testing connection to AWS...")
	out, err := exec.Command("aws", "sts", "get-caller-identity").Output()

	// Check if there was an error executing the command
	if err != nil {
		ui.PrintError("Error executing command")
		ui.PrintError("%v", err)
		os.Exit(1)
	}

	// Print the output of the command
	ui.PrintSuccess("Connection successful!")
	fmt.Println(string(out))
}
