// Package cmd provides the command line interface for the AWS CLI Manager.
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/user-cube/aws-cli-manager/v2/pkg/profile"
)

// selectCmd represents the 'select' command.
var selectCmd = &cobra.Command{
	Use:   "select [profile]",
	Short: "Select an AWS profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// If no profile is specified, show the selection menu
			// SelectProfile returns false if user cancels with Ctrl+C
			profile.SelectProfile()
		} else {
			// If a profile is specified, select it directly
			profileName := args[0]
			err := profile.SelectProfileByName(profileName)
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
				return
			}
		}
	},
}

func init() {
	profileCmd.AddCommand(selectCmd) // Adds the 'select' command as a subcommand to the 'profile' command
}
