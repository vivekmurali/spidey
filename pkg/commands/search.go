package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/search"
)

func Search(cmd *cobra.Command, args []string) {
	m := search.Search(strings.Join(args, " "))
	fmt.Println(m)
}
