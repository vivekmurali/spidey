package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Version(cmd *cobra.Command, args []string) {
	fmt.Println("Version: 0.0.1")
	// fmt.Println(viper.Get("name"))
}
