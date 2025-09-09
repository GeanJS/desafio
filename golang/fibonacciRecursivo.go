package main

func fibonacciRecursivo(numero int) int {
	if numero < 0 {
		return -1
	}

	if numero <= 1 {
		return numero
	}
	return fibonacciRecursivo(numero-1) + fibonacciRecursivo(numero-2)
}
