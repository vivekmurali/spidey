package pkg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Init(cmd *cobra.Command, args []string) {
	viper.WriteConfig()
}
