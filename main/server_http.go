package main

import (
  _ "fmt"
  "log"
  "net/http"
  "net/rpc"

  "github.com/peacedog123/gotest/server"
)

func main() {
  arith := new(server.Arith)
  rpc.Register(arith)
  rpc.HandleHTTP()

  err := http.ListenAndServe(":1234", nil)
  if err != nil {
    log.Fatal("error: ", err)
  }
}
