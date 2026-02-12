package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wellitonscheer/analise-log-file/utils"
)

func main() {
	file, err := os.Open("logs/EDLOGFILEREGISTROS_202602100903.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		split := utils.SplitAtPipe(line)
		fmt.Printf("\n%#+v\n", split)
		break
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
