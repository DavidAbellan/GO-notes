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
//%d indica que es decimal
//importando el paquete strconv también se podría, tiene muchas opciones

//uint --- sin signos + ó -

package main

//se importa el mismo package

import "fmt"

//import "fmt"

var numero int
var texto string
var bul bool

//func main(){----- la primera linea con { o si no no compila
func main() {

	numero3 := 4
	fmt.Println(numero3)
	estados()
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
