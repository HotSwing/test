package main 

import "strings"
import "fmt"
import "os"

func main() {
	fmt.Println("i am here")
	fmt.Println(strings.Join(os.Args[1:]," "))

}