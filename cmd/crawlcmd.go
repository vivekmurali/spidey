package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/commands"
)

func init() {
	crawlCmd.Flags().StringP("link_limit", "l", "", "Limit to URLs whose root is the given link")
	crawlCmd.Flags().IntP("number_limit", "n", -1, "Limit the number of layers the crawler goes")
	crawlCmd.Flags().StringP("seedfile", "f", "", "Absolute path to the seed.txt file")
	crawlCmd.Flags().BoolP("seed", "s", false, "Use the seed file as a starting point")

	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Search for the files",
	Run:   commands.Crawl,
}
