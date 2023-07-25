package internal

import (
	"bufio"
	"io"
	"os"
)

func ReadFromFile(filePath string) []string {
	byteList := make([]string, 0)

	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	r := bufio.NewReader(fi)
	w := bufio.NewWriter(fi)
	rw := bufio.NewReadWriter(r, w)
	for {
		n, err := rw.ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if string(n) == "" {
			break
		}

		byteList = append(byteList, string(n))
	}
	return byteList
}
