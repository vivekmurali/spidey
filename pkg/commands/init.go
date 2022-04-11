package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(home+"/spidey", 0765)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(home + "/spidey/" + "spidey.db")
	if err != nil{
		panic(err)
	}
	defer f.Close()
}
