package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

type MenuItem interface {
	choose()
	draw() int
}

type Tree struct {
	key rune
	description string
	expanded bool
	children []*Tree
}

func (t Tree) choose() {
	t.expanded = true
}

func (t Tree) draw() int {
	draw_cell(0, 0, 't')
	return 1
}

type Sound struct {
	key rune
	description string
	path string
}

func (s Sound) choose() {
	playMp3(s.path)
}

func (s Sound) draw() int {
	draw_cell(0, 0, 's')
	return 1
}

var topMenu = []MenuItem {
	Tree{ 'M', "Menu", false, nil},
	Sound{ 'S', "Sound", "/home/joshuanelsn/Downloads/zadoc_scream_1.mp3"},
}

var v_header string = "Voice"

var top_tree = []*Tree {
	{'G', "Global", false, global_tree},
	{'A', "Attack", false, nil},
	{'D', "Defend", false, nil},
	{'R', "Repair", false, repair_tree},
	{'B', "Base",   false, nil},
}

var global_tree = []*Tree {
	{'T', "Test",   false, action_tree},
	{'A', "Action", false, nil},
	{'S', "Super",  false, super_tree},
}

var super_tree = []*Tree {
	{'C', "Cheer", false, nil},
	{'L', "Laugh", false, nil},
	{'W', "Wave",  false, nil},
}

var action_tree = []*Tree {
	{'W', "Wave",  false, nil},
	{'T', "Taunt", false, nil},
}

var attack_tree = []*Tree {
	{'F', "Flag", false, nil},
}

var repair_tree = []*Tree {
	{'F', "Flag", false, nil},
	{'B', "Base", false, nil},
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

	var select_vtree []*Tree = top_tree
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
