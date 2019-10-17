package ast

var (
	BracketPair = map[rune]rune{
		'<': '>',
		'(': ')',
		'{': '}',
		'[': ']',
	}

	OpenBrackets  = []rune{'<', '(', '{', '['}
	CloseBrackets = []rune{'>', ')', '}', ']'}
)

func bracketContains(runes []rune, element rune) bool {
	for _, v := range runes {
		if element == v {
			return true
		}
	}
	return false
}
