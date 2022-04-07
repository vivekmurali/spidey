package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the database file",
	Run:   pkg.Init,
}
