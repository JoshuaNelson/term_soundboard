package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

type MenuItem interface {
	choose()
	draw(x, y int) int
}

type Tree struct {
	key rune
	description string
	expanded bool
	children []MenuItem
}

func (t Tree) choose() {
	t.expanded = true
}

func (t Tree) draw(x, y int) int {
	totalSubTreeLength := 0

	drawCell(x, y-1, 0x252C)

	drawText(x+3, y, string(t.key) + ": " + t.description)

	return totalSubTreeLength + 1
}

type Sound struct {
	key rune
	description string
	path string
}

func (s Sound) choose() {
	playMp3(s.path)
}

func (s Sound) draw(x, y int) int {
	drawText(x, y, string(s.key) + " " + s.description)
	return 1
}

var topMenu = []MenuItem{
	Tree{'G', "Global",      false, global_tree},
	Tree{'Z', "Zadoc Allen", false, zadoc_tree},
	Tree{'A', "Attack",      false, nil},
	Tree{'D', "Defend",      false, nil},
	Tree{'R', "Repair",      false, repair_tree},
	Tree{'B', "Base",        false, nil},
}

var global_tree = []MenuItem{
	Tree{'T', "Test",        false, action_tree},
	Tree{'A', "Action",      false, nil},
	Tree{'S', "Super",       false, super_tree},
}

var zadoc_tree = []MenuItem{
	Sound{'S', "Scream", "/home/joshuanelsn/Downloads/zadoc_scream_1.mp3"},
}

var super_tree = []MenuItem {
	Tree{'C', "Cheer",       false, nil},
	Tree{'L', "Laugh",       false, nil},
	Tree{'W', "Wave",        false, nil},
}

var action_tree = []MenuItem {
	Tree{'W', "Wave",        false, nil},
	Tree{'T', "Taunt",       false, nil},
}

var attack_tree = []MenuItem {
	Tree{'F', "Flag",        false, nil},
}

var repair_tree = []MenuItem {
	Tree{'F', "Flag",        false, nil},
	Tree{'B', "Base",        false, nil},
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var runePath []rune

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawMenu(runePath)
	termbox.Flush()

	var selectMenu *[]MenuItem = &topMenu
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}

			//selectMenu = rune_to_tree(ev.Ch, selectMenu)
			if selectMenu == nil {
				selectMenu = &topMenu
				//unexpand_tree(topMenu)
				runePath = runePath[:0]
			} else {
				runePath = append(runePath,
				    unicode.ToUpper(ev.Ch))
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			drawMenu(runePath)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
