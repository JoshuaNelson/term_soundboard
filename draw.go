package main

import (
	"github.com/nsf/termbox-go"
//	"unicode"
)

var drawMargin int = 2
var drawY int = drawMargin - 1
var headerWidth = 18

func drawMenu() {
	// Draw Header
	drawHeader("Voice")

	// Draw Menu
	soundMenu.Draw(drawMargin, drawY)

	// Reset Draw
	drawReset()
}

func drawReset() {
    drawY = drawMargin - 1
}

func drawText(x, y int, text string) {
	for idx, ch := range text {
		drawCell(x+idx, y, ch)
	}
}

func drawCell(x, y int, ch rune) {
	termbox.SetCell(x, y, ch, termbox.ColorWhite, termbox.ColorDefault)
}

func draw_soundboard(header string, vtree []*Tree, rune_map []rune) {
	drawHeader(header)
	//draw_voice_tree(vtree, 1)
	//draw_cur_rune_map(rune_map)
	drawReset()
}

func drawHeader(header string) {
	// Draw box top
	drawCell(drawMargin-1, drawY, 0x250C)
	drawCell(drawMargin-1+headerWidth, drawY, 0x2510)
	for i := 1; i < headerWidth; i++ {
		drawCell(drawMargin-1+i, drawY, 0x2500)
	}
	drawY++

	drawCell(drawMargin-1, drawY, 0x2502)
	drawCell(drawMargin-1+headerWidth, drawY, 0x2502)

	// Draw header text plus sidewalls
	for idx, letter := range header {
		drawCell(drawMargin+idx+1, drawY, letter)
	}
	drawY++

	// Draw box bottom
	drawCell(drawMargin-1, drawY, 0x2514)
	drawCell(drawMargin-1+headerWidth, drawY, 0x2518)
	for i := 1; i < headerWidth; i++ {
		drawCell(drawMargin-1+i, drawY, 0x2500)
	}
	drawY++
}
