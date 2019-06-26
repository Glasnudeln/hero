package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  image := make([][]uint8, dy)
  for y := range image {
    image[y] = make([]uint8, dx)
  }
}

func main() {
  pic.Show(Pic)
}
