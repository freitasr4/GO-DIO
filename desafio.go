// declarar meu pacote principal
package main

// importar uma função fmt
import "fmt"

// declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoK float64 = 373.15

// função principal
func main() {

	tempK := ebulicaoK   // variável do valor da temperatura em graus k
	tempC := tempK - 273 //Conversão de K para C

	fmt.Printf("A temperatura de ebulição da água em K = %g e temperatura de ebulição da água em C = %g ", tempK, tempC)

}
