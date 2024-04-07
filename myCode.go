package main

import (
	"fmt"
	"os"
)

func main(){
	f, err  := os.OpenFile("file.txt",os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 1; i <= 10; i++{
		for j := 1; j <= 10; j++{
			a := fmt.Sprintf("%d * %d = %d\n", i,j,i*j)
			if _, err := f.WriteString(a); err != nil {
				 panic(err)}
		}
		if _, err := f.WriteString("\n"); err != nil {
			panic(err)
		}
	}
}
