package main

import (
  "fmt"
  "time"
)

func main() {
  tick := time.Tick(1 * time.Second)
  tick1 := time.Tick(2 * time.Second)

  go forPrint(tick)
  go rangePrint(tick1)

  time.Sleep(10 * time.Second)
}

func forPrint(tick <-chan time.Time) {
  for {
    now := <-tick

    fmt.Println("forPrint: ", now)
  }
}

func rangePrint(tick <-chan time.Time) {
  for now := range tick {
    fmt.Println("rangePrint: ", now)
  }
}
