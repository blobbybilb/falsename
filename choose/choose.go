package choose

import (
	"falsename/types"

	"github.com/nsf/termbox-go"
)

func DisplayAliasesMenu(options []types.Command, selected int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i, option := range options {
		offsetX := 2
		offsetY := i + 1
		if i == selected {
			termbox.SetCell(offsetX, offsetY, '>', termbox.ColorLightCyan, termbox.ColorDefault)
		}

		toWrite := option.Name + ": " + option.Command

		for _, ch := range toWrite {
			termbox.SetCell(offsetX+2, offsetY, ch, termbox.ColorDefault, termbox.ColorDefault)
			offsetX++
		}
	}
	termbox.Flush()
}

func ChooseAlias(options []types.Command) int {
	termbox.Init()
	defer termbox.Close()

	selected := 0

	DisplayAliasesMenu(options, selected)

	for {
		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyArrowUp:
			if selected > 0 {
				selected--
				DisplayAliasesMenu(options, selected)
			}
		case termbox.KeyArrowDown:
			if selected < (len(options) - 1) {
				selected++
				DisplayAliasesMenu(options, selected)
			}
		case termbox.KeyEnter:
			termbox.Close()
			return selected
		case termbox.KeyEsc:
			termbox.Close()
			return -1
		}
	}
}