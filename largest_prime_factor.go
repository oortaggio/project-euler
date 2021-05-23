package main

import(
  "fmt"
)

func main(){
  const nmb int = 600851475143;
  var largest_factor []int = make([]int,0);
  var num_cr int = nmb;

  for k := 2; k < num_cr; k++ {
    var primo = nextPrime(k);
    if num_cr % primo == 0 {
      largest_factor = append(largest_factor, primo);
      fmt.Println("Numero: ",num_cr," / ",primo )
      num_cr = num_cr / primo;
    }
  }
  fmt.Println(largest_factor);
  fmt.Println(largest_factor[len(largest_factor)-1]);
}

func nextPrime(numPrime int) int {
  numPrime++
  for isPrime(numPrime) == false {
    numPrime++;
  }
  return numPrime;
}

func isPrime(n int) bool {
  for i := 2; i < n-1 ; i++ {
    if n % i == 0 {
      return false;
    }
  }
  return true;
}
