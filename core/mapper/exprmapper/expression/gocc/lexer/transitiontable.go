// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 36: // ['$','$']
			return 4
		case r == 37: // ['%','%']
			return 5
		case r == 38: // ['&','&']
			return 6
		case r == 39: // [''',''']
			return 7
		case r == 40: // ['(','(']
			return 8
		case r == 41: // [')',')']
			return 9
		case r == 42: // ['*','*']
			return 10
		case r == 43: // ['+','+']
			return 11
		case r == 44: // [',',',']
			return 12
		case r == 45: // ['-','-']
			return 13
		case r == 46: // ['.','.']
			return 14
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 16
		case r == 58: // [':',':']
			return 17
		case r == 60: // ['<','<']
			return 18
		case r == 61: // ['=','=']
			return 19
		case r == 62: // ['>','>']
			return 20
		case r == 63: // ['?','?']
			return 21
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 101: // ['a','e']
			return 22
		case r == 102: // ['f','f']
			return 23
		case 103 <= r && r <= 109: // ['g','m']
			return 22
		case r == 110: // ['n','n']
			return 24
		case 111 <= r && r <= 115: // ['o','s']
			return 22
		case r == 116: // ['t','t']
			return 25
		case 117 <= r && r <= 122: // ['u','z']
			return 22
		case r == 124: // ['|','|']
			return 26
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 27
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 28
		case r == 92: // ['\','\']
			return 29
		default:
			return 3
		}
	},
	// S4
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 30
		case r == 36: // ['$','$']
			return 30
		case r == 45: // ['-','-']
			return 30
		case r == 46: // ['.','.']
			return 30
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 30
		case r == 91: // ['[','[']
			return 31
		case r == 93: // [']',']']
			return 30
		case r == 95: // ['_','_']
			return 30
		case 97 <= r && r <= 122: // ['a','z']
			return 30
		case r == 123: // ['{','{']
			return 30
		case r == 125: // ['}','}']
			return 30
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 32
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 33
		case r == 92: // ['\','\']
			return 34
		default:
			return 7
		}
	},
	// S8
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 35
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 37
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 37
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 38
		case 48 <= r && r <= 57: // ['0','9']
			return 16
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 39
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 39
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 40
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 41
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 42
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case r == 97: // ['a','a']
			return 43
		case 98 <= r && r <= 122: // ['b','z']
			return 22
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 104: // ['a','h']
			return 22
		case r == 105: // ['i','i']
			return 44
		case 106 <= r && r <= 116: // ['j','t']
			return 22
		case r == 117: // ['u','u']
			return 45
		case 118 <= r && r <= 122: // ['v','z']
			return 22
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 113: // ['a','q']
			return 22
		case r == 114: // ['r','r']
			return 46
		case 115 <= r && r <= 122: // ['s','z']
			return 22
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 124: // ['|','|']
			return 47
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		default:
			return 48
		}
	},
	// S30
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 30
		case r == 36: // ['$','$']
			return 30
		case r == 45: // ['-','-']
			return 30
		case r == 46: // ['.','.']
			return 30
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 30
		case r == 91: // ['[','[']
			return 31
		case r == 93: // [']',']']
			return 30
		case r == 95: // ['_','_']
			return 30
		case 97 <= r && r <= 122: // ['a','z']
			return 30
		case r == 123: // ['{','{']
			return 30
		case r == 125: // ['}','}']
			return 30
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 49
		default:
			return 50
		}
	},
	// S32
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		default:
			return 51
		}
	},
	// S35
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 52
		case r == 45: // ['-','-']
			return 52
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 54
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 55
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 55
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 56
		case r == 45: // ['-','-']
			return 56
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 57
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 107: // ['a','k']
			return 22
		case r == 108: // ['l','l']
			return 58
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 107: // ['a','k']
			return 22
		case r == 108: // ['l','l']
			return 59
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 107: // ['a','k']
			return 22
		case r == 108: // ['l','l']
			return 60
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 116: // ['a','t']
			return 22
		case r == 117: // ['u','u']
			return 61
		case 118 <= r && r <= 122: // ['v','z']
			return 22
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 28
		case r == 92: // ['\','\']
			return 29
		default:
			return 3
		}
	},
	// S49
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 30
		case r == 36: // ['$','$']
			return 30
		case r == 45: // ['-','-']
			return 30
		case r == 46: // ['.','.']
			return 30
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 30
		case r == 91: // ['[','[']
			return 31
		case r == 93: // [']',']']
			return 30
		case r == 95: // ['_','_']
			return 30
		case 97 <= r && r <= 122: // ['a','z']
			return 30
		case r == 123: // ['{','{']
			return 30
		case r == 125: // ['}','}']
			return 30
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 49
		default:
			return 50
		}
	},
	// S51
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 33
		case r == 92: // ['\','\']
			return 34
		default:
			return 7
		}
	},
	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 62
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 53
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 54
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 63
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 63
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 64
		case r == 45: // ['-','-']
			return 64
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 65
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 66
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 57
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 114: // ['a','r']
			return 22
		case r == 115: // ['s','s']
			return 67
		case 116 <= r && r <= 122: // ['t','z']
			return 22
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 107: // ['a','k']
			return 22
		case r == 108: // ['l','l']
			return 68
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 69
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 62
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 70
		case r == 45: // ['-','-']
			return 70
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 72
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 65
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 66
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 73
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 74
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 72
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 74
		}
		return NoState
	},
}
