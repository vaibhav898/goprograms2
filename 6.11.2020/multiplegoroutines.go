package main
 
import (
    "fmt"
    "runtime"
    "sync"
)
 
func main() {
   
    runtime.GOMAXPROCS(3)
 
   
    var processTest sync.WaitGroup
   
    processTest.Add(3)
     
   
    go func() {
        defer processTest.Done()
        for i := 0; i < 30; i++ {
            for j := 51; j <= 100; j++ {
                fmt.Printf(" %d", j)
                if j == 100{
                    fmt.Println()
                }
            }
        }
    }()
    go func() {
        defer processTest.Done()
        for j := 0; j < 10; j++ {
            for char := 'A'; char < 'A'+26; char++ {
                fmt.Printf("%c ", char)
                if char == 'Z' {
                    fmt.Println()
                }
 
            }
        }
    }()
    go func() {
        defer processTest.Done()
        for i := 0; i < 30; i++ {
            for j := 0; j <= 50; j++ {
                fmt.Printf(" %d", j)
                if j == 50 {
                    fmt.Println()
                }
            }
        }
    }()
 
    processTest.Wait()  
}