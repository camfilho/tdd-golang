package loop

func Repeat(char string) (result string) {
	for i := 0; i < 5; i++ {
		result = result + char
	}

	return
}