// Package tempconf performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }

func (c Celsius) String() string    { return fmt.Sprintf("%.1f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f°F", f) }
