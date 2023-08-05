# Promodoro timer for tracking work or study
- Din't wanna download an app for this 
1. `osascript -e 'display notification "Notification Enabled" with title "Time Tracker"'` to verify if Applescript is enabled.
    Please note that for macOS, the code uses the osascript command with AppleScript to display the notification. The notify function takes care of sending notifications using the os/exec package. Make sure you have AppleScript enabled on your macOS system for this to work.
2. `./notify`

