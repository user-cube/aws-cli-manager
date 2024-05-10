package cmd

import (
	"aws-cli-manager/aws"
	"aws-cli-manager/configurator"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "AWS Profile Manager is a tool to manage AWS CLI profiles",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available profiles",
	Run: func(cmd *cobra.Command, args []string) {
		aws.ListProfiles()
	},
}

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a profile",
	Args:  cobra.ArbitraryArgs,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetProfileNames(), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		aws.SelectProfile()
	},
}

var exportCredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Export credentials to environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		aws.ExportCredentialsToEnvironmentVariables()
	},
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure AWS CLI Profile",
	Run: func(cmd *cobra.Command, args []string) {
		configurator.ConfigureAWSCLI()
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	profileCmd.Flags()
	profileCmd.AddCommand(listCmd)
	profileCmd.AddCommand(selectCmd)
	profileCmd.AddCommand(exportCredentialsCmd)
	profileCmd.AddCommand(configureCmd)

}
