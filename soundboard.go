package main

import "github.com/nsf/termbox-go"

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

var top_tree = []*voice_tree{
	{voice_menu{'G', "Global"}, global_tree},
	{voice_menu{'A', "Attack"}, nil},
	{voice_menu{'D', "Defend"}, nil},
	{voice_menu{'R', "Repair"}, nil},
	{voice_menu{'B', "Base"}, nil},
}

var global_tree = []*voice_tree{
	{voice_menu{'T', "Test"}, nil},
	{voice_menu{'A', "Action"}, nil},
	{voice_menu{'S', "Super"}, nil},
}

func draw_h1(x int, y int, header string) {
	termbox.SetCell(x-2, y, 0x250C, fg, bg)
	for idx, letter := range header {
		termbox.SetCell(x+idx, y, letter, fg, bg)
	}
}

func draw_voice_tree(x int, y int, tree []*voice_tree) {
	//As we draw down, increase max_y
	max_y := y
	for _, vtree := range tree {
		termbox.SetCell(x-2, max_y, 0x251C, fg, bg)
		termbox.SetCell(x, max_y, vtree.menu.key, fg, bg)
		termbox.SetCell(x+1, max_y, ':', fg, bg)

		for idx, letter := range vtree.menu.title {
			termbox.SetCell(x+3+idx, max_y, letter, fg, bg)
		}

		max_y++
		if vtree.tree != nil {
			draw_voice_tree(x+1, max_y, vtree.tree)
			menu_len := len(vtree.tree)
			//Draw connecting line
			termbox.SetCell(x-2, max_y, 0x251C, fg, bg)
			for loc_y := 1; loc_y < menu_len; loc_y++ {
				termbox.SetCell(x-2, max_y+1, 0x2502, fg, bg)
				max_y++
			}
			max_y++
		}
	}
	max_y--
	// Draw connection to menu+1
	termbox.SetCell(x-3, y, 0x2514, fg, bg)
	termbox.SetCell(x-2, y, 0x252C, fg, bg)
	// Draw endline
	termbox.SetCell(x-2, max_y, 0x2514, fg, bg)
}

func draw_soundboard(x int, y int, header string, vtree []*voice_tree) {
	//Menu Header
	draw_h1(x+2, y, header)

	// Menu Lines
	draw_voice_tree(x+3, y+1, vtree)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw_margin := 1
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	draw_soundboard(draw_margin, draw_margin, v_header, top_tree)
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_soundboard(draw_margin, draw_margin, v_header, top_tree) 
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
