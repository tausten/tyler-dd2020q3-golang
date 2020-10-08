package main

import (
	"os"
	"time"

	"github.com/tausten/tyler-dd2020q3-golang/mocking"
)

func other_main() {
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(sleeper, os.Stdout)
}
