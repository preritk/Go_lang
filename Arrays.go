//Arrays in Go
//Arrays follows indexing like C/C++. Indexing starts from 0 .
package main
import "fmt"
func main(){
    var array_name [5] int //int is type of array elements
    //By default integer array is initialised by 0 in Go .
    fmt.Println(array_name) //Result is [0 0 0 0 0].
    array_name[4] = 100     //Changing the value of 5th element .
    fmt.Println(array_name) //Result is [0 0 0 0 100]
    //Below is another method of declaring and initialosing an array .
    arr := [5]int{1,2,3,4,5}
    fmt.Println(arr)

    //Now how to declare a 2D matrix in Go .
    var brr [3][3]int 
    for i :=0;i<3;i++{
        for j:=0;j<3;j++{            
            brr[i][j] = i*j
        }
    }
    fmt.Println(brr)  //No need to write a double loop to print 2D matrix in Go.
    
}
