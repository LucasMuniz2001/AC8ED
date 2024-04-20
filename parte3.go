package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var codificado string
		var desloca int
		fmt.Scan(&codificado, &desloca)

		decodificado := decodificar(codificado, desloca)
		fmt.Println(decodificado)
	}
}

func decodificar(codificado string, desloca int) string {
	var decodificado string
	for _, char := range codificado {
		// Convertendo o caractere para sua posição relativa no alfabeto (0-25)
		posicaoCaractere := int(char - 'A')

		// Decodificando o caractere
		caractereDecodificado := 'A' + (posicaoCaractere-desloca+26)%26

		decodificado += string(caractereDecodificado)
	}
	return decodificado
}
