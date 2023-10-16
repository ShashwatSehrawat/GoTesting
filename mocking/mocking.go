package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

// type defaultSleeper struct{}

// func (d *defaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
		// time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(writer, "Go!")

}
func main() {
	// sleeper := &defaultSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
