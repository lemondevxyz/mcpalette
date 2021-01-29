package main

import (
	"fmt"
	"strings"

	"gopkg.in/gookit/color.v1"
)

const template = "\x1b[38;2;0;255;0;48;2;%d;%d;%dmTRUECOLOR\x1b[0m\n"

func main() {
	for _, v := range colors {
		clr := p.Get500(v)

		rgb := color.HexToRGB(clr)
		r, g, b := uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])

		x := 6 - len(v)
		if x > 0 {
			v += strings.Repeat(" ", x)
		}

		str := color.RGB(r, g, b, false).Sprint(v)
		fmt.Print(str, " ")
	}
	fmt.Println()

	for _, s := range order {
		for _, v := range colors {
			if len(p[v][s]) > 0 {
				rgb := color.HexToRGB(p[v][s])
				r, g, b := uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])

				fg := color.RGB(255, 255, 255, false)
				if ((float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255) > 0.5 {
					fg = color.RGB(0, 0, 0, false)
				}

				space := " "
				x := len(v) - 6
				if x > 0 {
					space = strings.Repeat(space, x+1)
				}

				str := fg.Sprint(color.RGB(r, g, b, true).Sprint(p[v][s][1:]))
				fmt.Print(str, space)
			}
		}
		fmt.Println()
	}
}
