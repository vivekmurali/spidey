package main

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/vivekmurali/spidey/cmd"
)

func main() {
	f, err := os.Create("prof/cpu.prof")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	cmd.Execute()
}
