package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalKeyWord = "Go!\n"

func CountDown(buffer io.Writer){
	for i := countDownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(buffer, i)
	}
	
	time.Sleep(1 * time.Second)
	io.WriteString(buffer, finalKeyWord)
}

func main(){
	CountDown(os.Stdout)
}