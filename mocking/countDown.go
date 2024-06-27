package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalKeyWord = "Go!\n"
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type CountDownOperationsSpy struct {
	Calls []string
}

func (s *SpySleeper)Sleep(){
	s.Calls++
}

func (s *CountDownOperationsSpy)Write(p []byte)(n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

func (s *CountDownOperationsSpy)Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func CountDown(buffer io.Writer, sleeper Sleeper){
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
	}
	
	sleeper.Sleep()
	io.WriteString(buffer, finalKeyWord)
}

type DefaultSleeper struct {}

func (d *DefaultSleeper)Sleep(){
	time.Sleep(1 * time.Second)
}

func main(){
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}