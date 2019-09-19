package main

import (
	"fmt"

	"github.com/amwolff/pesels"
)

func main() {
	input := "17281939323"

	p, err := pesels.Decode(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"%s:\n\t"+
			"Date of birth: %s\n\t"+
			"Ordinal number: %d\n\t"+
			"Sex: %s\n\t"+
			"Check digit: %d\n",
		input,
		p.DateOfBirth.Format("Mon Jan 2 -0700 MST 2006"),
		p.OrdinalNumber,
		p.Sex,
		p.CheckDigit)
}
