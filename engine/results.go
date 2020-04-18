package engine

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

// ShowResults : gives results of a search in a formatted table
func ShowResults(w io.Writer, key string, hits []string) {

	tw := new(tabwriter.Writer)
	tw.Init(w, 5, 8, 1, '\t', 0)

	fmt.Fprintln(tw, "File\tCount\tFirst Match\t")

	for i := 0; i < len(hits); i += 2 {

		hit, score := hits[i], hits[i+1]

		// find first line matching word
		r, err := os.Open(hit)

		if err != nil {
			log.Fatal("couldnt open file")
		}

		s := bufio.NewScanner(r)
		s.Split(bufio.ScanLines)
		for s.Scan() {
			line := s.Text()
			if strings.Contains(line, key) {
				fmt.Fprintln(tw, hit+"\t"+string(score)+"\t"+line+"\t")
				break
			}
		}

		// txt := "- " + hit + ":\n"
		// w.Write([]byte(txt))
	}
	tw.Flush()
}
