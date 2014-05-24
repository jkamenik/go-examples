package main

import (
  "io"
  "log"
  "net"
)

const listen = "localhost:4000"

func main() {
  println("Starting")
  l, err := net.Listen("tcp", listen)
  if err != nil {
    log.Fatal(err)
  }
  println("Listening on ", listen)

  for {
    println("Waiting for connection")
    c, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }

    println("Connection received")

    go io.Copy(c, c)
  }
}
