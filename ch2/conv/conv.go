package conv

func FToM(f Feet) Meters { return Meters(f / 3.2808) }

func MToF(m Meters) Feet { return Feet(m * 3.2808) }

func PToK(p Pounds) Kilograms { return Kilograms(p / 2.2046) }

func KToP(k Kilograms) Pounds { return Pounds(k * 2.2046) }
