package main

import (
  "fmt"
  "strconv"
  "sort"
)

func main(){
  var palindromes []int = make([]int,0)
  for i := 100; i <= 999; i++ {
    for j := 100; j <= 999; j++ {
      if(isPalindrome(i * j)){
        palindromes = append(palindromes, i * j)
      }
    }
  }
  sort.Ints(palindromes)
  fmt.Println(palindromes[len(palindromes)-1])
}

func isPalindrome(n int) bool {
  var str string = strconv.Itoa(n)
  for i := 0; i < len(str)/2; i++ {
    if str[i] != str[len(str)-i-1] {
      return false
    }
  }
  return true
}
