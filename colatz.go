package main
import (
	"fmt"
)
func main(){
	collatzalgtorythm(9903)
}

func collatzalgtorythm(x int64){
//
	count := 1
	inputvalue := x
	
	for {	
		if x!=1 {
			if x%2 != 0 {
				z := 3*x+1
				fmt.Print(z, " ")
				count++
				x = z

			} else {
				z := x/2
				fmt.Print(z, " ")	
				count++
				x = z
				
			}
			
		} else {
			fmt.Println("")
			fmt.Println("For number", inputvalue, "was done:", count, "steps.")
			break
		}
	}
}
