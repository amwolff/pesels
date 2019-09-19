package main

import (
	"fmt"

	"github.com/amwolff/pesels"
)

func main() {
	input := "17281939323"

	pesel, err := pesels.Decode(input)
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
		pesel.DateOfBirth.Format("Mon Jan 2 -0700 MST 2006"),
		pesel.OrdinalNumber,
		pesel.Sex,
		pesel.CheckDigit)
}
