//@author Zsam
//Declaración de variables en Go
//github.com/starsaminf/golang-examples
package main

import (
	"fmt"
	"reflect"
)

func main() {

  var index int = 100 //En este caso la variable sera un int
  fmt.Println("Soy un", reflect.TypeOf(index)) // Soy un int

  var cadena = "Soy una cadena" //No es necesario poner el tipo de variable
                                //Go reconoce el tipo de variable
  fmt.Println("Soy un", reflect.TypeOf(cadena)) //Soy un string

  var numero = 10.1
  fmt.Println("Soy un", reflect.TypeOf(numero)) //Soy un float64

  //Otra forma de declarar variables
  index2 := 200 //modo corto
  fmt.Println("Soy un", reflect.TypeOf(index2)) //Soy un int

  cadena2 := "Soy una cadena"
  fmt.Println("Soy un", reflect.TypeOf(cadena2)) //Soy un string

}

