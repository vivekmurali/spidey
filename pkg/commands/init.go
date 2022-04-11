package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(home+"/spidey", 0765)
	if err != nil {
		fmt.Println("Directory spidey already exists")
	}

	path := filepath.Join(home, "spidey", "spidey.db")

	f, err := os.Create(path)
	if err != nil {
		// fmt.Println("spidey.db already exists")
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}

	path = filepath.Join(home, "spidey", "seed.txt")
	f, err = os.Create(path)
	if err != nil {
		// fmt.Println("seed.txt already exists")
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte("https://vivekmurali.in\nhttps://github.com/vivekmurali"))
	if err != nil {
		// fmt.Println("Could not add default links to seed.txt")
		panic(err)
	}
}
