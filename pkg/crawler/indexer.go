package crawler

import (
	"strings"
	"sync"

	"github.com/bbalet/stopwords"
	"github.com/vivekmurali/spidey/pkg/db"
	"go.etcd.io/bbolt"
)

func index(u, s string, wg *sync.WaitGroup) {
	defer wg.Done()
	clean := stopwords.CleanString(s, "en", false)
	cleanSlice := strings.Split(clean, " ")

	// _ = v
	_ = db.KV.Batch(func(tx *bbolt.Tx) error {
		for _, v := range cleanSlice {

			old := string(tx.Bucket([]byte("bucket")).Get([]byte(v)))
			if strings.Contains(old, u) {
				continue
			}
			err := tx.Bucket([]byte("bucket")).Put([]byte(v), []byte(old+";"+u))
			if err != nil {
				return nil
			}
		}
		return nil
	})
}

// use map for weightage in search
