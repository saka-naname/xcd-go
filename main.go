package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"syscall"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

type DirItem struct {
	Name  string
	IsDir bool
}

type Flags struct {
	ShowHiddenFiles bool
}

func showCursor() {
	fmt.Fprint(os.Stderr, "\x1b[?25h")
}

func hideCursor() {
	fmt.Fprint(os.Stderr, "\x1b[?25l")
}

func clearTerm() {
	fmt.Fprint(os.Stderr, "\x1b[1;0H\x1b[0J\x1b[1;0H")
}

func initTerm() {
	fmt.Fprint(os.Stderr, "\x1b[2J")
}

func renderItems(dir string, items []DirItem, height int, scr int) {
	clearTerm()
	fmt.Fprintf(os.Stderr, "\x1b[32m\x1b[1m%s\x1b[0m\n", dir)

	i := 0 - scr
	for _, item := range items {
		if i < 0 {
			i++
			continue
		} else if i > height-3 {
			break
		}

		if item.IsDir {
			fmt.Fprintf(os.Stderr, "  \x1b[34m\x1b[1m%s\x1b[0m\n", item.Name)
		} else {
			fmt.Fprintf(os.Stderr, "  \x1b[37m%s\x1b[0m\n", item.Name)
		}
		i++
	}
}

func renderCursor(height int, cur int) {
	if cur > 0 {
		fmt.Fprintf(os.Stderr, "\x1b[%d;1H ", cur+1)
	}
	fmt.Fprintf(os.Stderr, "\x1b[%d;1H>", cur+2)
	if cur+3 <= height {
		fmt.Fprintf(os.Stderr, "\x1b[%d;1H ", cur+3)
	}
}

func loadItems(path string, flags Flags) ([]DirItem, int, int) {
	cur := 0
	files, _ := os.ReadDir(path)
	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() {
			if files[j].IsDir() {
				return files[i].Name() < files[j].Name()
			} else {
				return true
			}
		} else {
			if files[j].IsDir() {
				return false
			} else {
				return files[i].Name() < files[j].Name()
			}
		}
	})

	items := []DirItem{}

	if path != "/" {
		items = append(items, DirItem{Name: "..", IsDir: true})
		cur = 1
	}

	for _, f := range files {
		if !flags.ShowHiddenFiles && f.Name()[0] == '.' {
			continue
		}
		items = append(items, DirItem{Name: f.Name(), IsDir: f.IsDir()})
	}

	if len(items) == 1 {
		cur = 0
	}

	return items, cur, 0
}

func main() {
	// Initialization
	flags := Flags{
		ShowHiddenFiles: false,
	}

	_, height, err := term.GetSize(syscall.Stdin)
	if err != nil {
		panic(err)
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// Set current directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirReturn := dir

	// Load files in current directory
	items, cur, scr := loadItems(dir, flags)

	// Render
	hideCursor()
	initTerm()
	renderItems(dir, items, height, scr)
	renderCursor(height, cur)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if char == 'q' {
			dir = dirReturn
			break
		}

		if key == keyboard.KeyEnter {
			break
		}

		switch key {
		case keyboard.KeyArrowDown:
			{
				cur++
				if cur >= height-3 {
					cur--
					if (cur + scr + 1) < len(items) {
						scr++
						renderItems(dir, items, height, scr)
					}
				} else if cur >= len(items) {
					cur--
					break
				}
				renderCursor(height, cur)
			}
		case keyboard.KeyArrowUp:
			{
				cur--
				if cur < 0 {
					cur = 0
					if scr > 0 {
						scr--
						renderItems(dir, items, height, scr)
					} else {
						break
					}
				}
				renderCursor(height, cur)
			}
		case keyboard.KeyArrowLeft:
			{
				if dir == "/" {
					fmt.Fprint(os.Stderr, "\a")
					break
				}
				dir = filepath.Join(dir, "..")
				items, cur, scr = loadItems(dir, flags)
				renderItems(dir, items, height, scr)
				renderCursor(height, cur)
			}
		case keyboard.KeyArrowRight:
			{
				if !items[cur+scr].IsDir {
					fmt.Fprint(os.Stderr, "\a")
					break
				}
				dir = filepath.Join(dir, items[cur+scr].Name)
				items, cur, scr = loadItems(dir, flags)
				renderItems(dir, items, height, scr)
				renderCursor(height, cur)
			}
		}
	}

	clearTerm()
	showCursor()
	fmt.Print(dir)
}
