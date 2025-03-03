package main

import (
	"log"
	"os"

	"github.com/fnxpt/roadmap/tools/roadmap/commands"
)

func main() {
	app := commands.App()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
