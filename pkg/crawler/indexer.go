package crawler

import "github.com/bbalet/stopwords"

func index(s string) {
	clean := stopwords.CleanString(s, "en", false)
	_ = clean
}
