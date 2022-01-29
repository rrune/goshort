package util

import (
	"bufio"
	"os"
)

func WaitForExit() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			os.Exit(0)
		}
	}
}
