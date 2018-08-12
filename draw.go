package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

var draw_margin int = 2
var draw_y int = draw_margin - 1
var v_header_width = 18

func reset_draw() {
    draw_y = draw_margin - 1
}

func draw_cell(x, y int, ch rune) {
	termbox.SetCell(x, y, ch, termbox.ColorWhite, termbox.ColorDefault)
}

func draw_soundboard(header string, vtree []*voice_tree, rune_map []rune) {
	draw_h1(header)
	draw_voice_tree(vtree, 1)
	draw_cur_rune_map(rune_map)
	reset_draw()
}

func draw_h1(header string) {
	// Draw box top
	draw_cell(draw_margin-1, draw_y, 0x250C)
	draw_cell(draw_margin-1+v_header_width, draw_y, 0x2510)
	for i := 1; i < v_header_width; i++ {
		draw_cell(draw_margin-1+i, draw_y, 0x2500)
	}
	draw_y++

	draw_cell(draw_margin-1, draw_y, 0x2502)
	draw_cell(draw_margin-1+v_header_width, draw_y, 0x2502)

	// Draw header text plus sidewalls
	for idx, letter := range header {
		draw_cell(draw_margin+idx+1, draw_y, letter)
	}
	draw_y++

	// Draw box bottom
	draw_cell(draw_margin-1, draw_y, 0x2514)
	draw_cell(draw_margin-1+v_header_width, draw_y, 0x2518)
	for i := 1; i < v_header_width; i++ {
		draw_cell(draw_margin-1+i, draw_y, 0x2500)
	}
	draw_y++
}

func draw_voice_tree (tree []*voice_tree, level int) int {
	//As we draw down, increase max_y
	level++
	loc_x := draw_margin + 1 + level
	total_len_subtrees := 0

	draw_cell(loc_x-3, draw_y-1, 0x252C)
	for idx, vtree := range tree {
		// Draw Tree Lines
		last_line := (idx+1) == len(tree)
		if last_line {
			draw_cell(loc_x-3, draw_y, 0x2514)
		} else {
			draw_cell(loc_x-3, draw_y, 0x251C)
		}
		draw_cell(loc_x-2, draw_y, 0x2500)

		// Draw Shortcut Key
		draw_cell(loc_x, draw_y, vtree.menu.key)
		draw_cell(loc_x+1, draw_y, ':')

		// Draw Name
		for idx, letter := range vtree.menu.descrip {
			draw_cell(loc_x+3+idx, draw_y, letter)
		}

		draw_y++

		// Draw Subtree and Line Extension
		len_subtrees := 0
		if vtree.tree != nil && vtree.expanded == true {
			len_subtrees += draw_voice_tree(vtree.tree, level)
			total_len_subtrees += len_subtrees
			if last_line { continue }
			// Draw Line Extension
			for y := 1; y <= len_subtrees; y++ {
				draw_cell(loc_x-3, draw_y-y, 0x2502)
			}
		}
	}

	return len(tree) + total_len_subtrees
}

func draw_cur_rune_map (rune_map []rune) {
	draw_y++
	for idx, r := range rune_map {
		draw_cell(draw_margin+idx, draw_y, r)
	}
}

func rune_to_tree(ch rune, vtree []*voice_tree) []*voice_tree {
	for _, child_tree := range vtree {
		if ch == child_tree.menu.key ||
		    ch == unicode.ToLower(child_tree.menu.key) {
			child_tree.expanded = true
			return child_tree.tree
		}
	}
	return nil
}

func unexpand_tree(vtree []*voice_tree) {
	for _, child_tree := range vtree {
		if child_tree.tree != nil {
			unexpand_tree(child_tree.tree)
			child_tree.expanded = false
		}
	}
}
