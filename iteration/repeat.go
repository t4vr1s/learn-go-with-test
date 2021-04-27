package iteration

func Repeat(char string, cant int) string {
	var repeat string
	for i := 0; i < cant; i++ {
		repeat += char
	}
	return repeat
}
