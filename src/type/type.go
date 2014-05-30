package main

import "fmt"

type Myint int

func (m Myint) String() string {
  return fmt.Sprintf("%d-int", m)
}

func main() {
  m := Myint(3)

  fmt.Println(m)
  print(int(m))
}

func print(x int) {
  fmt.Println(x)
}
