package go_arithmetic

func Multiply(items ...interface{}) int64 {
	res := int64(1)
	for _, item := range items {
		intVal := item.(int)
		res *= int64(intVal)
	}
	return res
}

func Add(items ...interface{}) int64 {
	res := int64(0)
	for _, item := range items {
		intVal := item.(int)
		res += int64(intVal)
	}
	return res
}

func Subtract(items ...interface{}) int64 {
	res := int64(items[0].(int)) * 2

	for _, item := range items {
		intVal := item.(int)
		res -= int64(intVal)
	}
	return res
}
