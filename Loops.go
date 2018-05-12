//Loops in Go
//There are only for and while loop in go language .
//A third type of loop is foreach loop .
//No dowhile loop in Go .
package main
import "fmt"
func main(){
fmt.Println("For loop running::::")
    for i := 1 ; i<=10 ; i++{
        fmt.Println(i)
    }
fmt.Println("While Loop running::::")
    j := 0
    for j<=5{
        fmt.Println(j)
        j++
    }
fmt.Println("Foreach loop running::::")
//foreach loop is used to iterate over arrays,strings,slices etc.
//It uses range keyword .
//If used for associative data structure ,the iterator holds the key - value pair .
//for _, value returns only values.
//for _, key returns only keys .
    studentGrades := [3]int{10,15,20}
    for _, grades := range studentGrades {
        fmt.Println("Grade is: ",grades)
    }
}
