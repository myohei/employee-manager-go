package util

import (
	"bufio"
	"os"
)

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

