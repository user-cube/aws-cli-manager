package profile

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/user-cube/aws-cli-manager/v2/pkg/models"
	"github.com/user-cube/aws-cli-manager/v2/pkg/settings"
	"github.com/user-cube/aws-cli-manager/v2/pkg/sharedModules"
	"gopkg.in/yaml.v2"
)

func SelectProfile() {

	profileNames, awsProfiles := GetProfiles()

	// Prompt user to select a profile
	var selectedProfile string
	prompt := &survey.Select{
		Message: "Select an AWS profile:",
		Options: profileNames,
	}

	err := survey.AskOne(prompt, &selectedProfile)
	if err != nil {
		log.Fatalf("No profiles found, please execute 'aws-cli-manager profile add' to add a profile")
	}

	setProfile(selectedProfile, awsProfiles)

}

func GetProfiles() (profileNames []string, awsProfiles models.AwsProfile) {
	filePath := fmt.Sprintf("%s/%s/%s",
		sharedModules.GetHomeDirectory(),
		settings.AwsDefaultConfigurationFolder,
		settings.ConfigurationFilename,
	)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the configuration file if it doesn't exist
		awsProfiles = models.AwsProfile{
			Profiles: make(map[string]models.ProfileDetails),
		}
		saveProfiles(awsProfiles)
	}

	// Read the YAML file into yamlData
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Parse the YAML file into the config structure
	err = yaml.Unmarshal(yamlData, &awsProfiles)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	for profile := range awsProfiles.Profiles {
		profileNames = append(profileNames, profile)
	}

	return profileNames, awsProfiles
}

func setProfile(selectedProfile string, awsProfiles models.AwsProfile) {
	// Display the selected profile details
	selectedProfileDetails := awsProfiles.Profiles[selectedProfile]

	// We need to save credentials into ~/.aws/credentials
	// We need to save config into ~/.aws/config
	config := selectedProfileDetails.Config
	credentials := selectedProfileDetails.Credentials
	ssoEnabled := selectedProfileDetails.SSOEnabled

	header := "[default]"

	if selectedProfile == "global" {
		header = ""
	}

	// Define file paths
	homeDir := sharedModules.GetHomeDirectory()
	credentialsFilePath := fmt.Sprintf(
		"%s/%s/%s",
		homeDir,
		settings.AwsDefaultConfigurationFolder,
		settings.AwsCredentialsFilename,
	)
	configFilePath := fmt.Sprintf(
		"%s/%s/%s",
		homeDir,
		settings.AwsDefaultConfigurationFolder,
		settings.AwsConfigFilename,
	)

	// Write credentials to ~/.aws/credentials
	credentialsContent := fmt.Sprintf("%s\n%s", header, credentials)
	err := os.WriteFile(credentialsFilePath, []byte(credentialsContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write credentials file: %v", err)
	}

	// Write config to ~/.aws/config
	configContent := fmt.Sprintf("%s\n%s", header, config)
	err = os.WriteFile(configFilePath, []byte(configContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write config file: %v", err)
	}

	fmt.Println("Profile set successfully")

	if ssoEnabled {
		checkSSOSession()
	}
}

func checkSSOSession() {
	// Check if the SSO session is valid
	err := exec.Command("aws", "sts", "get-caller-identity").Run()
	if err != nil {
		fmt.Println("SSO session is invalid, starting a new session...")
		err = exec.Command("aws", "sso", "login").Run()
		if err != nil {
			log.Fatalf("Failed to start a new SSO session: %v", err)
		}
	}

}

func PromptProfileName() string {
	var profileName string
	for {
		prompt := &survey.Input{
			Message: "Enter the profile name:",
		}

		err := survey.AskOne(prompt, &profileName)
		if err != nil {
			log.Fatalf("Failed to get profile name: %v", err)
		}

		if profileName == "global" {
			fmt.Println("The profile name 'global' is reserved. Please enter a different name.")
		} else {
			break
		}
	}

	return profileName
}

func PromptProfileDetails() models.ProfileDetails {
	var profileDetails models.ProfileDetails

	// Prompt the user for profile details
	prompt := &survey.Input{
		Message: "Enter the region:",
	}
	err := survey.AskOne(prompt, &profileDetails.Region)
	if err != nil {
		log.Fatalf("Failed to get region: %v", err)
	}

	boolPrompt := &survey.Confirm{
		Message: "Is SSO enabled?",
	}
	err = survey.AskOne(boolPrompt, &profileDetails.SSOEnabled)
	if err != nil {
		log.Fatalf("Failed to get SSO information: %v", err)
	}

	multiLinePrompt := &survey.Multiline{
		Message: "Enter the config file information (e.g., sso_start_url, sso_region, etc.):",
	}
	err = survey.AskOne(multiLinePrompt, &profileDetails.Config)
	if err != nil {
		log.Fatalf("Failed to get config file information: %v", err)
	}

	// if region is not inside the config file, add it by using the region from the prompt
	if !strings.Contains(profileDetails.Config, "region") {
		profileDetails.Config = fmt.Sprintf("%s\nregion = %s", profileDetails.Config, profileDetails.Region)
	}

	multiLinePrompt = &survey.Multiline{
		Message: "Enter the credentials file information " +
			"(e.g., aws_access_key_id, aws_secret_access_key, aws_session_token):",
	}
	err = survey.AskOne(multiLinePrompt, &profileDetails.Credentials)
	if err != nil {
		log.Fatalf("Failed to get credentials file information: %v", err)
	}

	// Add the profile to the configuration file
	return profileDetails
}

func AddProfile(profileName string, profileDetails models.ProfileDetails) {
	// Get the existing profiles
	_, awsProfiles := GetProfiles()

	// Check if the global profile exists, if not create it
	if _, ok := awsProfiles.Profiles["global"]; !ok {
		awsProfiles.Profiles["global"] = models.ProfileDetails{
			Config:      "",
			Credentials: "",
			SSOEnabled:  false,
		}
	}

	// Append the new profile details to the global profile
	globalProfile := awsProfiles.Profiles["global"]
	globalProfile.Config += fmt.Sprintf("\n[profile %s]\n%s", profileName, profileDetails.Config)
	globalProfile.Credentials += fmt.Sprintf("\n[%s]\n%s", profileName, profileDetails.Credentials)
	globalProfile.SSOEnabled = false

	// Update the global profile in the map
	awsProfiles.Profiles["global"] = globalProfile

	// Save the updated profiles to the configuration file
	saveProfiles(awsProfiles)
}

func saveProfiles(awsProfiles models.AwsProfile) {
	// Define file path
	filePath := fmt.Sprintf("%s/%s/%s",
		sharedModules.GetHomeDirectory(),
		settings.AwsDefaultConfigurationFolder,
		settings.ConfigurationFilename,
	)

	// Marshal the profiles into YAML
	yamlData, err := yaml.Marshal(awsProfiles)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}

	// Write the YAML data to the file
	err = os.WriteFile(filePath, yamlData, 0644)
	if err != nil {
		log.Fatalf("Failed to write the file: %v", err)
	}
}
