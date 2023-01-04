//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
)

func main() {
	// Create a process
	c := make(chan os.Signal,1)
	signal.Notify(c, os.Interrupt)
	proc := MockProcess{}
	// Run the process (blocking)
	go func() {
		for range c {		
			proc.Stop()
		}
		
	}()

	go func() {
		for range c {
			os.Exit(0)
		}
	}()
	proc.Run()
	

}
