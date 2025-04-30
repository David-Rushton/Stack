package main

type KeyCode int

const (
	KeyNone                KeyCode = 0   // Null character.
	KeyBackspace           KeyCode = 8   // The BACKSPACE key.
	KeyTab                 KeyCode = 9   // The TAB key.
	KeyLineFeed            KeyCode = 10  // The LF key.
	KeyEnter               KeyCode = 13  // The ENTER key.
	KeyEscape              KeyCode = 27  // The ESC (ESCAPE) key.
	KeySpacebar            KeyCode = 32  // The SPACEBAR key.
	KeyExclamationMark     KeyCode = 33  // The exclamation mark key.
	KeyQuotation           KeyCode = 34  // The quotation key.
	KeyNumberSign          KeyCode = 35  // The number sign key.
	KeyDollar              KeyCode = 36  // The dollar key.
	KeyPercent             KeyCode = 37  // The percent key.
	KeyAmpersand           KeyCode = 38  // The ampersand key.
	KeyApostrophe          KeyCode = 39  // The apostrophe key.
	KeyLeftParenthesis     KeyCode = 40  // The left parenthesis key.
	KeyRightParenthesis    KeyCode = 41  // The right parenthesis key.
	KeyAsterisk            KeyCode = 42  // The asterisk key.
	KeyPlus                KeyCode = 43  // The plus key.
	KeyComma               KeyCode = 44  // The comma key.
	KeyHyphen              KeyCode = 45  // The hyphen key.
	KeyPeriod              KeyCode = 46  // The period key.
	KeySlash               KeyCode = 47  // The slash key.
	KeyD0                  KeyCode = 48  // The 0 key.
	KeyD1                  KeyCode = 49  // The 1 key.
	KeyD2                  KeyCode = 50  // The 2 key.
	KeyD3                  KeyCode = 51  // The 3 key.
	KeyD4                  KeyCode = 52  // The 4 key.
	KeyD5                  KeyCode = 53  // The 5 key.
	KeyD6                  KeyCode = 54  // The 6 key.
	KeyD7                  KeyCode = 55  // The 7 key.
	KeyD8                  KeyCode = 56  // The 8 key.
	KeyD9                  KeyCode = 57  // The 9 key.
	KeyColon               KeyCode = 58  // The colon key.
	KeySemicolon           KeyCode = 59  // The semicolon key.
	KeyLessThan            KeyCode = 60  // The less than key.
	KeyEqualTo             KeyCode = 61  // The equal to key.
	KeyGreaterThan         KeyCode = 62  // The greater than key.
	KeyQuestionMark        KeyCode = 63  // The question mark key.
	KeyAtSign              KeyCode = 64  // The at sign key.
	KeyUpperA              KeyCode = 65  // The A key.
	KeyUpperB              KeyCode = 66  // The B key.
	KeyUpperC              KeyCode = 67  // The C key.
	KeyUpperD              KeyCode = 68  // The D key.
	KeyUpperE              KeyCode = 69  // The E key.
	KeyUpperF              KeyCode = 70  // The F key.
	KeyUpperG              KeyCode = 71  // The G key.
	KeyUpperH              KeyCode = 72  // The H key.
	KeyUpperI              KeyCode = 73  // The I key.
	KeyUpperJ              KeyCode = 74  // The J key.
	KeyUpperK              KeyCode = 75  // The K key.
	KeyUpperL              KeyCode = 76  // The L key.
	KeyUpperM              KeyCode = 77  // The M key.
	KeyUpperN              KeyCode = 78  // The N key.
	KeyUpperO              KeyCode = 79  // The O key.
	KeyUpperP              KeyCode = 80  // The P key.
	KeyUpperQ              KeyCode = 81  // The Q key.
	KeyUpperR              KeyCode = 82  // The R key.
	KeyUpperS              KeyCode = 83  // The S key.
	KeyUpperT              KeyCode = 84  // The T key.
	KeyUpperU              KeyCode = 85  // The U key.
	KeyUpperV              KeyCode = 86  // The V key.
	KeyUpperW              KeyCode = 87  // The W key.
	KeyUpperX              KeyCode = 88  // The X key.
	KeyUpperY              KeyCode = 89  // The Y key.
	KeyUpperZ              KeyCode = 90  // The Z key.
	KeyLeftSquareBracket   KeyCode = 91  // left square bracket
	KeyBlackslash          KeyCode = 92  // backslash
	KeyRightSquareBracket  KeyCode = 93  // right square bracket
	KeyCaret               KeyCode = 94  // caret
	KeyUnderscore          KeyCode = 95  // underscore
	KeyGraveAccent         KeyCode = 96  // grave accent
	KeyLowerA              KeyCode = 97  // lowercase a
	KeyLowerB              KeyCode = 98  // lowercase b
	KeyLowerC              KeyCode = 99  // lowercase c
	KeyLowerD              KeyCode = 100 // lowercase d
	KeyLowerE              KeyCode = 101 // lowercase e
	KeyLowerF              KeyCode = 102 // lowercase f
	KeyLowerG              KeyCode = 103 // lowercase g
	KeyLowerH              KeyCode = 104 // lowercase h
	KeyLowerI              KeyCode = 105 // lowercase i
	KeyLowerJ              KeyCode = 106 // lowercase j
	KeyLowerK              KeyCode = 107 // lowercase k
	KeyLowerL              KeyCode = 108 // lowercase l
	KeyLowerM              KeyCode = 109 // lowercase m
	KeyLowerN              KeyCode = 110 // lowercase n
	KeyLowerO              KeyCode = 111 // lowercase o
	KeyLowerP              KeyCode = 112 // lowercase p
	KeyLowerQ              KeyCode = 113 // lowercase q
	KeyLowerR              KeyCode = 114 // lowercase r
	KeyLowerS              KeyCode = 115 // lowercase s
	KeyLowerT              KeyCode = 116 // lowercase t
	KeyLowerU              KeyCode = 117 // lowercase u
	KeyLowerV              KeyCode = 118 // lowercase v
	KeyLowerW              KeyCode = 119 // lowercase w
	KeyLowerX              KeyCode = 120 // lowercase x
	KeyLowerY              KeyCode = 121 // lowercase y
	KeyLowerZ              KeyCode = 122 // lowercase z
	KeyLeftCurlyBrace      KeyCode = 123 // left curly brace
	KeyVerticalBar         KeyCode = 124 // vertical bar
	KeyRightCurlyBrace     KeyCode = 125 // right curly brace
	KeyTilde               KeyCode = 126 // tilde
	KeyDelete              KeyCode = 127 // Delete
	KeyHome                KeyCode = 500 // Home
	KeyInsert              KeyCode = 501 // Insert
	KeyEnd                 KeyCode = 503 // End
	KeyPageUp              KeyCode = 504 // PageUp
	KeyPageDown            KeyCode = 505 // PageDown
	KeyPoundSterlingSymbol KeyCode = 506 // Pound sterling symbol £
	KeyF1                  KeyCode = 508 // F1
	KeyF2                  KeyCode = 509 // F2
	KeyF3                  KeyCode = 510 // F3
	KeyF4                  KeyCode = 511 // F4
	KeyF5                  KeyCode = 512 // F5
	KeyF6                  KeyCode = 513 // F6
	KeyF7                  KeyCode = 514 // F7
	KeyF8                  KeyCode = 515 // F8
	KeyF9                  KeyCode = 516 // F9
	KeyF10                 KeyCode = 517 // F10
	KeyF11                 KeyCode = 518 // F11
	KeyF12                 KeyCode = 519 // F12
	KeyF13                 KeyCode = 521 // F13
	KeyF14                 KeyCode = 522 // F14
	KeyF15                 KeyCode = 523 // F15
	KeyF16                 KeyCode = 524 // F16
	KeyF17                 KeyCode = 525 // F17
	KeyF18                 KeyCode = 526 // F18
	KeyF19                 KeyCode = 527 // F19
	KeyF20                 KeyCode = 528 // F20
	KeyUpArrow             KeyCode = 529 // Up arrow
	KeyDownArrow           KeyCode = 530 // Down arrow
	KeyRightArrow          KeyCode = 531 // Right arrow
	KeyLeftArrow           KeyCode = 532 // Left arrow
	KeyNumPad5             KeyCode = 533 // Num pad 5
)

