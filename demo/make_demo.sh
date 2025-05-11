#!/bin/bash
# Run demo_script and record
asciinema rec aws-cli-manager.cast -c "./demo_script.sh" --overwrite

# Convert to GIF
asciicast2gif aws-cli-manager.cast aws-cli-manager.gif

# Optimize GIF
gifsicle -O3 --colors 256 aws-cli-manager.gif -o aws-cli-manager-demo.gif

echo "✅ Done! Your demo GIF is ready as aws-cli-manager-demo.gif"