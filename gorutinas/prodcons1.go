package main

// Producer - Consumer

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
    rand.Seed( time.Now().UnixNano() )

    start := time.Now()

    buffer := make (chan int)

    fmt.Printf("starting time=%s\n", start)

    go producer("P1", buffer, 800)
    go consumer("C1", buffer, 1000)

    time.Sleep(2000 * time.Millisecond)

    fmt.Printf("finishing time=%s\n", time.Now().String())
    fmt.Printf("%.2fs elapse\n", time.Since(start).Seconds())
}

func producer(name string, b chan<- int, delay time.Duration) {
  x := 1
  for {
    b <- x
    fmt.Printf("%s produced %d\n", name, x)
    x++
    time.Sleep(delay * time.Millisecond)
  }
}

func consumer(name string, b <-chan int, delay time.Duration) {
  var res int
  for {
    res = <- b
    fmt.Printf("\t%s consumed %d\n", name, res)

    time.Sleep(delay * time.Millisecond)
  }
}