var (
	KeyCodes = map[int]KeyCode{
		0:   KeyNone,
		8:   KeyBackspace,
		9:   KeyTab,
		10:  KeyLineFeed,
		13:  KeyEnter,
		27:  KeyEscape,
		32:  KeySpacebar,
		33:  KeyExclamationMark,
		34:  KeyQuotation,
		35:  KeyNumberSign,
		36:  KeyDollar,
		37:  KeyPercent,
		38:  KeyAmpersand,
		39:  KeyApostrophe,
		40:  KeyLeftParenthesis,
		41:  KeyRightParenthesis,
		42:  KeyAsterisk,
		43:  KeyPlus,
		44:  KeyComma,
		45:  KeyHyphen,
		46:  KeyPeriod,
		47:  KeySlash,
		48:  KeyD0,
		49:  KeyD1,
		50:  KeyD2,
		51:  KeyD3,
		52:  KeyD4,
		53:  KeyD5,
		54:  KeyD6,
		55:  KeyD7,
		56:  KeyD8,
		57:  KeyD9,
		58:  KeyColon,
		59:  KeySemicolon,
		60:  KeyLessThan,
		61:  KeyEqualTo,
		62:  KeyGreaterThan,
		63:  KeyQuestionMark,
		64:  KeyAtSign,
		65:  KeyUpperA,
		66:  KeyUpperB,
		67:  KeyUpperC,
		68:  KeyUpperD,
		69:  KeyUpperE,
		70:  KeyUpperF,
		71:  KeyUpperG,
		72:  KeyUpperH,
		73:  KeyUpperI,
		74:  KeyUpperJ,
		75:  KeyUpperK,
		76:  KeyUpperL,
		77:  KeyUpperM,
		78:  KeyUpperN,
		79:  KeyUpperO,
		80:  KeyUpperP,
		81:  KeyUpperQ,
		82:  KeyUpperR,
		83:  KeyUpperS,
		84:  KeyUpperT,
		85:  KeyUpperU,
		86:  KeyUpperV,
		87:  KeyUpperW,
		88:  KeyUpperX,
		89:  KeyUpperY,
		90:  KeyUpperZ,
		91:  KeyLeftSquareBracket,
		92:  KeyBlackslash,
		93:  KeyRightSquareBracket,
		94:  KeyCaret,
		95:  KeyUnderscore,
		96:  KeyGraveAccent,
		97:  KeyLowerA,
		98:  KeyLowerB,
		99:  KeyLowerC,
		100: KeyLowerD,
		101: KeyLowerE,
		102: KeyLowerF,
		103: KeyLowerG,
		104: KeyLowerH,
		105: KeyLowerI,
		106: KeyLowerJ,
		107: KeyLowerK,
		108: KeyLowerL,
		109: KeyLowerM,
		110: KeyLowerN,
		111: KeyLowerO,
		112: KeyLowerP,
		113: KeyLowerQ,
		114: KeyLowerR,
		115: KeyLowerS,
		116: KeyLowerT,
		117: KeyLowerU,
		118: KeyLowerV,
		119: KeyLowerW,
		120: KeyLowerX,
		121: KeyLowerY,
		122: KeyLowerZ,
		123: KeyLeftCurlyBrace,
		124: KeyVerticalBar,
		125: KeyRightCurlyBrace,
		126: KeyTilde,
		127: KeyDelete,
		500: KeyHome,
		501: KeyInsert,
		502: KeyDelete,
		503: KeyEnd,
		504: KeyPageUp,
		505: KeyPageDown,
		508: KeyF1,
		509: KeyF2,
		510: KeyF3,
		511: KeyF4,
		512: KeyF5,
		513: KeyF6,
		514: KeyF7,
		515: KeyF8,
		516: KeyF9,
		517: KeyF10,
		518: KeyF11,
		519: KeyF12,
		521: KeyF13,
		522: KeyF14,
		523: KeyF15,
		524: KeyF16,
		525: KeyF17,
		526: KeyF18,
		527: KeyF19,
		528: KeyF20,
		529: KeyUpArrow,
		530: KeyDownArrow,
		531: KeyRightArrow,
		532: KeyLeftArrow,
		533: KeyNumPad5,
	}
)
