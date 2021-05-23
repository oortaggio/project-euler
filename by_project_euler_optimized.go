package main

import (
  "fmt"
  "math"
)

// not my solution, optimized after https://projecteuler.net/overview=007 check
// omitted bruteforce solution
// prime number theorem probabilistic solution abandoned for now
func main(){
  limit := 10001
  count := 1
  candidate := 1
  for count != limit {
    candidate = candidate + 2
    if isPrime(candidate) {
      count += 1
    }
  }
  fmt.Println(candidate)
}

func isPrime(n int) bool {
  if n==1 {
    return false
  } else if n < 4 {
    return true
  } else if n % 2 == 0 {
    return false
  } else if n < 9 {
    return true
  } else if n % 3 == 0 {
    return false
  } else {
    r:= math.Floor(math.Sqrt(float64(n)));
    f:= 5
    for f<=int(r) {
      if n % f == 0 {
        return false
      }
      if n % (f+2) == 0 {
        return false
      }
      f = f+6
    }
    return true
  }
}
