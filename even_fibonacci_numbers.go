package main

import (
  "fmt"
)

func main(){
  var acc int = 0;
  var cn int;
  for n := 1; n <= 33; n++ {
    cn = Fibonacci(n);
    if(isEven(cn)){
      acc += cn;
    }
  }
  fmt.Println(acc);
}

func Fibonacci(n int) int {
  if( n < 2 ){
    return n;
  } else {
    return Fibonacci(n-1) + Fibonacci(n-2);
  }
}

func isEven(n int) bool{
  return n % 2 == 0;
}
