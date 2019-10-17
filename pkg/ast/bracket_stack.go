package ast

type BracketStack struct {
	brackets []rune
}

func (s *BracketStack) Push(bracket rune) {
	s.brackets = append(s.brackets, bracket)
}

func (s *BracketStack) Pop() rune {
	if len(s.brackets) > 0 {
		bracket := s.brackets[len(s.brackets)-1]
		s.brackets = s.brackets[:len(s.brackets)-1]
		return bracket
	}
	return rune(-1)
}

func (s *BracketStack) Last() rune {
	if len(s.brackets) > 0 {
		return s.brackets[len(s.brackets)-1]
	}
	return rune(-1)
}
