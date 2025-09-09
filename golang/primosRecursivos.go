package main

import "math"

func primosRecursivos(numero int) []int  {
	if numero < 2 {
		return nil
	}

	primos := primosRecursivos(numero - 1)
	ehPrimo := true
	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero % i == 0 {
			ehPrimo = false
			break
		}
	}
	if ehPrimo {
		primos = append(primos, numero)
	}

	return primos
}

