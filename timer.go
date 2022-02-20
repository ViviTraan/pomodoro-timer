package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
)

type timer struct {
	start       time.Time
	workingMode bool
}

//Will get the elasped time in seconds
func (t timer) getElapsedTimeInSeconds() int {
	return int(time.Since(t.start).Seconds())
}

//function to switch mode
func (t *timer) switchMode() {
	t.start = time.Now()
	t.workingMode = !t.workingMode
}

//alert function will show a message depending on which mode it's on
func (t timer) alert() {
	message := " - It's time to take a break"
	t.beepAlert(message)
	if !t.workingMode {
		message = " - New session starting"
	}
	fmt.Println(message)

}

//beepAlert function will beep once when it's time for a pause and two times when its time to get back to work
func (t timer) beepAlert(message string) {
	os := runtime.GOOS
	if os == "darwin" {
		go exec.Command("say", message).Output()
	}
	if !t.workingMode {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	} else {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	}

}

const sessionDuration = 25 * 60 // 25 minutes
const pausDuration = 5 * 60     // 15 minutes

//switches mode depending on the elapsed time
func (t timer) shouldSwitchMode(elapsedTime int) bool {
	return elapsedTime == t.getDuration()
}

//setting the duration variable depending on the mode
func (t timer) getDuration() int {
	duration := sessionDuration
	if !t.workingMode {
		duration = pausDuration
	}
	return duration
}

//printing out which mode you're on
func (t timer) getMode() string {
	mode := "WORK MODE"
	if !t.workingMode {
		mode = "PAUS TIME"
	}
	return mode
}

//printing out and formatting time
func (t timer) printTimeRemaining(elapsed int) {
	timeRemaining := t.getDuration() - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining % 60
	fmt.Printf("\r%v: %02d:%02d ", t.getMode(), minutes, seconds)
}
