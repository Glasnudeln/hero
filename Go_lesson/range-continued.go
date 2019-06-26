package main

import "fmt"

func main() {
  pow := make([]int, 16)
  for i := range pow {
    pow[i] = 2 << uint(i) // == 2**i
  }
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
