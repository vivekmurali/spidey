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

	db.Initalize()

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
		// data := make([]db.Data, 0)
		data := []db.Data{}
		ch := make(chan *db.Data, 5)

		go func() {
			for {
				select {
				case d := <-ch:
					data = append(data, *d)
				}
			}
		}()

		for _, v := range links {
			wg.Add(1)
			go crawler.GetPage(ch, v, &wg)
		}
		wg.Wait()

		//TODO: REMOVE DUPLICATES

		db.Insert(data)
	}
}
