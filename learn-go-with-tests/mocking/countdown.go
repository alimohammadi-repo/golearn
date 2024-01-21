package mocking

import (
	"fmt"
	"io"
	"time"
)

//func Countdown(buffer *bytes.Buffer) {
//
//	fmt.Fprint(buffer,"3")
//
//}

//func Countdown(buffer io.Writer) {
//
//	fmt.Fprint(buffer, "3")
//
//}

//func Countdown(buffer io.Writer) {
//
//	for i := 3; i > 0; i-- {
//		fmt.Fprintln(buffer, i)
//	}
//
//	fmt.Fprint(buffer, "GO!")
//
//}

const finalWord = "GO!"
const countdownStart = 3

//func Countdown(out io.Writer) {
//	for i := countdownStart; i > 0; i-- {
//		fmt.Fprintln(out, i)
//	}
//	fmt.Fprint(out, finalWord)
//}

//func Countdown(out io.Writer, sleeper Sleeper) {
//	for i := countdownStart; i > 0; i-- {
//		fmt.Fprintln(out, i)
//		sleeper.Sleep()
//
//	}
//	fmt.Fprint(out, finalWord)
//}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s SpySleeper) Sleep() {
	s.Calls++

}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()

	}

	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	//}

	fmt.Fprint(out, finalWord)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {

	c.sleep(c.duration)

}
