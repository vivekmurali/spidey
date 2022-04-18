package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/crawler"
	"github.com/vivekmurali/spidey/pkg/db"
)

//Flags:
// Number limit
// Seed path
// Link limit

func Crawl(cmd *cobra.Command, args []string) {

	// Check if already crawled before and crawl from the pages that are not done

	count, err := cmd.Flags().GetInt("number_limit")
	if err != nil {
		panic(err)
	}
	if count < 0 {
		count = 2
	}

	//TODO: move the seed file stuff to another function
	//For loop with count to go to same page
	// Add to DB

	// Work with crawlers from here based on flags
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }

	// path := filepath.Join(home, "spidey", "seed.txt")
	// f, err := os.Open(path)

	// if err != nil {
	// 	panic(err)
	// }

	// sc := bufio.NewScanner(f)

	// var wg sync.WaitGroup
	// for sc.Scan() {
	// 	_, err = url.ParseRequestURI(sc.Text())
	// 	if err != nil {
	// 		fmt.Printf("%s is not a valid URL, please update the seed.txt file and try again", sc.Text())
	// 		os.Exit(1)
	// 	}

	// 	wg.Add(1)
	// 	go crawler.GetPage(sc.Text(), &wg, count)
	// }

	// if err := sc.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error in reading: ", err)
	// }
	// wg.Wait()

	// for i := 0; i < count; i++ {

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

	} else {
		links, err = db.GetLinks()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(links)

	var wg sync.WaitGroup
	for _, v := range links {
		wg.Add(1)
		go crawler.GetPage(v, &wg)
	}
	wg.Wait()

	// }
}
