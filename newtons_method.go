package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 10.0
  for delta := 10.0; delta > 1e-100; {
    next := z - (((z * z) - x) / (2 * z))
    delta = z - next
    z = next
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
