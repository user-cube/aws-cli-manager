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

func checkDependencies() {
	checkIfCurlIsInstalled()
	checkIfUnzipIsInstalled()
}

func checkIfAWSCLIIsInstalled() {
	// We need to check if aws cli is installed on the system
	cmd := exec.Command("aws", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Println("AWS CLI is not installed on your system")

			checkDependencies()

			// Ask user if he wants to install it
			installAWSCLI()
		}
	}
}

func checkIfCurlIsInstalled() {
	cmd := exec.Command("curl", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Println("Curl is not installed on your system, please install it and execute this program again")
			fmt.Println("Exiting...")
			panic("Missing curl on your system")
		}
	}
}

func checkIfUnzipIsInstalled() {
	cmd := exec.Command("unzip", "--version")
	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Println("Unzip is not installed on your system, please install it and execute this program again")
			fmt.Println("Exiting...")
			panic("Missing unzip on your system")
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
		panic("Unsupported OS detected")
	}
}

func installAWSCLILinux() {

	url := selectDownloadingURLAccordingToArch()

	downloadCmd := exec.Command("curl", "-o", "awscliv2.zip", url)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		panic(err)
	}

	// Unzip AWS CLI
	unzipCmd := exec.Command("unzip", "awscliv2.zip")
	unzipCmd.Stdout = os.Stdout
	unzipCmd.Stderr = os.Stderr
	if err := unzipCmd.Run(); err != nil {
		panic(err)
	}

	// Install AWS CLI
	installCmd := exec.Command("sudo", "-S", "--", "./aws/install")
	installCmd.Stdin = os.Stdin
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	fmt.Println("Installing AWS CLI...")

	if err := installCmd.Run(); err != nil {
		panic(err)
	}

	// Remove downloaded zip file
	if err := os.Remove("awscliv2.zip"); err != nil {
		panic(err)
	}

	// Remove extracted AWS CLI folder
	if err := os.RemoveAll("aws"); err != nil {
		panic(err)
	}
}

func installAWSCLIMac() {
	// Download AWS CLI package
	downloadCmd := exec.Command("curl", "-o", "AWSCLIV2.pkg", "https://awscli.amazonaws.com/AWSCLIV2.pkg")
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		panic(err)
	}

	// Install AWS CLI package
	installCmd := exec.Command("sudo", "installer", "-pkg", "AWSCLIV2.pkg", "-target", "/")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		panic(err)
	}

	// Remove downloaded package
	if err := os.Remove("AWSCLIV2.pkg"); err != nil {
		panic(err)
	}
}

func installAWSCLI() {

	// We need to detect if this is a Linux or MacOs system to install aws cli accordingly

	osSystem := detectOS()

	if osSystem == "darwin" {
		installAWSCLIMac()
	} else {
		installAWSCLILinux()
	}

	fmt.Println("AWS CLI installed successfully")
}
