package http

import (
  "log"
  "net/http"
  "net/http/pprof"
)

func HttpProfile(){
  mutex:=http.NewServeMux()
  mutex.HandleFunc("/debug/pprof/",pprof.Index)
  mutex.HandleFunc("/debug/pprof/cmdline",pprof.Cmdline)
  mutex.HandleFunc("/debug/pprof/goprofile",pprof.Profile)
  mutex.HandleFunc("/debug/pprof/symbol",pprof.Symbol)
  mutex.HandleFunc("/debug/pprof/trace",pprof.Trace)
  log.Fatal(http.ListenAndServe("0.0.0.0:8080",mutex))
}
