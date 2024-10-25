package main

import (
	"fmt"
	"scope/scoped"
)

func main() {
	fmt.Println(scoped.PublicVar)
	scoped.Exported()

	const CONSTVAR = 45

	// This will not work
	// fmt.Println(scoped.privateVar)

	// This will not work
	// scoped.notExported()

}
