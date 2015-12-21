// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float32
type Fahrenheit float32

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0.0
    BoilingC      Celsius = 100.0
)

func (c Celsius) String() string { return fmt.Sprintf("%gªC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gªF", f) }

