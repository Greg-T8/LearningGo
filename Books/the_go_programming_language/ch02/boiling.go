// Boiling prints the boiling point of water.

package main					// package-level declaration

import "fmt"

const boilingF = 212.0			// package-level declaration

func main() {
	f := boilingF				// local declaration
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g degrees Fahrenheit or %g degrees Celsius\n", f, c)
}
