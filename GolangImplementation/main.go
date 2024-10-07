// main.go
package main

import (
	"fmt"
	"os"

	"GolangImplementation/gi"
)

func main() {
	p := gi.NewProgram()
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка запуска программы: %v\n", err)
		os.Exit(1)
	}
}
