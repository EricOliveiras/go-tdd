package iteration

func Repeat(caracter string, amountRepeats int) string {
	var repeat string
	
	for i := 0; i < amountRepeats; i++ {
		repeat += caracter
	}

	return repeat
}