package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tempCmd)
}

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Create the database file",
	Run: func(cmd *cobra.Command, args []string) {
		// db.RunDB()
	},
}
