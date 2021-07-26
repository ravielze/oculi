package functions

import "math"

func CalculateTotalPages(totalCount int, limit int) int {
	tp := float64(totalCount) / float64(limit)
	tp = math.Ceil(tp)
	return int(tp)
}
