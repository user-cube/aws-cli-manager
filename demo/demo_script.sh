#!/bin/bash
# Green Profile Highlight Demo Script
# This script focuses on the green highlighting of the current profile
# in both the selection menu and success messages

# Enable script to be executable
# chmod +x green_profile_demo.sh

# Clear terminal and set up
# 1. First show the current profile
echo "First, let's see the current profile with green highlighting:"
echo "$ aws-cli-manager current"
aws-cli-manager current
echo
echo "Press Enter to continue..."
read

# 2. Show profile listing with current profile highlighted in green
clear
echo "Now, let's see the profile list with the current profile highlighted:"
echo "$ aws-cli-manager profile list"
aws-cli-manager profile list
echo
echo "Press Enter to continue..."
read

# 3. Show the selection interface with current profile in green
clear
echo "Next, let's see the selection interface with the current profile in green:"
echo "$ aws-cli-manager"
echo "Note: The current profile and '(current)' text should both be in green."
echo "The arrow (→) indicates the selected item."
echo "Take a screenshot during selection."
echo
echo "Press Enter to start profile selection..."
read
aws-cli-manager
echo
echo "Press Enter to continue..."
read

# 4. Show the success message with the profile name in green
clear
echo "Notice the success message format:"
echo "'Profile [profile-name] set successfully'"
echo "Where [profile-name] should be in green, but 'Profile' and 'set successfully' are in white."
echo
echo "Press Enter to see the current profile again..."
read
aws-cli-manager current
echo
