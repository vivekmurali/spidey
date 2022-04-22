package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vivekmurali/spidey/pkg/db"
	"go.etcd.io/bbolt"
)

func init() {
	rootCmd.AddCommand(tempCmd)
}

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Create the database file",
	Run: func(cmd *cobra.Command, args []string) {
		db.KV.View(func(tx *bbolt.Tx) error {
			c := tx.Bucket([]byte("bucket")).Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
			return nil
		})
	},
}
