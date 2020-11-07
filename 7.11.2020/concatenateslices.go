package main 
import "fmt"
  
func main() {
    a := []int{1, 2, 3, 4, 5}
    b := []int{6, 7, 8, 9, 10}
    c := []int{11, 12, 13, 14, 15}
     
    fmt.Printf("a: %v\n", a)
    fmt.Printf("cap(a): %v\n", cap(a))
     
    fmt.Printf("b: %v\n", b)
    fmt.Printf("cap(c): %v\n", cap(b))
     
    fmt.Printf("c: %v\n", c)
    fmt.Printf("cap(c): %v\n", cap(c))
     
    x := []int{}
    x = append(a,b...)  
    x = append(x,c...)  
     
    fmt.Printf("\n######### After Concatenation #########\n")
    fmt.Printf("x: %v\n", x)
    fmt.Printf("cap(x): %v\n", cap(x))
}