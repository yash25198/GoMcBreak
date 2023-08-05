package main

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	workDuration  = 25 * time.Minute // 25 minutes of work
	breakDuration = 5 * time.Minute  // 5 minutes break
	longBreak     = 15 * time.Minute // 15 minutes long break after 4 cycles
	cyclesBeforeLongBreak = 4
)

func main() {
	cycle := 0

	for {
		cycle++
		fmt.Printf("Cycle %d: Work for %s...\n", cycle, workDuration)
		timer := time.NewTimer(workDuration)
		<-timer.C

		notify("Time for a break!")

		if cycle%cyclesBeforeLongBreak == 0 {
			fmt.Printf("Long Break for %s...\n", longBreak)
			timer = time.NewTimer(longBreak)
		} else {
			fmt.Printf("Short Break for %s...\n", breakDuration)
			timer = time.NewTimer(breakDuration)
		}
		<-timer.C
		notify("break time up!")
	}
}

func notify(message string) {
	// Replace 'display notification' with the appropriate AppleScript for your version of macOS
	cmd := exec.Command("osascript", "-e", `display notification "`+message+`" with title "Time Tracker"`)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to send notification:", err)
	}
}
