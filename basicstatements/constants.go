package main

import "fmt"

const BoilingF = 212.0
const FreezingC = 0.0

func main() {
   fmt.Printf("%gºF = %gºC\n", CToF(FreezingC), FreezingC)
   fmt.Printf("%gºF = %gºC\n", BoilingF, FToC(BoilingF))
}

func CToF(c float32 ) float32 { return c*9/5 + 32 }
func FToC(f float32 ) float32 { return (f-32)*5 / 9 }

