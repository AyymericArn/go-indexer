package engine

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gomodule/redigo/redis"
)

func ScanFile(c redis.Conn, path string, info os.FileInfo, err error) error {
	fi, _ := os.Lstat(path)
	mode := fi.Mode()

	file, _ := os.Open(path)

	if mode.IsRegular() && IsText(file) {
		IndexFile(c, path)
	}

	return err
}

func IndexFile(c redis.Conn, file string) error {
	// data, err := ioutil.ReadFile(file)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("problem when opening file")
	}

	res := Scan(f)

	for word, score := range res {
		err = AddFile(c, word, file, score)
	}
	return err
}

func IndexDir(c redis.Conn, root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		ScanFile(c, path, info, err)
		return err
	})
}
