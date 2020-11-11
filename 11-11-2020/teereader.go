package main
 
import (
    "bytes"
    "io"
    "fmt"
    "strings"
)
 
func main() {
    testString := strings.NewReader("Jobs, Code, Videos and News for Go hackers.")
    var bufferRead bytes.Buffer
    example := io.TeeReader(testString, &bufferRead)
    readerMap := make([]byte, testString.Len())
    length, err := example.Read(readerMap)
    fmt.Printf("\nBufferRead: %s", &bufferRead)
    fmt.Printf("\nRead: %s", readerMap)
    fmt.Printf("\nLength: %d, Error:%v", length, err)
}