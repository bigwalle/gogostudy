package goprofile

import (
  "log"
  "os"
  "runtime/pprof"
)

func CPUProfile() {

  profileFile:="cpu.goprofile"

  f,err:=os.Create(profileFile)
  if err != nil {
    log.Fatalf("%v", err)
  }
  err = pprof.StartCPUProfile(f)
  if err!=nil{
    log.Fatalf("%v", err)
  }

  pprof.StopCPUProfile()
}

func HeapProfile(){
    heapprofileFile:="heap.goprofile"

    f,err:=os.Create(heapprofileFile)
    if err != nil {
    log.Fatalf("%v", err)
    }
    pprof.WriteHeapProfile(f)

}


