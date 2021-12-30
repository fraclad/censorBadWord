package helperFunc

import (
	"bufio"
	"io"
	"log"
)

func CreateTabooMap(file io.Reader) map[string]struct{} {
	tabooMap := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tabooMap[scanner.Text()] = struct{}{}
	}
	scanErr := scanner.Err()
	if scanErr != nil {
		log.Fatal(scanErr)
	}
	return (tabooMap)
}
