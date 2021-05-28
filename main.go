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

	//excepciones
	"log"
	"os"
	"time"

	//GO asume que dentro de la carpeta user hay un archivo user.go
	// solo se pueden importar carpetas
	//si no se llamara user.go no funcionaría
	usuario "./user"
	//Manejo de archivos
	"io/ioutil"
	//goroutines
	"strings"
	//redes
	"net/http"
)

// go env -w GO111MODULE=auto
//en línea de comando si no coge los paquetes que creas(user)

//import "fmt"

var numero int
var texto string
var bul bool

//Goroutines

//func main(){----- la primera linea con { o si no no compila
func main() {
	fmiddlewr()
	//detectarLocalhost()
	numero3 := 4
	fmt.Println(numero3)
	go miNombreLento("David Abellán Motos")
	estados()
	pedirPorTeclado()
	iteraciones()
	fmt.Println(uno(5))
	n, e := dos(1)
	fmt.Println(n)
	fmt.Println(e)
	tres(1)
	fmt.Println(cuatro(5, 34, 22, 12, 12, 22, 1))
	fmt.Println(cuatro(5, 34, 1))

	fmt.Printf("Sumo 5 + 7  = %d \n", calculo(5, 7))

	//calculo es una variable de tipo función y se puede redefinir(funciones anónimas)
	// cuando se redefine hay que respetar parámetros de entrada y salida
	calculo = func(i1, i2 int) int {
		return i1 - i2
	}

	fmt.Printf("Resto 51 - 7  = %d \n", calculo(51, 7))
	Operaciones()

	//Closure
	fmt.Println("Closure")

	tablaDel2 := 2
	MiTabla := tabla(tablaDel2)
	for po := 1; po < 11; po++ {
		fmt.Println(MiTabla())
	}
	//la instrucción go hace ejecutar miNombreLento de forma asíncrona
	fmt.Println("estoy aquí")
	var x string
	fmt.Scanln(&x)
	canales()
	//ejemploPanic()
	mostrarSlice()
	variante2()
	mapas()
	estructuras()
	interfeises()
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
	exponencia(2)
	//viendoDefer()

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

//Funciones

func uno(numero int) int {
	return numero * 2
}

//función que devuelve dos parámetros, se pueden devolver los parámetros
//cuantos se quieran
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 11, true
	}
}

//con alias para la salida
func tres(numero int) (z int) {
	z = numero * 2
	return
}

//con parámetros variables

func cuatro(numero ...int) int {
	total := 0
	// _ se pone cuando recibimos una variable que no vamos a usar
	//en este caso devuelve el index n de elementos que se mandan
	for _, num := range numero {

		total = total + num

	}
	return total
}

///funciones anónimas y CLOSURE

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	println("resultado : ", resultado())
}

///CLOSURE
//función tabla que devuelve una función anónima que devuelve un int
func tabla(valor int) func() int {

	//la primera vez que se ejecuta tabla crea las variables
	//numero y secuencia, las siguientes ejecuta solo la función
	//que hay dentro, o sea
	numero := valor
	secuencia := 0

	// a partir de aquí
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

//SLICES

func mostrarSlice() {
	var tabla2 [10]int

	tabla2[0] = 1
	tabla2[5] = 15

	fmt.Println(tabla2)

	tabla3 := [10]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 19}

	fmt.Println(tabla3)

	for i := 0; i < len(tabla3); i++ {
		fmt.Println(tabla3[i])

	}

	var matriz [5][7]int
	matriz[3][5] = 1
	fmt.Println(matriz)

	///SLICES === array y matrices dinámicos

	matriz2 := []int{1, 2, 4}
	fmt.Println(matriz2)

}
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//crea un array llamado porcion con valores desde
	// la posicion 3 del array elementos hasta la última
	porcion := elementos[3:]
	fmt.Println(porcion)
	//de la posicion 1 a la 4
	porcion = elementos[1:4]
	fmt.Println(porcion)

	//make crea un array de 5 elementos int pero
	// con capacidad en memoria de hasta 20
	// quiere decir que reserva hasta 20 no que solo se puedan meter 20
	elementos2 := make([]int, 5, 20)

	//printf para concatenar con %d
	fmt.Printf("largo %d, Capacidad %d", len(elementos2), cap(elementos2))

	nums := make([]int, 0, 0)
	for j := 0; j < 100; j++ {
		nums = append(nums, j)
	}
	fmt.Printf("largo %d, Capacidad %d", len(nums), cap(nums))

}

////MAPS

///como array pero con clave : valor

func mapas() {
	paises := make(map[string]string)
	paises = make(map[string]string, 5)

	paises["Méjico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   39,
		"Chivas":      23,
		"Real Madrid": 12,
		"Juventus":    8,
		"Leipzig":     2,
	}
	fmt.Println(campeonato)
	//lo muestra alfabéticamente por clave

	//añadir
	campeonato["River Plate"] = 22
	fmt.Println(campeonato)
	campeonato["Chivas"] = 44
	campeonato["Alavés"] = 43
	fmt.Println(campeonato)
	delete(campeonato, "River Plate")
	fmt.Println(campeonato)

	//el range ordena aleatoriamente
	for equipo, puntaje := range campeonato {
		fmt.Println(equipo)
		fmt.Println(puntaje)

	}
	puntaje, existe := campeonato["Atletico"]

	//%t imprime booleano
	fmt.Printf("El puntaje %d y el equipo existe %t", puntaje, existe)

	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje %d y el equipo existe %t", puntaje, existe)

}

