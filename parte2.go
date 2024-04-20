import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var valor string
		fmt.Scan(&valor)

		quantidadeLeds := calcularLeds(valor)
		fmt.Printf("%d leds", quantidadeLeds)
	}
}

func calcularLeds(numero string) int {
	leds := map[rune]int{
		'0': 6,
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
	}

	quantidadeLeds := 0
	for _, digito := range numero {
		quantidadeLeds += leds[digito]
	}

	return quantidadeLeds
}
