package validators

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ValidateAWSCLI() {
	checkIfAWSCLIIsInstalled()
}

func InstallAWSCLI() {
	fmt.Println(checkIfAWSCLIIsInstalled())
}

func checkDependencies() {
	checkIfCurlIsInstalled()
	checkIfUnzipIsInstalled()
}

func checkIfAWSCLIIsInstalled() string {
	// We need to check if aws cli is installed on the system
	cmd := exec.Command("aws", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Println("AWS CLI is not installed on your system")

			checkDependencies()

			// Ask user if he wants to install it
			installAWSCLI()

			return ""
		} else {
			return "AWS CLI is installed on your system"
		}
	} else {
		return "AWS CLI is installed on your system"
	}
}

func checkIfCurlIsInstalled() {
	cmd := exec.Command("curl", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			message := fmt.Errorf("curl is not installed on your system, please install it and execute this program again")
			fmt.Println(message)
			os.Exit(1)
		}
	}
}

func checkIfUnzipIsInstalled() {
	cmd := exec.Command("unzip", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			message := fmt.Errorf("unzip is not installed on your system, please install it and execute this program again")
			fmt.Println(message)
			os.Exit(1)
		}
	}
}

func selectDownloadingURLAccordingToArch() string {
	// We need to detect if this is a 32-bit or 64-bit system to download the correct AWS CLI package

	arch := runtime.GOARCH

	if arch == "amd64" {
		return "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"
	} else {
		return "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip"
	}
}

func detectOS() string {
	// We need to detect if this is a linux or Mac system to install aws cli accordingly

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("MacOS System detected, proceeding to install AWS CLI")
		return "darwin"
	case "linux":
		fmt.Println("Linux System detected, proceeding to install AWS CLI")
		return "linux"
	default:
		fmt.Println("Unsupported OS detected, please install AWS CLI manually")
		message := fmt.Errorf("unsupported OS detected, please use MacOS or Linux")
		fmt.Println(message)
		os.Exit(1)
	}

	return ""
}

func installAWSCLILinux() {

	url := selectDownloadingURLAccordingToArch()

	downloadCmd := exec.Command("curl", "-o", "awscliv2.zip", url)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		message := fmt.Errorf("error downloading aws-cli: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Unzip AWS CLI
	unzipCmd := exec.Command("unzip", "awscliv2.zip")
	unzipCmd.Stdout = os.Stdout
	unzipCmd.Stderr = os.Stderr
	if err := unzipCmd.Run(); err != nil {
		message := fmt.Errorf("error unzip aws-cli: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Install AWS CLI
	installCmd := exec.Command("sudo", "-S", "--", "./aws/install")
	installCmd.Stdin = os.Stdin
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	fmt.Println("Installing AWS CLI...")

	if err := installCmd.Run(); err != nil {
		message := fmt.Errorf("error installing aws-cli: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Remove downloaded zip file
	if err := os.Remove("awscliv2.zip"); err != nil {
		message := fmt.Errorf("error deleting aws zip file: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Remove extracted AWS CLI folder
	if err := os.RemoveAll("aws"); err != nil {
		message := fmt.Errorf("error deleting aws folder: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}
}

func installAWSCLIMac() {
	// Download AWS CLI package
	downloadCmd := exec.Command("curl", "-o", "AWSCLIV2.pkg", "https://awscli.amazonaws.com/AWSCLIV2.pkg")
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		message := fmt.Errorf("error downloading aws cli: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Install AWS CLI package
	installCmd := exec.Command("sudo", "installer", "-pkg", "AWSCLIV2.pkg", "-target", "/")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		message := fmt.Errorf("error installing aws cli: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}

	// Remove downloaded package
	if err := os.Remove("AWSCLIV2.pkg"); err != nil {
		message := fmt.Errorf("error removing AWSCLIV2.pkg: %v", err)
		fmt.Println(message)
		os.Exit(1)
	}
}

func installAWSCLI() {

	// We need to detect if this is a Linux or macOS system to install aws cli accordingly

	osSystem := detectOS()

	if osSystem == "darwin" {
		installAWSCLIMac()
	} else {
		installAWSCLILinux()
	}

	fmt.Println("AWS CLI installed successfully")
}
