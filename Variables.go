//Variables in Go
package main
import "fmt"
func main(){
   var a string = "David"
   fmt.Println(a)
   //The same thing can be done using := operator in Go . := declares as well as initialize the variable.
   b := "David"
   fmt.Println(b)
   //You can declare multiple variables at a time . GO will infer them all.
   var c,d int = 2,3
   fmt.Println(c,d)
   var e = 2
   //If you don't declare the type of variable Go is intelligent enough to understand the type .
   fmt.Println(e+2)   //Result is 4
}
