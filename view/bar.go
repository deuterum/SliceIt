package view

import (
	"fmt"
)

func Bar(iteration, total int) {
	barWidth := 50
	progress := float64(iteration) / float64(total)
	filled := int(progress * float64(barWidth))
	empty := barWidth - filled

	bar := fmt.Sprintf("\r[%s%s] %3d%%",
		repeat("=", filled),
		repeat(" ", empty),
		int(progress*100),
	)
	fmt.Print(bar)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
