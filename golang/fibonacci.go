package main

func fibonacci(numero int) int {
	if numero < 0 {
		return -1
	}
	if numero <= 1 {
		return numero
	}
	
	a, b := 0, 1
	i := 0
	for i < numero {
		a, b = b, a + b
		i++
	}
	return a
}

