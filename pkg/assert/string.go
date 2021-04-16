package assert

func NotEmpty(value string, msg string) {
	if value == "" {
		panic(msg)
	}
}
