package main

import (
  "fmt"
  "github.com/welcome112s/gogostudy/goprofile/http"
  "os"
  "os/signal"
  "sync"
  "syscall"
)

var mutex  sync.Mutex
var once   sync.Once
var wg sync.WaitGroup

var   num int

func main() {
  for  i:=0;i<1000;i++{
    wg.Add(1)
    go add (&wg)
  }
  wg.Wait()
  fmt.Println("num =", num)

  http.HttpProfile()

  c:=make(chan os.Signal,1)

  signal.Notify(c,syscall.SIGHUP,syscall.SIGQUIT,syscall.SIGTERM,syscall.SIGINT,os.Interrupt)

  for s := range c {
    switch s {
    case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT,os.Interrupt:
      os.Exit(0)
    case syscall.SIGHUP:
    default:
      return
    }
    fmt.Println("Got signal:", s)
  }
}


func add(wg *sync.WaitGroup){
  mutex.Lock()
    num ++
  mutex.Unlock()
    wg.Done()
}


var ad=func() {
  fmt.Println("adadadas")
}

func onceod(){
     fmt.Println("aaaa")
     wg.Done()
}
func do(a *int ){
  defer  mutex.Unlock()
  mutex.Lock()
  *a ++
  println("adads")
}
