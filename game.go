package main

const (
	STATE_WON GameState = iota
	STATE_LOST
	STATE_WRONG
	STATE_CORRECT
	STATE_REPETETIVE
)

type GameState int

type Game struct {
	hangman *Hangman
	tried   map[rune]struct{}
}

func NewGame(word string) Game {
	h := Game{
		hangman: NewHangman(word),
		tried:   make(map[rune]struct{}),
	}

	return h
}

func (h *Game) Try(r rune) GameState {

	if _, ok := h.tried[r]; ok {
		return STATE_REPETETIVE
	}

	h.tried[r] = struct{}{}

	c, w, o := h.hangman.Try(r)

	if w {
		return STATE_WON
	} else if o {
		return STATE_LOST
	} else if c {
		return STATE_CORRECT
	}

	return STATE_WRONG
}

func (h *Game) PrintStats() {
	h.hangman.Print()
}
