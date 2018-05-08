//Constants in Go
//Declared by const keyword.
//Can come at places where a variable can come
package main
import "fmt"
const a string = "David"  //Global constants
func main(){
   const b = 2
   fmt.Println(float32(b)+2)    //Local constants
   const c = 2.5
   fmt.Println(c)
}
