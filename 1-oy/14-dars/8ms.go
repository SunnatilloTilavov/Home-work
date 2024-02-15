func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	maxnum := candles[0]
	for _, num := range candles {
		if num > maxnum {
			maxnum = num
		}
	}
	s := 0
	for i := 0; i < len(candles); i++ {
		if candles[i] == maxnum {
			s++
		}
	}
	return int32(s)
}