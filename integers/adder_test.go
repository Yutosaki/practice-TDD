package integers

import (
	"fmt"
	"testing"
)
func TestAdder(t *testing.T){
	sum := Adder(2,4)
	expected := 6
	
	if sum != expected {
		t.Errorf("expected %d, but got %d",expected, sum)
	}
}

func ExampleAdder(){
	sum := Adder(1, 4)
	fmt.Println(sum)
	//output: 5
}