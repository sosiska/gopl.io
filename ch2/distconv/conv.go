package distconv

// CToF converts a Celsius temperature to Fahrenheit.
func FToM(f Foots) Meters { return Meters(f * 0.305) }

// FToC converts a Fahrenheit temperature to Celsius.
func MToF(m Meters) Foots { return Foots(m / 0.305) }
