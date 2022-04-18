package commands

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/crawler"
	"github.com/vivekmurali/spidey/pkg/db"
)

func Crawl(cmd *cobra.Command, args []string) {

	count, err := cmd.Flags().GetInt("number_limit")
	if err != nil {
		panic(err)
	}
	if count < 0 {
		count = 3
	}

	for i := 0; i < count; i++ {

		seed, err := cmd.Flags().GetBool("seed")
		if err != nil {
			panic(err)
		}

		var links []string
		if seed {
			home, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			path := filepath.Join(home, "spidey", "seed.txt")
			data, err := os.ReadFile(path)
			if err != nil {
				log.Fatal("Could not read file", err)
			}

			links = strings.Split(string(data), "\n")

			for _, v := range links {
				_, err = url.ParseRequestURI(v)
				if err != nil {
					log.Fatalf("%s is not a valid URL, please update the seed.txt file and try again", v)
				}
			}

		} else {
			links, err = db.GetLinks()
			if err != nil {
				log.Fatal(err)
			}
		}

		var wg sync.WaitGroup
		for _, v := range links {
			wg.Add(1)
			go crawler.GetPage(v, &wg)
		}
		wg.Wait()

	}
}
