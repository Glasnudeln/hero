package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 60000; {
    sum += sum
  }
  fmt.Println(sum)
}
