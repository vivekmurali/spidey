package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/search"
)

func Search(cmd *cobra.Command, args []string) {
	m := search.Search(strings.Join(args, " "))
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	for _, v := range keys {
		fmt.Println(v)
	}
}
