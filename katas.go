package main

import (
  "fmt"
)
import "math"

func Sqrt(x float64) float64 {
	var z float64=x/2;
	for i:=0; i<10;i++{
		z=z-((z*z)-x)/(2*z)
	}
	return z
}
func Sqrt2(x,param float64) float64 {
z := float64(x/2)
  var auxiliar float64 =0
  var contador int = 0
  for {
    contador++
    auxiliar=z
    z=z-(((z*z)-x)/(2*z))

    if auxiliar-z < param {
      fmt.Println(contador)
      return z
    }

  }
}
func main() {
  fmt.Println(Sqrt(8))
  fmt.Println(Sqrt2(8,1e-8))
  fmt.Println(math.Sqrt(8))
  fmt.Println("Hola nilton, digo mundo")
}