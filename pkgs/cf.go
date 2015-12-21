package main

import (
  "fmt"

  "github.com/ghuecas/gosamples/pkgs/tempconv"
)

func main() {
   fmt.Printf("%gºF = %gºC\n", tempconv.CToF(tempconv.FreezingC), tempconv.FreezingC)
   fmt.Printf("%gºF = %gºC\n", tempconv.CToF(tempconv.BoilingC), tempconv.BoilingC)
}

