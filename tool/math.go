package tool

func Two2TheNthPower(number int) int {
	if (number%2 != 0 && number != 1) || number <= 0 {
		return -1
	}
	var n = 0
	for number != 1 {
		number = number / 2
		n++
	}
	return n
}
