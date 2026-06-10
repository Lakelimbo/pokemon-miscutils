package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Pokémon Miscellaneous Utils",
		Usage: "A bunch of random things for Pokémon data",
		Commands: []*cli.Command{
			&effectivenessCmd,
		},
		EnableShellCompletion: true,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
