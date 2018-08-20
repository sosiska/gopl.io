// Package distconv performs Meters and Foots conversions.
package distconv

import "fmt"

type Foots float64
type Meters float64

func (f Foots) String() string  { return fmt.Sprintf("%g foots", f) }
func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
