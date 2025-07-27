// cmd/nexuspro/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nexuspro/internal/nexuspro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nexuspro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
