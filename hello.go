package main

import "fmt"

func main() {
  x := 5
  x = "string"
  if x < 6 {
	   fmt.Printf("hello, %d world\n", x)
  }
}
