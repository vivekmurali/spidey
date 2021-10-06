package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var Version string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spidey",
	Long:  `All software has versions. This is spidey's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spidey ", Version)
	},
}
