package reflection

func walk(x interface{}, fn func(input string)) {
	fn("Implement function inside walk")
}
