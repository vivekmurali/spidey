package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/commands"
)

func init() {
	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Search for the files",
	Run:   commands.Crawl,
}

// flags:
// seed files
// number limit
// link limit
