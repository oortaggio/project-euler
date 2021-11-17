package main

import (
  "fmt"
)

func main(){
  var results map[int]int
  results = make(map[int]int)
  for n := 1; n <= 1000000; n++ {
    results[n]= collatz(n)
  }
  var max = []int{0,0}
  var temp int
  for i, v := range results {
    temp = v
    if temp > max[1] {
      max[1] = temp
      max[0] = i
    }
  }
  fmt.Println(max[0], max[1])
}

func collatz(a int) int {
  n:= a
  chain_count:= 0
  for n > 1 {
    if n % 2 == 0 {
      n = n / 2
    } else {
      n = 3 * n + 1
    }
    chain_count++
  }
  return chain_count
}
