package integers

import (
	"fmt"
	"testing"
)
func TestAdder(t *testing.T){
	sum := adder(2,4)
	expected := 6
	
	if sum != expected {
		t.Errorf("expected %d, but got %d",expected, sum)
	}
}

func ExampleAdd(){
	sum := adder(1, 4)
	fmt.Println(sum)
	//output: 5
}