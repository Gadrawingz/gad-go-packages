package simpleInterest

/*
Let's calculate and return the simple interest for:
- a principal p,
- rate of interest r for
- time duration t years
*/

func CalculateInterest(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
