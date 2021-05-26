//go run main.go ---- ejecutar sin compilar
//go build main.go ---- compila no ejecuta

//si se declaran de esta forma
//numero3 := 4
//para asignar nuevos valores
//numero3= 15

//variables locales
//var numero int
//var numero int = 7
//func numeros (){

//Variables y funciones visibles de otros paquetes
//var Numero int
//func Numeros (){

//var texto string
//var bul bool

//var numero1, numero2, numero4 int
//numero5 , numero6 , numero7 := 5
//numero5 , numero6 , numero7 := 5,6,7
//numero5 , numero6 , numero7, text := 5,6,7,"aqui un string"

//int8
//int32
//float32

//se puede transformar entre int
//var numero int32
//numero2 = int(numero)
//int también es una función

//en otros casos es más complicado...
//convertir decimal a string
//texto = fmt.Sprintf('%d', moneda)
//%d indica que es base 10
//importando el paquete strconv también se podría, tiene muchas opciones

//uint --- sin signos + ó -

package main

//se importa el mismo package

import (
	"bufio"
	"fmt"
	"os"
)

//import "fmt"

var numero int
var texto string
var bul bool

//func main(){----- la primera linea con { o si no no compila
func main() {

	numero3 := 4
	fmt.Println(numero3)
	estados()
	pedirPorTeclado()
	iteraciones()
}

////Condicionales

func estados() {
	estado := true
	//if ---- { , }else{, debe de estar en la misma línea
	//si no no compila
	// se puede asignar una variable en el mismo if
	//if estado = false; estado==true
	//if <asignacion>   ;  <comparacion>

	//switch numero:=5 ; numero {
	//asignacion ---- comparacion
	//}
	if estado {
		fmt.Println(estado)
	} else {
		fmt.Println("Es Falso")
	}
}
func pedirPorTeclado() {
	var numero1 int
	var numero2 int
	var leyenda string
	var resultado int
	//fmt.Println("Ingrese numero :")
	//fmt.Scanf("%d", numero1)

	//fmt.Println("Ingrese numero :")
	//fmt.Scanf("%d", numero2)
	// de esta forma no funciona en Windows
	// porque toma enter como valor del numero2

	fmt.Println("Ingrese numero :")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero :")
	fmt.Scanln(&numero2)

	//fmt.Printf("%d", numero1+numero2)

	fmt.Println("Descripción :")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}

//BUCLES
//Sólo existe FOR

func iteraciones() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for y := 1; y < 10; y++ {
		fmt.Println(y)
	}

	numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scan(&numero)
		if numero == 100 {
			break
		}

	}
	///intercalar la variable
	var z = 0
	for z < 10 {
		fmt.Printf("\n Valor de z : %d", z)
		if z == 5 {
			fmt.Printf("multiplicamos por 2 \n")
			z = z * 2

			//continue devuelve a la linea primera del for	for z < 10 {
			continue
		}
		fmt.Printf("    Pasa por aquí\n")
		z++
	}

	var f int

	// crear una sección, no afecta a la ejecución
RUTINA:

	for f < 10 {
		if f == 4 {
			f = f + 2
			fmt.Println("Voy a RUTINA sumando 2 a f")
			goto RUTINA
		}
		fmt.Printf("valor de i: %d \n", f)
		f++
	}

}
