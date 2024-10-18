package unitconv

import "fmt"

type Meters float64
type Kilograms float64
type Feets float64
type Pounds float64

// FtoM converts Feets temperature to Meters.
func FtoM(f Feets) Meters { return Meters(f * 0.304) }

// MtoF converts Meters into Feet
func MtoF(m Meters) Feets { return Feets(m * 3.28084) }

func KtoP(k Kilograms) Pounds { return Pounds(k * 2.2) }

func PtoK(p Pounds) Kilograms { return Kilograms(p * 0.45359237) }

func (m Meters) String() string    { return fmt.Sprintf("%gm", m) }
func (f Feets) String() string     { return fmt.Sprintf("%gft", f) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%glb", p) }
