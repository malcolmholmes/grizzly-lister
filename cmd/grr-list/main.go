package main

import (
	"log"

	"github.com/go-clix/cli"
)

// Version is the current version of the grr command.
// To be overwritten at build time
var Version = "dev"

func main() {
	log.SetFlags(0)

	rootCmd := &cli.Command{
		Use:     "grr-list",
		Short:   "Grizzly",
		Version: Version,
	}

	// workflow commands
	rootCmd.AddCommand(
		foldersCmd(),
		dashboardsCmd(),
		datasourcesCmd(),
	)

	// Run!
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
