package main

import (
	"aws-cli-manager/menus"
	"aws-cli-manager/validators"
)

func main() {
	validators.ValidateAWSCLI()

	menus.Menu()
}
