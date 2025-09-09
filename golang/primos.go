package main

import (
	"math"
)

func primosLinear(numero int) []int {
	if numero < 2 {
		return nil
	}

	SliceDePrimos := []int{}
	
	for i := 2; i <= numero; i++ {
		ehPrimo := true
		limite := int(math.Sqrt(float64(i)))
		
		for j := 2; j <= limite; j++ {
			if i%j == 0 {
				ehPrimo = false
				break
			}
		}
		if ehPrimo {
			SliceDePrimos = append(SliceDePrimos, i)
		}
	}
	return SliceDePrimos
	
}




