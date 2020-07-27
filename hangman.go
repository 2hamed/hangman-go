package main

import "fmt"

const MAX = 7

type Hangman struct {
	word    map[rune][]int
	guessed []rune

	life               int
	hangman            string
	hangmanParts       []interface{}
	hangmanPartsFilled []interface{}
}

func NewHangman(word string) *Hangman {
	h := &Hangman{
		word: make(map[rune][]int),
		life: MAX,
		hangman: `|---------
|       %[1]c
|       %[2]c
|      %[4]c%[3]c%[5]c
|      %[6]c %[7]c
|
|___________` + "\n",
		hangmanParts:       []interface{}{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		hangmanPartsFilled: []interface{}{'|', 'O', '|', '/', '\\', '/', '\\'},
	}

	h.prepare(word)
	return h
}

func (h *Hangman) Hanged() bool {
	h.life--
	return h.life == 0
}

func (h *Hangman) prepare(word string) {
	h.guessed = make([]rune, len(word))

	for i, r := range word {
		h.guessed[i] = '_'

		if arr, ok := h.word[r]; ok {
			h.word[r] = append(arr, i)
		} else {
			h.word[r] = make([]int, 1)
			h.word[r][0] = i
		}
	}
}

func (h *Hangman) Try(r rune) (correct bool, won bool, gameOver bool) {
	if _, ok := h.word[r]; !ok {
		if h.Hanged() {
			return false, false, true
		}

		return false, false, false
	}

	for _, i := range h.word[r] {
		h.guessed[i] = r
	}

	delete(h.word, r)
	if len(h.word) == 0 {
		return true, true, false
	}
	return true, false, false
}

func (h *Hangman) Print() {
	if h.life != MAX {
		h.hangmanParts[MAX-h.life-1] = h.hangmanPartsFilled[MAX-h.life-1]
	}

	fmt.Printf(h.hangman, h.hangmanParts...)
	fmt.Printf("%c\n", h.guessed)
}
