package main

import (
	"fmt"

	docgen "github.com/wade-welles/docgen"
)

const cwd := "."

func main() {
	docgen.ParsePDFsInFolder(cwd)
}
