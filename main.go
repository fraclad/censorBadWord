package main

import (
	"bufio"
	"censorBadWord/helperFunc"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Println("Plz enter the filepath containing the bad words")
	fmt.Scan(&fileName)
	file := helperFunc.ReadFile(fileName)
	tabooMapCheck := helperFunc.CreateTabooMap(file)
	quit := false
	sentenceReader := bufio.NewReader(os.Stdin)
	for !quit {
		fmt.Println("Write the sentence you want to check [write 'exit' to bail]")
		sentenceToCheck, _ := sentenceReader.ReadString('\n')
		result := helperFunc.CensorBadWord(tabooMapCheck, sentenceToCheck)
		fmt.Println(result)
		if strings.TrimSpace(sentenceToCheck) == "exit" {
			fmt.Println("Bye!")
			quit = true
		}
	}

}
