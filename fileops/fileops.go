package fileops

import "math"

func CalculateFutureValues(investmentAmount float64, expectedReturnRate float64, years int) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	rfv = fv / math.Pow(1+2.5/100, float64(years)) // Assuming 2% inflation rate
	return fv, rfv
}
