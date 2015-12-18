package conv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g feet",   f) }
func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
