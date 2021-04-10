package operation

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	if y != 0 {
		return x / y
	} else {
		return 0
	}
}
