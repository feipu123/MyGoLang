package main
import "fmt"
func main() {
     fmt.Println("Hello, world.")
	 var a [10]int
	 a = [10]int{1,2,3,4,5}
	 a[0] = 1;
	 for i, v := range a {
		fmt.Println(i, v)
	 }
}