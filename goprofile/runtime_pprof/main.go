package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
  "time"
)

func main() {
	var (
		isCPUPprof bool
		isMemPorof bool
	)
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPorof, "mem", false, "turn mem pprof on")
  flag.Parse()
	if isCPUPprof {
		file, err := os.Create("./cpu.goprofile")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(file)

		defer pprof.StopCPUProfile()
	}
  for i := 0; i < 8; i++ {
    go loginCode()
  }
  time.Sleep(20 * time.Second)
	if isMemPorof {
		file1, err := os.Create("./mem.goprofile")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(file1)

		file1.Close()
	}
}

func loginCode() {
	var c chan int
	for {
		select {
		case x := <-c:
			fmt.Printf("recv from chan, value:%v\n", x)
		default:
      time.Sleep(time.Second)
		}
	}
}
