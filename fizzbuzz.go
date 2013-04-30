package main

import "fmt"
import "math"

var x = 1.00

func main() {
  for i := x; i < 101; i++ {
  if math.Mod(i, 3)==0 && math.Mod(i, 5)==0{
    fmt.Println("FizzBuzz")
    } else if math.Mod(i, 3)==0{
      fmt.Println("Fizz")
    } else if math.Mod(i, 5)==0{
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}
