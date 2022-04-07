package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the database file",
	Run: func(cmd *cobra.Command, args []string) {
		// link to the init file

	},
}
