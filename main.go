package main
import "fmt"
func main(){
	switch 5 {
	case 5,6,7:
		fmt.Println("the value is between 5 - 7")
		break
	case 8,9,10:
		fmt.Println("the value is between 8 - 10")
		break
	
	case 11,12,13:
		fmt.Println("the value is between 5 - 7")
		break
	default:
		fmt.Println("the value is not betweeen 5-13")
		break
	}
}