// Package cmd provides the command line interface for the AWS CLI Manager.
package cmd

import (
	"github.com/spf13/cobra" // Importing the cobra package for creating CLI applications
	"github.com/user-cube/aws-cli-manager/v2/pkg/aws"
	"github.com/user-cube/aws-cli-manager/v2/pkg/profile"
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

// addCmd represents the 'add' command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new profile",
	Run: func(cmd *cobra.Command, args []string) {
		aws.AddProfile() // Calls the AddProfile function from the aws package
	},
}

// useCmd represents the 'use' command.
var useCmd = &cobra.Command{
	Use:   "use [profile]",
	Short: "Use a specified AWS profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		err := profile.SelectProfileByName(profileName)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}
	},
}

// init function to initialize the root command and add subcommands.
func init() {
	rootCmd.AddCommand(profileCmd) // Adds the 'profile' command as a subcommand to the root command

	profileCmd.Flags()             // Initializes flags for the 'profile' command
	profileCmd.AddCommand(listCmd) // Adds the 'list' command as a subcommand to the 'profile' command
	profileCmd.AddCommand(addCmd)  // Adds the 'add' command as a subcommand to the 'profile' command
	profileCmd.AddCommand(useCmd)  // Adds the 'use' command as a subcommand to the 'profile' command
}
