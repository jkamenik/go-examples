package main

import "fmt"

func main() {
  slice1 := make([]int, 10)      //=> 10 size array
  slice2 := make([]int, 10, 100) //=> 100 size array, but a 10 size slice

  fmt.Println(slice1)
  fmt.Println(len(slice1))

  fmt.Println(slice2)
  fmt.Println(len(slice2))

  // No error because slice2's backing array is 100 large
  fmt.Println(slice2[0:100])
}
