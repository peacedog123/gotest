package main

import (
  "fmt"
  "log"
  "net/rpc"
  "os"

  "github.com/peacedog123/gotest/server"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: ", os.Args[0], "server address")
    os.Exit(1)
  }

  service := os.Args[1]

  client, err := rpc.Dial("tcp", service)
  if err != nil {
    log.Fatal("dialing: ", err)
  }

  args := server.Args{17, 8}
  var reply int
  err = client.Call("Arith.Multiply", args, &reply)
  if err != nil {
    log.Fatal("arith error: ", err)
  }
  fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

  var quot server.Quotient
  err = client.Call("Arith.Divide", args, &quot)
  if err != nil {
    log.Fatal("arith error:", err)
  }
  fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
