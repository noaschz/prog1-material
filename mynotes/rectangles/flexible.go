package rectangles

import "fmt"

func DrawRectangle(height, width int, inner, outer string) {
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			if row == 0 || row == height-1 || column == 0 || column == width-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}
