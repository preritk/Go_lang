package main
//If you want to import multiple modules use the syntax as used below .
import (
    "fmt"
    "time"
    )
    
//switch case is similar to C/C++
//No need of using break statement here between cases.Go is smart enough to do so.
func main(){
    i := 5
    switch i{
        case 2:
            fmt.Println("Number is 2")
        case 5:
            fmt.Println("Number is 5")
        default:
            fmt.Println("Number is neither 2 or 5.")
    }
//You can also use functions in switch condition .
//You can put multiple cases together as illustrated below .
    switch time.Now().Weekday(){
        case time.Saturday,time.Sunday :
            fmt.Println("Today is a weekend.")
        default :
            fmt.Println("Go to your work.")
    }
    
}
