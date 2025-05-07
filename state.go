package main

type State int

const (
	StateMenu State = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateWin
)

type GameState struct {
	current State
}

func (s *GameState) Set(state State) {
	s.current = state
}

func (s *GameState) Is(state State) bool {
	return s.current == state
}
