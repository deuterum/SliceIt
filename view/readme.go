package view

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

func ViewReadme() {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(execPath)

	data, err := os.ReadFile(filepath.Join(execDir, "README.md"))
	if err != nil {
		fmt.Println("Error reading README.md file:", err)
		return
	}

	out, err := glamour.Render(string(data), "dark")
	if err != nil {
		fmt.Println("Error rendering Markdown:", err)
		return
	}

	fmt.Print(out)
}
