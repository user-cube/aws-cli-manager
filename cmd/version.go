package cmd

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Version information
var (
	BuildDate = "unknown"
	GitCommit = "unknown"
	Version   = "v2.0.0"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of AWS CLI Manager",
	Long: `Display the version, build date, and git commit of your AWS CLI Manager installation.

Examples:
  # Show version information
  aws-cli-manager version`,
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		fmt.Printf("%s: %s\n", cyan("AWS CLI Manager Version"), Version)
		fmt.Printf("%s: %s\n", cyan("Git Commit"), GitCommit)
		fmt.Printf("%s: %s\n", cyan("Built"), BuildDate)
		fmt.Printf("%s: %s\n", cyan("Platform"), runtime.GOOS+"/"+runtime.GOARCH)
		fmt.Printf("%s: %s\n", cyan("Go Version"), runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
