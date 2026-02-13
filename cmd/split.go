package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wellitonscheer/analise-log-file/utils"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "split lines and print results",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := cmd.Flags().GetInt("lines")
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.Open("logs/EDLOGFILEREGISTROS_202602100903.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for i := 0; i < lines; i++ {
			scanner.Scan()
			line := scanner.Text()
			split := utils.SplitAtPipe(line)
			fmt.Printf("\n%#+v\n", split)
		}

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().IntP("lines", "l", 10, "Number of lines to split")
}
