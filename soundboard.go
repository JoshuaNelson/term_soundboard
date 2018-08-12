package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

type voice_menu struct {
	key rune
	descrip string
}

type voice_tree struct {
	menu voice_menu
	tree []*voice_tree
	expanded bool
}

var v_header string = "Voice"

var top_tree = []*voice_tree {
	{voice_menu{'G', "Global"}, global_tree, false},
	{voice_menu{'A', "Attack"}, nil, false},
	{voice_menu{'D', "Defend"}, nil, false},
	{voice_menu{'R', "Repair"}, repair_tree, false},
	{voice_menu{'B', "Base"}, nil, false},
}

var global_tree = []*voice_tree {
	{voice_menu{'T', "Test"}, action_tree, false},
	{voice_menu{'A', "Action"}, nil, false},
	{voice_menu{'S', "Super"}, super_tree, false},
}

var super_tree = []*voice_tree {
	{voice_menu{'C', "Cheer"}, nil, false},
	{voice_menu{'L', "Laugh"}, nil, false},
	{voice_menu{'W', "Wave"}, nil, false},
}

var action_tree = []*voice_tree {
	{voice_menu{'W', "Wave"}, nil, false},
	{voice_menu{'T', "Taunt"}, nil, false},
}

var attack_tree = []*voice_tree {
	{voice_menu{'F', "Flag"}, nil, false},
}

var repair_tree = []*voice_tree {
	{voice_menu{'F', "Flag"}, nil, false},
	{voice_menu{'B', "Base"}, nil, false},
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var cur_rune_map []rune

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	draw_soundboard(v_header, top_tree, cur_rune_map)
	termbox.Flush()

	var select_vtree []*voice_tree = top_tree
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}

			select_vtree = rune_to_tree(ev.Ch, select_vtree)
			if select_vtree == nil {
				select_vtree = top_tree
				unexpand_tree(top_tree)
				cur_rune_map = cur_rune_map[:0]
			} else {
				cur_rune_map = append(cur_rune_map,
				    unicode.ToUpper(ev.Ch))
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_soundboard(v_header, top_tree, cur_rune_map)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
