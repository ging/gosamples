package main

import (
  "fmt"
  "time"
)

func main() {
    go clock ()

    const n= 45
    fibN := fib(n) // quite slow

    fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}

func clock() {
  for {
    t := time.Now()
    fmt.Printf("\r%d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
    time.Sleep(1000 * time.Millisecond)
  }
}

func fib(x int) int {
  if x < 2 {
    return x
  }

  return fib(x-1) + fib(x-2)
}


