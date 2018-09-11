package main

import (
	"github.com/nsf/termbox-go"
	"unicode"
)

type MenuItem interface {
	Key() rune
	Item() Item
	Choose() Item
	Draw(x, y int) int
	Reset()
}

type Item struct {
	list []MenuItem
}

func (item Item) Draw(x, y int) int {
	origY := y
	totalSubItemLength := 0

	drawCell(x, y-1, 0x252C)

	for idx, subItem := range item.list {
		totalSubItemLength = 0

		lastLine := idx == len(item.list)-1
		if lastLine {
		    drawCell(x, y, 0x2514)
		} else {
		    drawCell(x, y, 0x251C)
		}
		drawCell(x+1, y, 0x2500)

		totalSubItemLength += subItem.Draw(x+1, y)
		y += totalSubItemLength

		if lastLine { continue }
		for i := 1; i < totalSubItemLength; i++ {
			drawCell(x, y-i, 0x2502)
		}
	}

	return totalSubItemLength + (y - origY)
}

type Tree struct {
	key rune
	description string
	expanded bool
	item Item
}

func (tree *Tree) setExpanded(val bool) {
	tree.expanded = val
}

func (tree *Tree) Choose() Item {
	tree.setExpanded(true)
	return tree.item
}

func (tree Tree) Draw(x, y int) int {
	totalSubTreeLength := 1

	drawCell(x+1, y, 0x2500)
	drawText(x+3, y, string(tree.key) + ": " + tree.description)

	if tree.expanded {
		totalSubTreeLength += tree.item.Draw(x+1, y+1)
		totalSubTreeLength--
	}

	return totalSubTreeLength
}

func (tree Tree) Key() rune {
	return tree.key
}

func (tree Tree) Item() Item {
	return tree.item
}

func (tree *Tree) Reset() {
	tree.setExpanded(false)
}

type Sound struct {
	key rune
	description string
	path string
}

func (sound Sound) Choose() Item {
	go playMp3(sound.path)
	return Item{}
}

func (sound Sound) Draw(x, y int) int {
	drawText(x+2, y, string(sound.key) + ": " + sound.description)
	return 1
}

func (sound Sound) Key() rune {
	return sound.key
}

func (sound Sound) Item() Item {
	return Item{}
}

func (sound Sound) Reset() { }

var soundMenu = Item{topMenu}

var topMenu = []MenuItem{
	&Tree{'K', "Kanye West",  false, Item{kanye_tree}},
	&Tree{'S', "Sea of Thieves", false, Item{sot_tree}},
	&Tree{'Z', "Zadoc Allen", false, Item{zadoc_tree}},

}

var soundPath string = "/home/joshuanelsn/Downloads/soundboard/"

var kanye_tree = []MenuItem{
	&Sound{'L', "Lift Yourself", soundPath + "liftyourself_hook.mp3"},
	&Sound{'S', "Scoop",         soundPath + "liftyourself_scoop1.mp3"},
}

var zadoc_tree = []MenuItem{
	Sound{'1', "He he he he he he", soundPath + "zadoc_crazylaugh_1.mp3"},
	Sound{'2', "Hahehahaahehahe hehehehe", soundPath + "zadoc_crazylaugh_2.mp3"},
	Sound{'3', "Hhhhheeehehehe heeh he", soundPath + "zadoc_crazylaugh_3.mp3"},
	Sound{'S', "They seen us", soundPath + "zadoc_theyseenus.mp3"},
	Sound{'T', "They're bringin' things up out of where they come from into the town", soundPath + "zadoc_intothetown.mp3"},
	Sound{'D', "Don't believe me, hey?", soundPath + "zadoc_dontbelievemeeh.mp3"},
	Sound{'G', "Get outta here!", soundPath + "zadoc_getouttahere.mp3"},
	Sound{'O', "GET OUATTA HERE!", soundPath + "zadoc_getouttahere_yell.mp3"},
	Sound{'L', "Get out for your life!", soundPath + "zadoc_getoutforyerlife.mp3"},
	Sound{'Y', "Yeeowhhaa", soundPath + "zadoc_scream_1.mp3"},
	Sound{'A', "Auuuugghh", soundPath + "zadoc_scream_2.mp3"},
}

var sot_tree = []MenuItem{
	Sound{'A', "Drop Anchor", soundPath + "anchor.mp3"},
	Sound{'M', "Megaladon Rush", soundPath + "meg_rushingwater.mp3"},
	Sound{'S', "Megaladon SpawnUnder", soundPath + "meg_spawn_under.mp3"},
}

func selectTree(ch rune, item Item) Item {
	for _, child := range item.list {
		if ch == unicode.ToLower(child.Key()) {
			child.Choose()
			return child.Item()
		}
	}
	return Item{}
}

func unexpandMenu(menu Item) {
	for _, subMenu := range menu.list {
		subMenu.Reset()
		unexpandMenu(subMenu.Item())
	}
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
	drawMenu()
	termbox.Flush()

	var selectItem Item = soundMenu
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			termbox.Clear(termbox.ColorDefault,
			    termbox.ColorDefault)
			if ev.Key == termbox.KeyCtrlX {
				break loop
			}

			selectItem = selectTree(ev.Ch, selectItem)
			if len(selectItem.list) == 0 {
				selectItem = soundMenu
				unexpandMenu(soundMenu)
				runePath = runePath[:0]
			} else {
				runePath = append(runePath,
				    unicode.ToUpper(ev.Ch))
			}

			drawMenu()
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
