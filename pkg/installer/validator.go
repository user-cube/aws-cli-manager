// Package installer provides functions to validate and install AWS CLI.
package installer

import (
	"aws-cli-manager/pkg/settings"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// InstallAWSCLI checks if AWS CLI is installed and installs it if not.
func InstallAWSCLI() {
	fmt.Println(checkIfAWSCLIIsInstalled())
}

// checkDependencies checks if curl and unzip are installed.
func checkDependencies() {
	checkIfCurlIsInstalled()
	checkIfUnzipIsInstalled()
}

// checkIfAWSCLIIsInstalled checks if AWS CLI is installed.
// If not, it installs AWS CLI.
func checkIfAWSCLIIsInstalled() string {
	// We need to check if aws cli is installed on the system
	cmd := exec.Command("aws", "--version")
	err := cmd.Run()

	if err != nil {
		var exitError *exec.ExitError
		if !errors.As(err, &exitError) {
			fmt.Println("AWS CLI is not installed on your system")

			checkDependencies()

			// Ask user if he wants to install it
			installAWSCLI()

			return ""
		}
	} else {
		return "AWS CLI is installed on your system"
	}

	// Add this return statement
	return "AWS CLI is installed on your system"
}

// checkIfCurlIsInstalled checks if curl is installed.
func checkIfCurlIsInstalled() {
	cmd := exec.Command("curl", "--version")
	err := cmd.Run()

	if err != nil {
		var exitError *exec.ExitError
		if !errors.As(err, &exitError) {
			message := fmt.Errorf(
				"curl is not installed on your system, please install it and execute this program again",
			)
			log.Fatalf("%v", message)
		}
	}
}

// checkIfUnzipIsInstalled checks if unzip is installed.
func checkIfUnzipIsInstalled() {
	cmd := exec.Command("unzip", "--version")
	err := cmd.Run()

	if err != nil {
		var exitError *exec.ExitError
		if !errors.As(err, &exitError) {
			message := fmt.Errorf(
				"unzip is not installed on your system, please install it and execute this program again",
			)
			log.Fatal(message)
		}
	}
}

// selectDownloadingURLAccordingToArch selects the downloading URL according to the system architecture.
func selectDownloadingURLAccordingToArch() string {
	// We need to detect if this is a 32-bit or 64-bit system to download the correct AWS CLI package

	arch := runtime.GOARCH

	if arch == "amd64" {
		return fmt.Sprintf("%s%s", settings.BaseLinuxInstallerUrl, "-x86_64.zip")
	} else {
		return fmt.Sprintf("%s%s", settings.BaseLinuxInstallerUrl, "-aarch64.zip")
	}
}

// detectOS detects the operating system.
func detectOS() string {
	// We need to detect if this is a linux or Mac system to install aws cli accordingly

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("MacOS System detected, proceeding to install AWS CLI")
		return "darwin"
	case "linux":
		fmt.Println("Linux System detected, proceeding to install AWS CLI")
		return "linux"
	case "windows":
		return "windows"
	default:
		message := fmt.Errorf("unsupported OS detected, please use MacOS, Linux or Windows")
		log.Fatalf("%v", message)
	}

	return ""
}

// installAWSCLILinux installs AWS CLI on Linux.
func installAWSCLILinux() {

	url := selectDownloadingURLAccordingToArch()

	downloadCmd := exec.Command("curl", "-o", "awscliv2.zip", url)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		message := fmt.Errorf("error downloading aws-cli: %v", err)
		log.Fatalf("%v", message)
	}

	// Unzip AWS CLI
	unzipCmd := exec.Command("unzip", "awscliv2.zip")
	unzipCmd.Stdout = os.Stdout
	unzipCmd.Stderr = os.Stderr
	if err := unzipCmd.Run(); err != nil {
		message := fmt.Errorf("error unzip aws-cli: %v", err)
		log.Fatalf("%v", message)
	}

	// Install AWS CLI
	installCmd := exec.Command("sudo", "-S", "--", "./aws/install")
	installCmd.Stdin = os.Stdin
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	fmt.Println("Installing AWS CLI...")

	if err := installCmd.Run(); err != nil {
		message := fmt.Errorf("error installing aws-cli: %v", err)
		log.Fatalf("%v", message)
	}

	// Remove downloaded zip file
	if err := os.Remove("awscliv2.zip"); err != nil {
		message := fmt.Errorf("error deleting aws zip file: %v", err)
		log.Fatalf("%v", message)
	}

	// Remove extracted AWS CLI folder
	if err := os.RemoveAll("aws"); err != nil {
		message := fmt.Errorf("error deleting aws folder: %v", err)
		log.Fatalf("%v", message)
	}
}

// installAWSCLIMac installs AWS CLI on MacOS.
func installAWSCLIMac() {
	// Download AWS CLI package
	downloadCmd := exec.Command(
		"curl",
		"-o",
		"AWSCLIV2.pkg",
		"https://awscli.amazonaws.com/AWSCLIV2.pkg",
	)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		message := fmt.Errorf("error downloading aws cli: %v", err)
		log.Fatalf("%v", message)
	}

	// Install AWS CLI package
	installCmd := exec.Command(
		"sudo",
		"installer",
		"-pkg",
		"AWSCLIV2.pkg",
		"-target",
		"/",
	)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		message := fmt.Errorf("error installing aws cli: %v", err)
		log.Fatalf("%v", message)
	}

	// Remove downloaded package
	if err := os.Remove("AWSCLIV2.pkg"); err != nil {
		message := fmt.Errorf("error removing AWSCLIV2.pkg: %v", err)
		log.Fatalf("%v", message)
	}
}

// installAWSCLIWindows installs AWS CLI on Windows.
func installAWSCLIWindows() {

	// Install AWS CLI package
	installCmd := exec.Command(
		"msiexec",
		"/i",
		"https://awscli.amazonaws.com/AWSCLIV2.msi",
	)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		message := fmt.Errorf("error installing aws cli: %v", err)
		log.Fatalf("%v", message)
	}
}

// installAWSCLI detects the operating system and installs AWS CLI accordingly.
func installAWSCLI() {

	// We need to detect if this is a Linux or macOS system to install aws cli accordingly

	osSystem := detectOS()

	if osSystem == "darwin" {
		installAWSCLIMac()
	} else if osSystem == "linux" {
		installAWSCLILinux()
	} else if osSystem == "windows" {
		installAWSCLIWindows()
	} else {
		message := fmt.Errorf("unsupported OS detected, please use MacOS, Linux or Windows")
		log.Fatalf("%v", message)
	}

	fmt.Println("AWS CLI installed successfully")
}
