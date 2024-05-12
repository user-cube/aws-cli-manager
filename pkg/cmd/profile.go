// Package cmd provides the command line interface for the AWS CLI Manager.
package cmd

import (
	"aws-cli-manager/pkg/aws"
	"aws-cli-manager/pkg/configurator"
	"github.com/spf13/cobra" // Importing the cobra package for creating CLI applications
)

// profileCmd represents the 'profile' command.
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "AWS Profile Manager is a tool to manage AWS CLI profiles",
}

// listCmd represents the 'list' command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available profiles",
	Run: func(cmd *cobra.Command, args []string) {
		aws.ListProfiles() // Calls the ListProfiles function from the aws package
	},
}

// selectCmd represents the 'select' command.
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a profile",
	Args:  cobra.ArbitraryArgs,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetProfileNames(), cobra.ShellCompDirectiveNoFileComp // Returns the profile names from the aws package
	},
	Run: func(cmd *cobra.Command, args []string) {
		aws.SelectProfile() // Calls the SelectProfile function from the aws package
	},
}

// exportCredentialsCmd represents the 'credentials' command.
var exportCredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Export credentials to environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		aws.ExportCredentialsToEnvironmentVariables() // Calls the ExportCredentialsToEnvironmentVariables function from the aws package
	},
}

// configureCmd represents the 'configure' command.
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure AWS CLI Profile",
	Run: func(cmd *cobra.Command, args []string) {
		configurator.ConfigureAWSCLI() // Calls the ConfigureAWSCLI function from the configurator package
	},
}

// init function to initialize the root command and add subcommands.
func init() {
	rootCmd.AddCommand(profileCmd) // Adds the 'profile' command as a subcommand to the root command

	profileCmd.Flags()                          // Initializes flags for the 'profile' command
	profileCmd.AddCommand(listCmd)              // Adds the 'list' command as a subcommand to the 'profile' command
	profileCmd.AddCommand(selectCmd)            // Adds the 'select' command as a subcommand to the 'profile' command
	profileCmd.AddCommand(exportCredentialsCmd) // Adds the 'credentials' command as a subcommand to the 'profile' command
	profileCmd.AddCommand(configureCmd)         // Adds the 'configure' command as a subcommand to the 'profile' command
}
