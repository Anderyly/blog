package main

import (
	"blog/cmd/serve"
	"blog/cmd/toTypecho"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "blog",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(serve.Cmd())
	rootCmd.AddCommand(toTypecho.Cmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
