package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ScanHits = map[string]int

func Scan(r io.Reader) ScanHits {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	hits := make(map[string]int)
	for s.Scan() {
		hits[s.Text()]++
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return hits
}

func IsText(r io.Reader) bool {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	t := http.DetectContentType(buf.Bytes())

	// TODO: change verification method
	return strings.HasPrefix(t, "text")
}
