package homework

func average(input [15]float32) (result float32) {
	var sums float32 = 0
	for _, v := range input {
		sums += v
	}
	return sums / float32(len(input))
}
