package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Wait and return the input
func Listen() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %s", err)
	}

	input = strings.TrimSpace(input)

	return input, nil
}
