Package pesels provides utilities for working with Polish PESEL numbers.

[![GoDoc](https://godoc.org/github.com/amwolff/pesels?status.svg)](https://godoc.org/github.com/amwolff/pesels)
[![Build Status](https://travis-ci.org/amwolff/pesels.svg?branch=master)](https://travis-ci.org/amwolff/pesels)

## Validating

```
package main

import (
	"fmt"

	"github.com/amwolff/pesels"
)

func main() {
	fmt.Printf("Valid() == %t\n", pesels.Valid("17281939323"))
}
```

**Output**:

```
Valid() == true
```

## Decoding

```
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
			"Checksum: %d\n",
		input,
		p.DateOfBirth.Format("Mon Jan 2 -0700 MST 2006"),
		p.OrdinalNumber,
		p.Sex,
		p.Checksum)
}
```

**Output**:

```
17281939323:
        Date of birth: Sat Aug 19 +0200 CEST 2017
        Ordinal number: 3932
        Sex: Female
        Checksum: 3
```
