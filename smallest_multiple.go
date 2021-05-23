package main

import (
  "fmt"
)

func main(){
    for i := 1; i > -1 ; i++ {
      if isDivisible20(i) {
        fmt.Println(i)
        break
      }
    }
}

func isDivisible20(a int) bool {
  for k := 1; k <= 20 ; k++ {
    if a % k != 0 {
      return false
    }
  }
  return true
}
