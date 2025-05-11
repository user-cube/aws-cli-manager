package profile

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/user-cube/aws-cli-manager/v2/pkg/models"
	"github.com/user-cube/aws-cli-manager/v2/pkg/settings"
	"github.com/user-cube/aws-cli-manager/v2/pkg/sharedModules"
	"github.com/user-cube/aws-cli-manager/v2/pkg/ui"
	"gopkg.in/yaml.v2"
)

func SelectProfile() bool {
	// Get current profile first
	currentProfile := GetCurrentProfile()
	profileNames, awsProfiles := GetProfiles()

	if len(profileNames) == 0 {
		ui.PrintError("No profiles found, please execute 'aws-cli-manager profile add' to add a profile")
		os.Exit(1)
	}

	// Sort profile names alphabetically and create display names with current marker
	var displayNames []string
	var sortedNames []string

	// First add current profile if it exists
	if currentProfile != "" {
		// Make both the profile name and (current) marker green
		displayNames = append(displayNames, ui.SuccessColor(fmt.Sprintf("%s (current)", currentProfile)))
		sortedNames = append(sortedNames, currentProfile)
	}

	// Then add remaining profiles in alphabetical order
	for _, name := range profileNames {
		if name != currentProfile {
			displayNames = append(displayNames, name)
			sortedNames = append(sortedNames, name)
		}
	}

	// Create and run the profile selection prompt
	prompt := ui.CreateSelectPrompt(ui.SelectOptions{
		Label: "Select an AWS profile",
		Items: displayNames,
	})

	index, _, err := prompt.Run()
	if err != nil {
		// Check if it's an interrupt error (Ctrl+C)
		if err == promptui.ErrInterrupt {
			fmt.Println("\nProfile selection cancelled")
			return false
		}
		// Handle other errors
		ui.PrintError("Error selecting profile: %v", err)
		os.Exit(1)
	}

	// Get the selected profile name
	selectedProfile := sortedNames[index]
	setProfile(selectedProfile, awsProfiles)
	return true
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

	// Get profile names and sort them alphabetically
	for profile := range awsProfiles.Profiles {
		profileNames = append(profileNames, profile)
	}

	// Sort the profile names alphabetically
	sort.Strings(profileNames)

	return profileNames, awsProfiles
}

// GetCurrentProfile returns the currently selected profile
func GetCurrentProfile() string {
	_, awsProfiles := GetProfiles()
	return awsProfiles.CurrentProfile
}

func ShowCurrentProfile() {
	currentProfile := GetCurrentProfile()
	if currentProfile == "" {
		ui.PrintInfo("No profile is currently selected. Run 'aws-cli-manager' to select a profile.")
		return
	}

	_, awsProfiles := GetProfiles()
	profileDetails := awsProfiles.Profiles[currentProfile]

	ui.PrintCurrentProfile(currentProfile, profileDetails.Region, profileDetails.SSOEnabled)
}

func setProfile(selectedProfile string, awsProfiles models.AwsProfile) {
	// Display the selected profile details
	selectedProfileDetails := awsProfiles.Profiles[selectedProfile]

	// Save the current profile
	awsProfiles.CurrentProfile = selectedProfile
	saveProfiles(awsProfiles)

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
		ui.PrintError("Failed to write config file: %v", err)
		os.Exit(1)
	}

	// Format the success message with the profile name in green
	fmt.Printf("%s %s %s\n",
		ui.NormalColor("Profile"),
		ui.SuccessColor(selectedProfile),
		ui.NormalColor("set successfully"))

	if ssoEnabled {
		checkSSOSession()
	}
}

func checkSSOSession() {
	// Check if the SSO session is valid
	err := exec.Command("aws", "sts", "get-caller-identity").Run()
	if err != nil {
		ui.PrintInfo("SSO session is invalid, starting a new session...")
		err = exec.Command("aws", "sso", "login").Run()
		if err != nil {
			ui.PrintError("Failed to start a new SSO session: %v", err)
			os.Exit(1)
		}
		ui.PrintSuccess("SSO session started successfully")
	}
}

