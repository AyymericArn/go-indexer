package cmd

import (
	"fmt"
	"log"
	"os"
	"search/engine"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Display files related to a given word.",
	Long:  `Display files related to a given word, with the number of occurence of the word in each file, and the first line where the word appears.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("querying index for \"%v\":\n", args[0])

		var key string = args[0]

		c, err := engine.Dial()
		if err != nil {
			log.Fatal("error with redis connection")
		}

		res, err := redis.Strings(c.Do("ZREVRANGE", key, 0, 2, "WITHSCORES"))

		engine.ShowResults(os.Stdout, key, res)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
