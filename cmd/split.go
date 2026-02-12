package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "split lines and print results",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := cmd.Flags().GetInt64("rows")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d\n", lines)
	},
}

// file, err := os.Open("logs/EDLOGFILEREGISTROS_202602100903.txt")
// if err != nil {
// 	log.Fatal(err)
// }
// defer file.Close()

// scanner := bufio.NewScanner(file)

// for scanner.Scan() {
// 	line := scanner.Text()
// 	split := utils.SplitAtPipe(line)
// 	fmt.Printf("\n%#+v\n", split)
// 	break
// }

// if err = scanner.Err(); err != nil {
// 	log.Fatal(err)
// }

func init() {
	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().Int64P("lines", "l", 10, "Number of lines to split")
}
