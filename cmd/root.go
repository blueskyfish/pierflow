package cmd

import (
	"github.com/blueskyfish/pierflow/cmd/serve"

	"github.com/spf13/cobra"
)

var CommandRoot = &cobra.Command{
	Use:   "pierflow",
	Short: "Pierflow is a tool manages your docker workflow",
	Long:  `Pierflow is a tool manages your docker workflow, it helps you to build, push, pull and run docker images easily.`,
}

func init() {
	CommandRoot.AddCommand(serve.CommandServe)
	CommandRoot.PersistentFlags().String("log", "info", "The log level for the application. Options are: none, debug, info, warn, error.")
}
