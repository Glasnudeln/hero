package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  j := 0
  for i := 1; i <= 10000000; i++ {
    z -= (z * z - x) / (2 * z)
    j++
  }
  fmt.Println(j)
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
