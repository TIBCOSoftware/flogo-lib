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
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 23
		case r == 102: // ['f','f']
			return 24
		case 103 <= r && r <= 104: // ['g','h']
			return 22
		case r == 105: // ['i','i']
			return 25
		case 106 <= r && r <= 109: // ['j','m']
			return 22
		case r == 110: // ['n','n']
			return 26
		case 111 <= r && r <= 115: // ['o','s']
			return 22
		case r == 116: // ['t','t']
			return 27
		case 117 <= r && r <= 122: // ['u','z']
			return 22
		case r == 123: // ['{','{']
			return 28
		case r == 124: // ['|','|']
			return 29
		case r == 125: // ['}','}']
			return 30
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
			return 31
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 32
		case r == 92: // ['\','\']
			return 33
		default:
			return 3
		}
	},
	// S4
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 34
		case r == 36: // ['$','$']
			return 34
		case r == 45: // ['-','-']
			return 34
		case r == 46: // ['.','.']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 34
		case 65 <= r && r <= 90: // ['A','Z']
			return 34
		case r == 91: // ['[','[']
			return 35
		case r == 93: // [']',']']
			return 34
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 34
		case r == 123: // ['{','{']
			return 34
		case r == 125: // ['}','}']
			return 34
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
			return 36
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 37
		case r == 92: // ['\','\']
			return 38
		default:
			return 7
		}
	},
	// S8
	func(r rune) int {
		switch {
		case r == 41: // [')',')']
			return 39
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
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 41
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 42
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 42
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
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 43
		case 48 <= r && r <= 57: // ['0','9']
			return 16
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 44
		case 70 <= r && r <= 90: // ['F','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 100: // ['a','d']
			return 22
		case r == 101: // ['e','e']
			return 44
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
			return 45
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 46
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 47
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
		case r == 45: // ['-','-']
			return 22
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
		case r == 45: // ['-','-']
			return 22
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
			return 48
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case r == 97: // ['a','a']
			return 49
		case 98 <= r && r <= 122: // ['b','z']
			return 22
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 101: // ['a','e']
			return 22
		case r == 102: // ['f','f']
			return 50
		case 103 <= r && r <= 122: // ['g','z']
			return 22
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 51
		case 106 <= r && r <= 116: // ['j','t']
			return 22
		case r == 117: // ['u','u']
			return 52
		case 118 <= r && r <= 122: // ['v','z']
			return 22
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 53
		case 115 <= r && r <= 122: // ['s','z']
			return 22
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
		case r == 124: // ['|','|']
			return 54
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		}
		return NoState
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
		default:
			return 55
		}
	},
	// S34
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 34
		case r == 36: // ['$','$']
			return 34
		case r == 45: // ['-','-']
			return 34
		case r == 46: // ['.','.']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 34
		case 65 <= r && r <= 90: // ['A','Z']
			return 34
		case r == 91: // ['[','[']
			return 35
		case r == 93: // [']',']']
			return 34
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 34
		case r == 123: // ['{','{']
			return 34
		case r == 125: // ['}','}']
			return 34
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 56
		default:
			return 57
		}
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		default:
			return 58
		}
	},
	// S39
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 41
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 59
		case r == 45: // ['-','-']
			return 60
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 61
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 62
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
	// S44
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 64
		case r == 45: // ['-','-']
			return 65
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 66
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
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
		case r == 45: // ['-','-']
			return 22
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
	// S49
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S50
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S51
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 69
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 70
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 71
		case 118 <= r && r <= 122: // ['v','z']
			return 22
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 32
		case r == 92: // ['\','\']
			return 33
		default:
			return 3
		}
	},
	// S56
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 34
		case r == 36: // ['$','$']
			return 34
		case r == 45: // ['-','-']
			return 34
		case r == 46: // ['.','.']
			return 34
		case 48 <= r && r <= 57: // ['0','9']
			return 34
		case 65 <= r && r <= 90: // ['A','Z']
			return 34
		case r == 91: // ['[','[']
			return 35
		case r == 93: // [']',']']
			return 34
		case r == 95: // ['_','_']
			return 34
		case 97 <= r && r <= 122: // ['a','z']
			return 34
		case r == 123: // ['{','{']
			return 34
		case r == 125: // ['}','}']
			return 34
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 93: // [']',']']
			return 56
		default:
			return 57
		}
	},
	// S58
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 37
		case r == 92: // ['\','\']
			return 38
		default:
			return 7
		}
	},
	// S59
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 72
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 61
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 61
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 62
		case 65 <= r && r <= 68: // ['A','D']
			return 22
		case r == 69: // ['E','E']
			return 73
		case 70 <= r && r <= 90: // ['F','Z']
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
	// S63
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 74
		case r == 45: // ['-','-']
			return 75
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 76
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
			return 77
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 66
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
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 66
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 78
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 79
		case 116 <= r && r <= 122: // ['t','z']
			return 22
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
		case r == 45: // ['-','-']
			return 22
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
			return 80
		case 109 <= r && r <= 122: // ['m','z']
			return 22
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 81
		case 102 <= r && r <= 122: // ['f','z']
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
		case r == 43: // ['+','+']
			return 82
		case r == 45: // ['-','-']
			return 83
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 84
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
			return 85
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 77
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S79
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
			return 86
		case 102 <= r && r <= 122: // ['f','z']
			return 22
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S81
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S82
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 87
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 84
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
		case r == 46: // ['.','.']
			return 22
		case 48 <= r && r <= 57: // ['0','9']
			return 84
		case 65 <= r && r <= 90: // ['A','Z']
			return 22
		case r == 95: // ['_','_']
			return 22
		case 97 <= r && r <= 122: // ['a','z']
			return 22
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 85
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 22
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
	// S87
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 87
		}
		return NoState
	},
}
