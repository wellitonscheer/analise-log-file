package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wellitonscheer/analise-log-file/internal/cli"
	"github.com/wellitonscheer/analise-log-file/internal/utils"
)

func main() {
	c := cli.NewCli()
	err := c.Listen()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%#v\n", c.Input)

	file, err := os.Open("logs/EDLOGFILEREGISTROS_202602100903.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		split := utils.SplitAtPipe(line)
		fmt.Printf("\n%#+v\n", split)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
