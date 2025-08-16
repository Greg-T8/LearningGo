package tempconv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15

	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
	AbsoluteZeroF Fahrenheit = -459.67
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
