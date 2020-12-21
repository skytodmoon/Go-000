package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
	path    string
)

func init() {
	rootCmd = newRootCommand()
	rootCmd.Flags().StringVarP(&path, "path", "p", "config/dev.yaml", "input config path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			app, err := BuildApp(path)
			if err != nil {
				panic(err)
			}
			app.Start()
		},
	}
	return cmd
}
