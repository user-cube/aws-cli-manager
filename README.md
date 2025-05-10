# AWS CLI Manager

This project is a command-line tool written in Go that helps manage the AWS CLI on your system. It checks if the AWS CLI is installed, and if not, it downloads and installs the correct version based on your system's architecture and operating system.

## Features

- Detects if AWS CLI is installed
- Downloads and installs AWS CLI if not present
- Supports both Linux and MacOS systems

## Build from source

To install this tool, you can clone the repository and build the project:

```bash
git clone https://github.com/user-cube/aws-cli-manager.git
cd aws-cli-manager
go build
```

## Install with go install

```bash
go install github.com/user-cube/aws-cli-manager@latest
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