package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"search/engine"

	"github.com/spf13/cobra"
)

// indexCmd represents the ecrire command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Launches indexation in given folder.",
	Long: `Launches indexation in given folder.
	you may provide relative path or absolute path`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("not enough arguments")
		}

		fmt.Printf("Walking %v", args[0])

		dir, err := filepath.Abs(args[0])
		if err != nil {
			log.Fatal("error converting relative path")
		}

		// path exists ?
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			log.Fatal("bad directory")
		}

		// redis connection
		c, err := engine.Dial()
		defer c.Close()

		if err != nil {
			log.Fatal("redis connection error")
		}

		engine.IndexDir(c, dir)
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
