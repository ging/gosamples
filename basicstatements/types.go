package main

import "fmt"

type Celsius float32     // looks the same
type Fahrenheit float32  // cannot be mixed, nor confused

const (
  AbsoluteZeroC Celsius    = -273.15
  FreezingC     Celsius    = 0.0
  BoilingF      Fahrenheit = 212.0
)

func main() {
   fmt.Printf("%gºF = %gºC\n", CToF(FreezingC), FreezingC)
   fmt.Printf("%gºF = %gºC\n", BoilingF, FToC(BoilingF))
}

func CToF(c Celsius ) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit ) Celsius { return Celsius((f - 32) * 5 / 9) }