///Estructuras (No existen las clases ni métodos)
func estructuras() {

	/*	type usuario struct {
			Id        int
			Nombre    string
			fechaAlta time.Time
			Status    bool
		}

		User := new(usuario)
		User.Id = 10
		User.Nombre = "Diego"
		fmt.Println(User)*/

	///Herencia

	type pepe struct {
		usuario.Usuario
	}

	user := new(pepe)
	user.AltaUsuario(1, "German Burgos", time.Now(), true)
	fmt.Println(user.Usuario)

}

//Interfaces

//herencia de funciones, comportamientos

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estarVivo() bool
}

type animal interface {
	respirar()
	estarVivo() bool

	esCarnivoro() bool
}
type vegetal interface {
	clasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	//variante heredando mujer de hombre
	esHombre bool

	//superClase EstarVivo
	vivo bool
}

/*type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}*/

//mujer hereda todos los campos de la estructura hombre
type mujer struct {
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	//return "Hombre"
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

/*func (m *mujer) respirar() {
	m.respirando = true
}
func (m *mujer) comer() {
	m.comiendo = true
}
func (m *mujer) pensar() {
	m.pensando = true
}
func (m *mujer) sexo() string {
	return "Mujer"
}*/

func humanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (h *perro) respirar() {
	h.respirando = true
}
func (h *perro) comer() {
	h.comiendo = true
}

func (p *perro) esCarnivoro() bool {
	return p.carnivoro
}
func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}
func AnimalesCarnivoros(an animal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}
func interfeises() {
	Pedro := new(hombre)
	//segunda variante
	Pedro.esHombre = true
	humanosRespirando(Pedro)

	Maria := new(mujer)
	humanosRespirando(Maria)

	totalCarnivoros := 0
	dogo := new(perro)
	dogo.carnivoro = true
	AnimalesRespirar(dogo)
	totalCarnivoros = +AnimalesCarnivoros(dogo)
	fmt.Printf("Total Carnivoros %d", totalCarnivoros)

	//Si tiene las funciones declaradas en la interfaz
	//func (h *perro) respirar() {
	//h.respirando = true
	//      }

	//func (p *perro) esCarnivoro() bool {
	//return p.carnivoro
	//}

	//automáticamente se implementa y se relaciona con ese tipo
	//structura perro---interfaz animal

	//permite que pertenezca a varias interfaces

}

type serVivo interface {
	estarVivo() bool
}

func (h *hombre) estarVivo() bool {
	return h.vivo
}
func (h *perro) estarVivo() bool {
	return h.vivo
}

////Manejo de archivos

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./text.txt")
	//null en GO es nil
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))

	}

}
func leoArchivo2() {
	archivo, err := os.Open("./text.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea >" + registro + "\n")

		}

	}
	//con os hay que cerrar el archivo
	archivo.Close()
}
func graboArchivo() {
	archivo, err := os.Create("./newtext.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Escribiendo en el archivo")
	//con os hay que cerrar el archivo
	archivo.Close()
}
func graboArchivo2() {
	FileName := "./newtext.txt"
	if agregar(FileName, "\n contenido de la función graboArchivo2") == false {
		fmt.Println("Error en la segunda Función")
	}

}
func agregar(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	//el guión bajo es para no recibir un parámetro de los que devuelve
	//la función , en GO pueden devolver dos parámetros
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true

}

///Recursividad

func exponencia(numero int) {
	if numero > 10000000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}

//Defer / Panic recover
//manejo de errores

//no hay try catch

func viendoDefer() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	//defer(va a ejecutarse siempre) va a cerrar el archivo independientemente se abra o no
	//para liberar buffer en memoria por el open
	defer f.Close()
	if err != nil {
		fmt.Println("error abriendo archivo")
		os.Exit(1)
	}
}
func ejemploPanic() {
	//para recuperarse del error, no se puede ejecutar nada
	//después de panic
	//ponemos defer para asegurarnos que podemos seguir ejecutando
	//defer sólo ejecuta una cosaç
	// recover es una función que recoge si hubo un panic
	defer func() {
		rcover := recover()
		if rcover != nil {
			//%v para poder enviar un valor variable reco es un objeto
			log.Fatalf("Ocurrió un error(panic) %v", rcover)
		}
	}()
	am := 1
	if am == 1 {
		panic("se encontró valor de 1")
	}
}

///GORoutines

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}

//channels
//cuando acaba el programa se cierra
//aunque esté ejecutándose una operación
//asíncrona de fondo termina

//channels reserva un espacio de memoria

func canales() {
	canal1 := make(chan time.Duration)

	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")

	///estoy esperando así que termine esta función
	//se detiene la ejecucíon

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	canal <- final.Sub(inicio)
}

//servidor web

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
func detectarLocalhost() {
	http.HandleFunc("/", home)
	//escuchando puerto
	http.ListenAndServe(":3000", nil)

}

//Middleware

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
func fmiddlewr() {
	fmt.Println("inicio")

	result := operacionesMid(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(multiplicar)(2, 3)
	fmt.Println(result)
}

//lo que recibe el middleware y lo que devuelve debe ser igual
func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de operación")
		return f(a, b)
	}

}
