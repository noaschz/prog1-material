package rectangles

import "fmt"

//Erwartet sind zwei Seitenlängen "height" und "width".
//Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
//Das Rechteck soll komplett mit "#"-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	//Zeichne eine Zeile auf die Konsole
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
