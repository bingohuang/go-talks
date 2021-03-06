// +build OMIT

package main

import (
	"os"
	"runtime/trace"
	"time"
)

// START OMIT
func leaker() {
	time.Sleep(1000000 * time.Minute)
}

func main() {
	trace.Start(os.Stderr)
	for i := 0; i < 1000; i++ {
		go leaker()
	}
	trace.Stop()
}
// STOP OMIT