func PromptProfileName() string {
	var profileName string
	for {
		prompt := ui.CreateTextPrompt(ui.TextPromptOptions{
			Label:    "Enter the profile name",
			Validate: nil,
		})

		result, err := prompt.Run()
		if err != nil {
			// Check if it's an interrupt error (Ctrl+C)
			if err == promptui.ErrInterrupt {
				fmt.Println("\nProfile creation cancelled")
				os.Exit(0)
			}
			// Handle other errors
			ui.PrintError("Failed to get profile name: %v", err)
			os.Exit(1)
		}

		profileName = result

		if profileName == "global" {
			ui.PrintError("The profile name 'global' is reserved. Please enter a different name.")
		} else {
			break
		}
	}

	return profileName
}

func PromptProfileDetails() models.ProfileDetails {
	var profileDetails models.ProfileDetails

	// Prompt for region
	regionPrompt := ui.CreateTextPrompt(ui.TextPromptOptions{
		Label:    "Enter the region",
		Default:  "us-east-1",
		Validate: nil,
	})

	region, err := regionPrompt.Run()
	if err != nil {
		// Check if it's an interrupt error (Ctrl+C)
		if err == promptui.ErrInterrupt {
			fmt.Println("\nProfile creation cancelled")
			os.Exit(0)
		}
		ui.PrintError("Failed to get region: %v", err)
		os.Exit(1)
	}
	profileDetails.Region = region

	// Prompt for SSO
	ssoPrompt := ui.CreateSelectPrompt(ui.SelectOptions{
		Label: "Is SSO enabled?",
		Items: []string{"Yes", "No"},
	})

	ssoIndex, _, err := ssoPrompt.Run()
	if err != nil {
		// Check if it's an interrupt error (Ctrl+C)
		if err == promptui.ErrInterrupt {
			fmt.Println("\nProfile creation cancelled")
			os.Exit(0)
		}
		ui.PrintError("Failed to get SSO information: %v", err)
		os.Exit(1)
	}
	profileDetails.SSOEnabled = ssoIndex == 0 // Index 0 is "Yes"

	// Prompt for config details
	configPrompt := ui.CreateTextPrompt(ui.TextPromptOptions{
		Label:    "Enter the config file information (e.g., sso_start_url, sso_region, etc.)",
		Validate: nil,
	})

	config, err := configPrompt.Run()
	if err != nil {
		// Check if it's an interrupt error (Ctrl+C)
		if err == promptui.ErrInterrupt {
			fmt.Println("\nProfile creation cancelled")
			os.Exit(0)
		}
		ui.PrintError("Failed to get config file information: %v", err)
		os.Exit(1)
	}
	profileDetails.Config = config

	// if region is not inside the config file, add it by using the region from the prompt
	if !strings.Contains(profileDetails.Config, "region") {
		profileDetails.Config = fmt.Sprintf("%s\nregion = %s", profileDetails.Config, profileDetails.Region)
	}

	// Prompt for credentials
	credentialsPrompt := ui.CreateTextPrompt(ui.TextPromptOptions{
		Label:    "Enter the credentials file information (e.g., aws_access_key_id, aws_secret_access_key, aws_session_token)",
		Validate: nil,
	})

	credentials, err := credentialsPrompt.Run()
	if err != nil {
		// Check if it's an interrupt error (Ctrl+C)
		if err == promptui.ErrInterrupt {
			fmt.Println("\nProfile creation cancelled")
			os.Exit(0)
		}
		ui.PrintError("Failed to get credentials file information: %v", err)
		os.Exit(1)
	}
	profileDetails.Credentials = credentials

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

// SelectProfileByName selects a profile by name
func SelectProfileByName(profileName string) error {
	_, awsProfiles := GetProfiles()

	// Check if the profile exists
	if _, ok := awsProfiles.Profiles[profileName]; !ok {
		return fmt.Errorf("profile %s does not exist", profileName)
	}

	// Set the profile
	setProfile(profileName, awsProfiles)
	return nil
}
