package h7z

func urshift(number, bits int) int {
	return int(uint(number) >> uint(bits))
}
