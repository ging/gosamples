package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    histogram := make (map[int]int)
    for i := 0 ; i < 6; i++ {
        histogram[ rand.Intn(10) ]++
    }

    fmt.Println("HISTOGRAM IN ORDER");
    for i := 0; i < 10; i++ {
        fmt.Printf("Value=%d appeared %d times\n", i, histogram[i])
    }

    fmt.Println("HISTOGRAM NO ORDER");
    for value, amount := range histogram {
        fmt.Printf("Value=%d appeared %d times\n", value, amount)
    }

    fmt.Println("HISTOGRAM PERCENTAGE");
    sum := 0
    for _, amount := range histogram {
        sum += amount
    }
    for value, amount := range histogram {
        fmt.Printf("Value=%d appeared %.2f%% times\n", value, 100.0 * float32(amount)/ float32(sum))
    }

}

