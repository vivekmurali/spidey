package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Version(cmd *cobra.Command, args []string) {
	fmt.Println("Version command")
}
