package main

import (
  "fmt"
)

func main(){
  var n int = 1
  for {
    a := triangolo(n)
    if n_divis(a) > 50 {
      fmt.Println(a)
      return
    }
    n++
  }
}

func triangolo(a int) int {
  var tri int
  for j := 1; j <= a; j++ {
      tri += j
  }
  return tri
}

func n_divis(n int) int {
  var count int
  for i := 1; i < n; i++ {
    if n % i == 0 {
      count++
    }
  }
  return count
}
