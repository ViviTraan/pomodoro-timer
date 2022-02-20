package main

import (
	"fmt"
	"pomodoro-timer/timer"
	"time"

	"github.com/inancgumus/screen"
)

func welcomeMessage() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("The POMODORO timer will start in about 5 seconds\n1 beep means it's time for a pause of 5 minutes\n2 beeps means it's time to get back to work.\n 'CTRL + C' to exit the program - GOOD LUCK!")
	fmt.Println("-----------------------------------------------")
}

func main() {
	screen.MoveTopLeft()
	screen.Clear()
	welcomeMessage()
	time.Sleep(7 * time.Second)

	t := timer.Timer{
		Start:       time.Now(),
		WorkingMode: true,
	}
	prevElapsedTime := 0
	for {
		elapsedTime := t.GetElapsedTimeInSeconds()
		if elapsedTime != prevElapsedTime {
			t.PrintTimeRemaining(elapsedTime)
			prevElapsedTime = elapsedTime
			if t.ShouldSwitchMode(elapsedTime) {
				t.Alert()
				t.SwitchMode()
			}
		}
	}
}
