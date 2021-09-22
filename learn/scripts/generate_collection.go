package main 
		 
// how to test
// go run -gcflags=-G=3 learn/scripts/generate_collection.go
import "fmt" 
		 
func print[T any] (arr []T) { 
	for _, v := range arr { 
		fmt.Print(v) 
		fmt.Print(" ") 
	} 
	fmt.Println("") 
} 
		 
func main() { 
	strs := []string{"Hello", "World",  "Generics"} 
	decs := []float64{3.14, 1.14, 1.618, 2.718 } 
	nums := []int{2,4,6,8} 
 
	print(strs) 
	print(decs) 
	print(nums) 
}
		
		
