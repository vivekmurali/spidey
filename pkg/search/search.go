package search

import (
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/vivekmurali/spidey/pkg/db"
	"go.etcd.io/bbolt"
)

func Search(q string) map[string]int {
	// map for search weight
	m := make(map[string]int)
	clean := stopwords.CleanString(q, "en", false)
	cleanSlice := strings.Split(clean, " ")

	db.KV.View(func(tx *bbolt.Tx) error {
		for _, v := range cleanSlice {
			stored := tx.Bucket([]byte("bucket")).Get([]byte(v))
			storedSlice := strings.Split(string(stored), ";")

			for _, w := range storedSlice {
				m[w]++
			}
		}

		return nil
	})

	return m
}
