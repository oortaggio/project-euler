package main

import (
  "fmt"
)

func main(){
  var acc int = 0;
  for n := 1; n < 1000; n++ {
    if (n % 3 == 0 || n % 5 == 0){
      acc += n;
    }
    fmt.Println(n);
  }
  fmt.Println(acc);
}
