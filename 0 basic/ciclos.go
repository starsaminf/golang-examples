//@author Zsam | starsaminf
//Ciclos repetitivos
//En golang un while = for
//github.com/starsaminf/golang-examples
package main/

import "fmt"

func main() {

 fmt.Println("Distintas formas de usar un for")

 i:= 'a'
 for {
   if i > 'd' {
    break
   }
  fmt.Printf("%c ", i) //%c = Character
  i++
 }

 fmt.Println()
 limit:= 0
 for limit  < 10 {
  fmt.Printf("%d ", limit) //%d = int
  limit++
 }

 fmt.Println()
 //Jugando con el printf
 for i:= 0; i < 20; i++ {
   fmt.Printf("Decimal = %d, octal = %o , hexadecimal = %X \n", i, i, i) //%d base 10, %o octal y %X o %x hexadecimal
 }
 fmt.Println()
}
