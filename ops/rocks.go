package ops

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

// Rocks is the entry point command in the application
var Rocks = &cobra.Command{
	Use:   "rocks",
	Short: "RocksDB Ops CLI",
	Long: `Perform common ops related tasks on one or many RocksDB instances.

Find more details at https://github.com/ind9/rocks`,
}

// CommandHandler is the wrapper interface that all commands to be implement as part of their "Run"
type CommandHandler func(args []string) error

// AttachHandler is a wrapper method for all commands that needs to be exposed
func AttachHandler(handler CommandHandler) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		start := time.Now()
		err := handler(args)
		elapsed := time.Since(start)
		fmt.Printf("This took  %s\n", humanize.RelTime(start, start.Add(elapsed), "", ""))
		if err != nil {
			log.Printf("[Error] %s", err.Error())
			os.Exit(1)
		}
	}
}
