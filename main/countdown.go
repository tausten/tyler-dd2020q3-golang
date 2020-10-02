package main

import (
	"os"
	"time"

	"github.com/tausten/tyler-dd2020q3-golang/mocking"
)

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func other_main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(sleeper, os.Stdout)
}
