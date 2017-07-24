package main 
import (
	"bufio"
	"fmt"
	"os"
) 

func main() {
	for i := 1; i < len(os.Args); i++ {
		    arg := os.Args[i]
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n",  err)
				continue
			}
			input := bufio.NewScanner(f)
			for input.Scan() {
				counts[input.Text()]++
			}
			f.Close()
			for line, n := range counts {
					if n>1 {
						fmt.Printf("%s\t%s", arg , line)
					}
			}
	}
}