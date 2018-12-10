package main

import (
  "fmt"
  _ "log"
  "net"
  "net/rpc"
  "os"

  "github.com/peacedog123/gotest/server"
)

func checkError(err error) {
  if err != nil {
    fmt.Println("Fatal error ", err.Error())
    os.Exit(1)
  }
}

func main() {
  arith := new(server.Arith)
  rpc.Register(arith)

  tcpAddr, err := net.ResolveTCPAddr("tcp", ":4321")
  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    rpc.ServeConn(conn)
  }
}
