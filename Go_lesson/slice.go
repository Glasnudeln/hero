package main

import "fmt"

func main() {
  s := []int{2, 4, 5, 7, 33, 13}
  q := []int{2, 3, 5, 7, 11, 13}

  r := []bool{true, false, true, true, false, false}
  fmt.Println(r)
  fmt.Println(q)

  str := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, false},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(str)
  fmt.Println(str[0])
  printSlice(q)
//  printSliceS(str.i)

  s = s[1:4]
  fmt.Println(s)

  s = s[:2]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d%v\n", len(s), cap(s), s)
}

func printSliceS(s []int) {
  fmt.Printf("len=%d cap=%d%v\n", len(s), cap(s), s)
}
