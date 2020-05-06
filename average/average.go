package average

// Average finds the average of the given numbers
func Average(numbers ...float64) float64 {
	numSlice := []float64(numbers)
	total := 0.0
	for _, num := range numSlice {
		total += num
	}
	return total / float64(len(numSlice))
}
