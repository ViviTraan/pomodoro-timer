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

func (t timer) getElapsedTimeInSeconds() int {
	return int(time.Since(t.start).Seconds())
}

func (t *timer) switchMode() {
	t.start = time.Now()
	t.workingMode = !t.workingMode
}

func (t timer) alert() {
	message := " - It's time to take a break"
	t.beepAlert(message)
	if !t.workingMode {
		message = " - New session starting"
	}
	fmt.Println(message)

}

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

func (t timer) shouldSwitchMode(elapsedTime int) bool {
	return elapsedTime == t.getDuration()
}

func (t timer) getDuration() int {
	duration := sessionDuration
	if !t.workingMode {
		duration = pausDuration
	}
	return duration
}

func (t timer) getMode() string {
	mode := "WORK MODE"
	if !t.workingMode {
		mode = "PAUS TIME"
	}
	return mode
}

func (t timer) printTimeRemaining(elapsed int) {
	timeRemaining := t.getDuration() - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining % 60
	fmt.Printf("\r%v: %02d:%02d ", t.getMode(), minutes, seconds)
}
