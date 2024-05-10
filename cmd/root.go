package cmd

import (
	"aws-cli-manager/validators"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "1.2.0"

var rootCmd = &cobra.Command{
	Use:   "aws-cli-manager",
	Short: "AWS CLI Manager is a tool to manage AWS CLI profiles",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AWS CLI Manager version " + version)
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install AWS CLI tool from AWS official website",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing AWS CLI")
		validators.InstallAWSCLI()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(installCmd)
}
