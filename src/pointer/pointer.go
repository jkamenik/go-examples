package main

import "fmt"

type S string

func (s *S) Value() string { return string(*s) }
func (s S) Copy() *S       { return &s }

func main() {
  myS := new(S)
  s := S("Testing")

  fmt.Printf("%v %s\n", myS, myS)
  fmt.Println(myS.Value())
  fmt.Printf("%v\n", myS.Copy())
  fmt.Println("-----")
  fmt.Printf("%v %s\n", &s, s)
  fmt.Println(s.Value())
  fmt.Printf("%v\n", s.Copy())
}
