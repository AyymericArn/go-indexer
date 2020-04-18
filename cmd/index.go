package cmd

import (
	"fmt"
	"log"
	"os"
	"search/engine"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

// ecrireCmd represents the ecrire command
var ecrireCmd = &cobra.Command{
	Use:   "index",
	Short: "Launches indexation in given folder.",
	Long: `Launches indexation in given folder.
	you may provide relative path with the ./ syntax (.. is not supported yet)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("not enough arguments")
		}

		fmt.Printf("Walking %v", args[0])

		dir := args[0]

		// if is absolute path, turn it into absolute
		// only works for ./
		if string(dir[0]) == "." {
			_, i := utf8.DecodeRuneInString(dir)
			dir = dir[i:]
			mydir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			dir = mydir + dir
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
	rootCmd.AddCommand(ecrireCmd)
}
