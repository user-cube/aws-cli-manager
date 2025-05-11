// Package ui provides reusable UI components for the AWS CLI Manager.
package ui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// Colors defines standard colors used throughout the UI
var (
	SuccessColor = color.New(color.FgGreen).SprintFunc()
	ErrorColor   = color.New(color.FgRed).SprintFunc()
	InfoColor    = color.New(color.FgYellow).SprintFunc()
	HeaderColor  = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	NormalColor  = color.New(color.FgWhite).SprintFunc()
)

// SelectOptions contains options for select prompts
type SelectOptions struct {
	Label    string
	Items    []string
	Selected string
}

// CreateSelectPrompt creates a select prompt with standard styling
func CreateSelectPrompt(options SelectOptions) promptui.Select {
	return promptui.Select{
		Label: options.Label,
		Items: options.Items,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "→ {{ . | cyan }}",
			Inactive: "  {{ . }}", // No white formatting to allow for custom colors in items
			Selected: "✓ {{ . | green }}",
		},
	}
}

// TextPromptOptions contains options for text prompts
type TextPromptOptions struct {
	Label    string
	Default  string
	Validate func(string) error
}

// CreateTextPrompt creates a text prompt with standard styling
func CreateTextPrompt(options TextPromptOptions) promptui.Prompt {
	return promptui.Prompt{
		Label:    options.Label,
		Default:  options.Default,
		Validate: options.Validate,
	}
}

// PrintSuccess prints a success message in green
func PrintSuccess(message string, args ...interface{}) {
	color.Green(message, args...)
}

// PrintError prints an error message in red
func PrintError(message string, args ...interface{}) {
	color.Red(message, args...)
}

// PrintInfo prints an info message in yellow
func PrintInfo(message string, args ...interface{}) {
	color.Yellow(message, args...)
}

// PrintCurrentProfile prints information about the current profile
func PrintCurrentProfile(profileName, region string, ssoEnabled bool) {
	fmt.Printf("Current profile: %s\n", SuccessColor(profileName))
	color.White("Region: %s", region)
	if ssoEnabled {
		color.Green("SSO Enabled: Yes")
	} else {
		color.Red("SSO Enabled: No")
	}
}
