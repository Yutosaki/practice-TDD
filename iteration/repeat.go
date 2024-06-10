package iteration

// Repeat 
//
// Example usage:
// 	result := Repeat("a", 5) //result = "aaaaa"
//
// Parameters:
// 	- character: the string to be repeated
// 	- count: the number of times to repeat the character
//
// Returns:
// 	- A new string consisting of the input character repeated count times.

func Repeat(character string, repeatedCount int)string{
	var repeated string

	for i:=0; i< repeatedCount; i++{
		repeated += character
	}
	return repeated
}
