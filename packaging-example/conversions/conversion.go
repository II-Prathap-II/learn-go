package conversions

type Celcius float64
type Farenheit float64

func FToC(f Farenheit) Celcius {
	c := (f - 32) * 5 / 9
	return Celcius(c)
}

func CToF(c Celcius) Farenheit {
	return Farenheit(c*9/5 + 32)
}
