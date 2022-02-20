package timer

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
)

type Timer struct {
	Start       time.Time
	WorkingMode bool
}

//Will get the elasped time in seconds
func (t Timer) GetElapsedTimeInSeconds() int {
	return int(time.Since(t.Start).Seconds())
}

//function to switch mode
func (t *Timer) SwitchMode() {
	t.Start = time.Now()
	t.WorkingMode = !t.WorkingMode
}

//alert function will show a message depending on which mode it's on
func (t Timer) Alert() {
	message := " - It's time to take a break"
	t.BeepAlert(message)
	if !t.WorkingMode {
		message = " - New session starting"
	}
	fmt.Println(message)

}

//beepAlert function will beep once when it's time for a pause and two times when its time to get back to work
func (t Timer) BeepAlert(message string) {
	os := runtime.GOOS
	if os == "darwin" {
		go exec.Command("say", message).Output()
	}
	if !t.WorkingMode {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	} else {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	}

}

const sessionDuration = 25 * 60 // 25 minutes
const pausDuration = 5 * 60     // 15 minutes

//switches mode depending on the elapsed time
func (t Timer) ShouldSwitchMode(elapsedTime int) bool {
	return elapsedTime == t.GetDuration()
}

//setting the duration variable depending on the mode
func (t Timer) GetDuration() int {
	duration := sessionDuration
	if !t.WorkingMode {
		duration = pausDuration
	}
	return duration
}

//printing out which mode you're on
func (t Timer) GetMode() string {
	mode := "WORK MODE"
	if !t.WorkingMode {
		mode = "PAUS TIME"
	}
	return mode
}

//printing out and formatting time
func (t Timer) PrintTimeRemaining(elapsed int) {
	timeRemaining := t.GetDuration() - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining % 60
	fmt.Printf("\r%v: %02d:%02d ", t.GetMode(), minutes, seconds)
}
