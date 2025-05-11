# AWS CLI Manager

This project is a command-line tool written in Go that helps manage the AWS CLI on your system. It checks if the AWS CLI is installed, and if not, it downloads and installs the correct version based on your system's architecture and operating system.

![demo](/demo/aws-cli-manager-demo.gif)

## Features

- Detects if AWS CLI is installed
- Downloads and installs AWS CLI if not present
- Supports both Linux and MacOS systems
- Manage multiple AWS profiles with ease
- Interactive and colorful UI for better user experience
- Track and switch between profiles quickly
- Test AWS connections

## Build from source

To install this tool, you can clone the repository and build the project:

```bash
git clone https://github.com/user-cube/aws-cli-manager.git
cd aws-cli-manager
go build
```

## Install with go install

```bash
go install github.com/user-cube/aws-cli-manager/v2@latest
```

## Install from compiled binary

You can also download the compiled binary from the [releases](https://github.com/user-cube/aws-cli-manager/releases/latest) page.

Please change `VERSION` and `ARCH` to the desired version and architecture before running the following command:
```bash
VERSION=1.2.0
ARCH=linux_amd64
wget https://github.com/user-cube/aws-cli-manager/releases/download/$VERSION/aws-cli-manager_$VERSION_$ARCH.tar.gz
```

After downloading the tarball, extract the contents and run the tool:
```bash
tar -xvf aws-cli-manager_$VERSION_$ARCH.tar.gz
sudo cp aws-cli-manager /usr/local/bin/aws-cli-manager
```

## Usage
After building the project, you can run the tool with:
```bash
./aws-cli-manager
```

If you copied the binary to `/usr/local/bin`, you can run the tool with:
```bash
aws-cli-manager
```

## Commands

### Profile Management
- `aws-cli-manager` - Interactive profile selection
- `aws-cli-manager profile add` - Add a new AWS profile
- `aws-cli-manager profile select` - Select a profile interactively
- `aws-cli-manager profile select [NAME]` - Select a specific profile by name
- `aws-cli-manager current` - Show the currently selected profile

### Installation and Configuration
- `aws-cli-manager install` - Install AWS CLI if not already installed
- `aws-cli-manager test` - Test the connection to AWS
- `aws-cli-manager completion [SHELL]` - Generate shell completion scripts

## UI Features

The AWS CLI Manager now features an enhanced UI with:

- Color-coded outputs for better readability
- Interactive prompts with improved selection indicators (→)
- Profile highlighting to show your current active profile
- Formatted tables for profile information
- Clearer success/error messaging

## Demo Scripts

The `demo` directory contains several scripts to demonstrate the features of AWS CLI Manager:

### Screenshot Demos
To take screenshots for documentation or presentations, you can use these scripts:

```bash
# Set up demo scripts
cd demo
chmod +x *.sh

# Generate dummy AWS profiles (safe for screenshots)
./generate_dummy_profiles.sh

# Run the dummy profile screenshot script
./make_demo.sh
```

The dummy profiles include:
- `demo-dev` (us-east-1) - Basic profile
- `demo-staging` (eu-west-1) - Initially selected profile
- `demo-prod` (us-west-2) - SSO-enabled profile
- `demo-admin` (us-east-2) - SSO-enabled profile

These profiles use fake credentials and are safe to include in screenshots.

## Autocompletion
To enable autocompletion for the tool, you can run the following command:
```bash
echo 'source <(aws-cli-manager completion SHELL)' >> ~/.SHEELrc
```
Please replace `SHELL` with the shell you are using (bash, zsh, fish).

If you don't to add the command to your shell configuration file, you can run the following command:
```bash
aws-cli-manager  completion SHELL -h
```

This will output a command that you can run to enable autocompletion without modifying your shell configuration file.