//conditional statements are same as C/C++.
//No paranthesis needed here.
//Braces are required .
package main

import "fmt"

func main() {
   if 7%2 ==0 {
       fmt.Println("7 is even.")
   }else{
       fmt.Println("7 is odd")
   }
   
   num := 9
   if num > 0 {
       fmt.Println("Number is non negative")
   }else{
       fmt.Println("Number is zero or negative")
   }
}
