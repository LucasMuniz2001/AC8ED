package main
 
import (
    "fmt"
)
 
func main() {
 var d, n int

    for {
        fmt.Scan(&d, &n)
        
        if d == 0 && n == 0 {
            break
        }

        fmt.Println(calculaValor(d, n))
    }

}

func calculaValor(d, n int) int {
    var novoNumero int

    for n > 0 {
        digito := n % 10

        if digito != d {
            novoNumero = novoNumero*10 + digito
        }

        n /= 10
    }

    return novoNumero
}
