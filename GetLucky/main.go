package main

import (
	"strconv"
	"strings"
)

func getLucky(s string, k int) int {
	// Paso 1: Convertir cada letra en su posición en el alfabeto
	var numStr strings.Builder
	for _, char := range s {
		numStr.WriteString(strconv.Itoa(int(char - 'a' + 1)))
	}

	// Convertir el número construido a entero para la siguiente fase
	num, _ := strconv.Atoi(numStr.String())

	// Paso 2: Transformar la suma de los dígitos k veces
	for i := 0; i < k; i++ {
		num = sumDigits(num)
	}

	return num
}

// Función auxiliar para sumar los dígitos de un número
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}