package iteration

const repeatCount = 5

func Repeat(char string) string {

	var repeat string
	for i := 0; i < repeatCount; i++ {

		repeat += char
	}

	return repeat
}
