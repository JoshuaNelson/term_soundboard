package main

import "github.com/nsf/termbox-go"

type voice_menu struct {
	key rune
	descrip string
}

type voice_tree struct {
	menu voice_menu
	tree []*voice_tree
}

var v_header string = "Voice"

var top_tree = []*voice_tree {
	{voice_menu{'G', "Global"}, global_tree},
	{voice_menu{'A', "Attack"}, nil},
	{voice_menu{'D', "Defend"}, nil},
	{voice_menu{'R', "Repair"}, repair_tree},
	{voice_menu{'B', "Base"}, nil},
}

var global_tree = []*voice_tree {
	{voice_menu{'T', "Test"}, action_tree},
	{voice_menu{'A', "Action"}, nil},
	{voice_menu{'S', "Super"}, super_tree},
}

var super_tree = []*voice_tree {
	{voice_menu{'C', "Cheer"}, nil},
	{voice_menu{'L', "Laugh"}, nil},
	{voice_menu{'W', "Wave"}, nil},
}

var action_tree = []*voice_tree {
	{voice_menu{'W', "Wave"}, nil},
	{voice_menu{'T', "Taunt"}, nil},
}

var attack_tree = []*voice_tree {
	{voice_menu{'F', "Flag"}, nil},
}

var repair_tree = []*voice_tree {
	{voice_menu{'F', "Flag"}, nil},
	{voice_menu{'B', "Base"}, nil},
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

	var ch_bug []rune
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}
			if ev.Ch == 'c' {
				ch_bug = append(ch_bug, ev.Ch)
				ch_bug[0] = ev.Ch
				v_header = string(ch_bug[0])
			} else {
				v_header = "Voice"
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_soundboard(v_header, top_tree)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
