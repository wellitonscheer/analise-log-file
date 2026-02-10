package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("logs/EDLOGFILEREGISTROS_202602100903.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text() == string(scanner.Bytes()))
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
