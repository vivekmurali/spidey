package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Crawl(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
