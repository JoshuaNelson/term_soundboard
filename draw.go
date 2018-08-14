package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

var drawMargin int = 2
var drawY int = drawMargin - 1
var headerWidth = 18

func drawMenu(runePath []rune) {
	// Draw Header
	drawHeader("Voice")

	// Draw Menu
	for _, menu := range topMenu {
		drawY += menu.draw(drawMargin, drawY)
	}

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
	draw_voice_tree(vtree, 1)
	draw_cur_rune_map(rune_map)
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

func draw_voice_tree (tree []*Tree, level int) int {
	//As we draw down, increase max_y
	level++
	loc_x := drawMargin + 1 + level
	total_len_subtrees := 0

	drawCell(loc_x-3, drawY-1, 0x252C)
	for idx, vtree := range tree {
		// Draw Tree Lines
		last_line := (idx+1) == len(tree)
		if last_line {
			drawCell(loc_x-3, drawY, 0x2514)
		} else {
			drawCell(loc_x-3, drawY, 0x251C)
		}
		drawCell(loc_x-2, drawY, 0x2500)

		// Draw Shortcut Key
		drawCell(loc_x, drawY, vtree.key)
		drawCell(loc_x+1, drawY, ':')

		// Draw Name
		for idx, letter := range vtree.description {
			drawCell(loc_x+3+idx, drawY, letter)
		}

		drawY++

		// Draw Subtree and Line Extension
		len_subtrees := 0
		if vtree.children != nil && vtree.expanded == true {
			//len_subtrees += draw_voice_tree(vtree.children, level)
			total_len_subtrees += len_subtrees
			if last_line { continue }
			// Draw Line Extension
			for y := 1; y <= len_subtrees; y++ {
				drawCell(loc_x-3, drawY-y, 0x2502)
			}
		}
	}

	return len(tree) + total_len_subtrees
}

func draw_cur_rune_map (rune_map []rune) {
	drawY++
	for idx, r := range rune_map {
		drawCell(drawMargin+idx, drawY, r)
	}
}

func rune_to_tree(ch rune, vtree []*Tree) []MenuItem{
	for _, child := range vtree {
		if ch == child.key ||
		    ch == unicode.ToLower(child.key) {
			child.expanded = true
			return child.children
		}
	}
	return nil
}

func unexpand_tree(vtree []*Tree) {
	for _, child := range vtree {
		if child.children != nil {
			//unexpand_tree(child.children)
			child.expanded = false
		}
	}
}
