package iteration

// const repeatCount = 5

func Repeat(str string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return repeated
}
