package main

import (
  "fmt"
  "os"
  "time"
)

/*
 Generator \
            Fan in -> output
 Generator /
*/
func main() {
  done := make(chan bool)
  gen1 := generator("Gen1", 1, 30, done)
  gen2 := generator("Gen2", 2, 15, done)
  gen3 := generator("Gen3", 3, 10, done) // => 60 iterations

  merged := fanIn(gen1, gen2, gen3)

  for {
    select {
    case <-done:
      fmt.Println("First one done, quitting")
      os.Exit(0)
    case value := <-merged:
      fmt.Println(value)
    }
  }
}

/*
 Takes params and returns an output channel
*/
func generator(name string, sleep int64, count int, done chan bool) <-chan string {
  output := make(chan string)
  duration := time.Duration(sleep) * time.Second

  go func() {
    for i := 0; i < count; i++ {
      output <- fmt.Sprintf("%s %d", name, i)
      time.Sleep(duration)
    }
    done <- true
  }() // () are important

  return output
}

/*
 Aggragates N number of output channels into a single output channel
*/
func fanIn(channels ...<-chan string) <-chan string {
  out := make(chan string)
  output := func(c <-chan string) {
    for n := range c {
      out <- n
    }
  }

  for _, c := range channels {
    go output(c)
  }

  return out
}
