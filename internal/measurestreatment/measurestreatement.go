package measuretreatement

import data "Moniport/internal/data"

func getAverageFromMeasures(measures []data.Measure) float64 {
	sum := .0

	for _, measure := range measures {
		sum += measure.Value
	}

	return sum / float64(len(measures))
}
