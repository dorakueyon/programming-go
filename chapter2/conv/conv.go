package conv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type KelvinScale float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

func (c Celsius) String() string     { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string  { return fmt.Sprintf("%gかし", f) }
func (k KelvinScale) String() string { return fmt.Sprintf("%gK", k) }
