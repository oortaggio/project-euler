package main

import (
  "fmt"
  "math"
)

func main(){
  fmt.Println(squareOfSum(100)-sumOfSquares(100))
}

func sumOfSquares(n int) int {
  var sum_of_sq int
  for k := 1; k <= n; k++ {
    sum_of_sq += pow(k,2)
  }
  return sum_of_sq
}

func squareOfSum(n int) int {
  var square_of_sum int
  for j := 1; j <= n; j++ {
    square_of_sum += j
  }
  square_of_sum = pow(square_of_sum,2)
  return square_of_sum
}

func pow(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}
