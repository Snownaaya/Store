package extensionflow

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func UserInput(prompt string) string{
	
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(text)
}