# AWS CLI Manager

This project is a command-line tool written in Go that helps manage the AWS CLI on your system. It checks if the AWS CLI is installed, and if not, it downloads and installs the correct version based on your system's architecture and operating system.

## Features

- Detects if AWS CLI is installed
- Downloads and installs AWS CLI if not present
- Supports both Linux and MacOS systems

## Installation

To install this tool, you can clone the repository and build the project:

```bash
git clone https://github.com/user-cube/aws-cli-manager.git
cd aws-cli-manager
go build
```

## Usage
After building the project, you can run the tool with:
```bash
./aws-cli-manager
```