package main

import (
	re "regexp"
	"fmt"
)

func main(){
	s := "this is some string"
	pa := re.Compile(`some string`)

	// MatchString return bool 

	if pa.MatchString(s) {

	}

	// Find return result
	result := pa.Find(s)
	fmt.Println(result)
}



