package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	SleepDuration time.Duration
	SleepFunc     func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.SleepDuration)
}

func Countdown(sleeper Sleeper, writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}
