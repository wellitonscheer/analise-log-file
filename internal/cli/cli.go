package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cli struct {
	Input string
}

func NewCli() *Cli {
	return &Cli{}
}

// Wait and return the input
func (c *Cli) Listen() error {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %s", err)
	}

	c.Input = strings.TrimSpace(input)
	return nil
}

func (c *Cli) Parse() {

}
