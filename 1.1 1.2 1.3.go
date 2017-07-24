package main

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
	//输出参数的索引和值
	for i:= 1; i<len(os.Args); i++ {
		fmt.Println("第", i,"次索引为：", i-1,";第", i,"次索引的值：", os.Args[i], "\n")
	}
}