package main

import "github.com/nsf/termbox-go"

var draw_margin int = 2
var draw_y int = draw_margin - 1

type voice_menu struct {
	key rune
	title string
}

type voice_tree struct {
	menu voice_menu
	tree []*voice_tree
}

var fg termbox.Attribute = termbox.ColorWhite
var bg termbox.Attribute = termbox.ColorDefault

var v_header string = "Voice"
var v_header_width int = 18

var top_tree = []*voice_tree {
	{voice_menu{'G', "Global"}, global_tree},
	{voice_menu{'A', "Attack"}, nil},
	{voice_menu{'D', "Defend"}, nil},
	{voice_menu{'R', "Repair"}, nil},
	{voice_menu{'B', "Base"}, nil},
}

var global_tree = []*voice_tree {
	{voice_menu{'T', "Test"}, action_tree},
	{voice_menu{'A', "Action"}, nil},
	{voice_menu{'S', "Super"}, nil},
}

var action_tree = []*voice_tree {
	{voice_menu{'W', "Wave"}, nil},
	{voice_menu{'T', "Taunt"}, nil},
}

var attack_tree = []*voice_tree {
	{voice_menu{'F', "Flag"}, nil},
}

func draw_h1(header string) {
	// Draw Box
	termbox.SetCell(draw_margin-1, draw_y, 0x250C, fg, bg)
	termbox.SetCell(draw_margin-1+v_header_width, draw_y, 0x2510, fg, bg)
	for i := 1; i < v_header_width; i++ {
		termbox.SetCell(draw_margin-1+i, draw_y, 0x2500, fg, bg)
	}
	//termbox.SetCell(draw_margin-1, draw_y, 0x250C, fg, bg)
	draw_y++

	termbox.SetCell(draw_margin-1, draw_y, 0x2502, fg, bg)
	termbox.SetCell(draw_margin-1+v_header_width, draw_y, 0x2502, fg, bg)
	// Draw Header
	for idx, letter := range header {
		termbox.SetCell(draw_margin+idx+1, draw_y, letter, fg, bg)
	}
	draw_y++

	termbox.SetCell(draw_margin-1, draw_y, 0x2514, fg, bg)
	termbox.SetCell(draw_margin-1+v_header_width, draw_y, 0x2518, fg, bg)
	for i := 1; i < v_header_width; i++ {
		termbox.SetCell(draw_margin-1+i, draw_y, 0x2500, fg, bg)
	}
	draw_y++
}

func draw_voice_tree (tree []*voice_tree, level int) int {
	//As we draw down, increase max_y
	level++
	loc_x := draw_margin + 1 + level
	len_subtrees := 0

	len_tree := len(tree)
	termbox.SetCell(loc_x-3, draw_y-1, 0x252C, fg, bg)
	for idx, vtree := range tree {
		// Draw Tree Lines
		last_line := (idx+1) == len_tree
		if last_line {
			termbox.SetCell(loc_x-3, draw_y, 0x2514, fg, bg)
		} else {
			termbox.SetCell(loc_x-3, draw_y, 0x251C, fg, bg)
		}
		termbox.SetCell(loc_x-2, draw_y, 0x2500, fg, bg)

		// Draw Shortcut Key
		termbox.SetCell(loc_x, draw_y, vtree.menu.key, fg, bg)
		termbox.SetCell(loc_x+1, draw_y, ':', fg, bg)

		// Draw Name
		for idx, letter := range vtree.menu.title {
			termbox.SetCell(loc_x+3+idx, draw_y, letter, fg, bg)
		}

		draw_y++

		// Draw Subtree and Line Extension
		if vtree.tree != nil {
			len_subtrees += draw_voice_tree(vtree.tree, level)
			// Draw Line Extension
			for y := 1; y <= len_subtrees; y++ {
				termbox.SetCell(loc_x-3, draw_y-y, 0x2502, fg, bg)
			}
		}
	}

	return len_tree + len_subtrees
}

func draw_soundboard(header string, vtree []*voice_tree) {
	//Menu Header
	draw_h1(header)

	// Menu Lines
	draw_voice_tree(vtree, 1)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	draw_soundboard(v_header, top_tree)
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_soundboard(v_header, top_tree)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
