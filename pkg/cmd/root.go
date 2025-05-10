// Package cmd provides the command line interface for the AWS CLI Manager.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra" // Importing the cobra package for creating CLI applications
	"github.com/user-cube/aws-cli-manager/v2/pkg/aws"
	"github.com/user-cube/aws-cli-manager/v2/pkg/config"
	"github.com/user-cube/aws-cli-manager/v2/pkg/installer"
	"github.com/user-cube/aws-cli-manager/v2/pkg/profile"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "aws-cli-manager",
	Short: "AWS CLI Manager is a tool to manage AWS CLI profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profile.SelectProfile()
	},
}

// versionCmd represents the 'version' command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		config.LogVersion()
	},
}

// installCmd represents the 'install' command.
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install AWS CLI tool from AWS official website",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing AWS CLI")
		installer.InstallAWSCLI() // Calls the InstallAWSCLI function from the installer package
	},
}

var testConnectionCmd = &cobra.Command{
	Use:   "test",
	Short: "Test connection to AWS",
	Run: func(cmd *cobra.Command, args []string) {
		aws.TestConnection() // Calls the TestConnection function from the aws package
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute() // Executes the root command
	if err != nil {
		os.Exit(1) // Exits the program if there is an error
	}
}

// init function to initialize the root command and add subcommands.
func init() {
	rootCmd.Flags()                       // Initializes flags for the root command
	rootCmd.AddCommand(versionCmd)        // Adds the 'version' command as a subcommand to the root command
	rootCmd.AddCommand(installCmd)        // Adds the 'install' command as a subcommand to the root command
	rootCmd.AddCommand(testConnectionCmd) // Adds the 'testConnection' command as a subcommand to the root command
}
