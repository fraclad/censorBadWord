package helperFunc

import (
	"io"
	"log"
	"os"
)

func ReadFile(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return (file)
}
