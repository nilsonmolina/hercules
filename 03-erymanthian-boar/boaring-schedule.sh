#!/bin/bash
# On macOS, enable atrun utility with the following code: 
# sudo launchctl load -w /System/Library/LaunchDaemons/com.apple.atrun.plist

echo "echo 'oink oink mother f*cker!' >> boar.txt | open boar.txt" | at 08:42 AM 12/21/2018
