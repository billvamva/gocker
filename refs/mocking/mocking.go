package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWord = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}

func Countdown(b io.Writer, s Sleeper) {
	for i:=countdownStart; i>0; i--{
		fmt.Fprintln(b, i)
		s.Sleep()
	}
	fmt.Fprint(b, finalWord)
}
