package cmd

import (
	"fmt"
	"log"
	"search/engine"
	"sort"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dumps all recorded entries.",
	Long:  `dumps all the recorded entries in every indexed files.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dumping index")

		c, err := engine.Dial()
		defer c.Close()

		if err != nil {
			log.Fatal("redis connection failed")
		}

		// SCAN
		var keys []string

		// redis query
		var iterator int
		for {
			// thanks @FGM for the example
			arr, err := redis.Values(c.Do("SCAN", iterator))
			if err != nil {
				log.Fatal("error during scan")
			}

			k, _ := redis.Strings(arr[1], nil)

			keys = append(keys, k...)
			iterator, _ = redis.Int(arr[0], nil)
			if iterator == 0 {
				break
			}
		}

		sort.Strings(keys)

		for _, key := range keys {
			fmt.Println(key)
			files, err := engine.Get(c, key)
			if err != nil {
				log.Fatal("get failed")
			}

			for _, file := range files {
				fmt.Println(file)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
