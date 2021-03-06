package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/commands"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var Version string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spidey",
	Run:   commands.Version,
}
