package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    heads := 0
    tails := 0
    for i := 0 ; i < 10; i++ {
        switch throwCoin() {
  	case "heads" : heads++
  	case "tails" : tails++
  	default : fmt.Println("landed on edge!")
	}
    }

    fmt.Printf("Heads=%d, Tails=%d\n", heads, tails)
}

func throwCoin() string {
    switch rand.Intn(2) {
    case 0 : return "heads"
    case 1 : return "tails"
    default : return "edge"
    }
}

