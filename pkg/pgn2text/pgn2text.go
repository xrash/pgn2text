package pgn2text

import (
	"github.com/xrash/pgn"
	"io"
	"regexp"
	"strings"
)

var __sanRegexp *regexp.Regexp

var __piecesTranslationTable = map[string]string{
	"K": "king",
	"Q": "queen",
	"R": "rook",
	"B": "bishop",
	"N": "knight",
	"P": "pawn",
}

var __promotionTranslationTable = map[string]string{
	"=Q": "queen",
	"=R": "rook",
	"=B": "bishop",
	"=N": "knight",
	"Q": "queen",
	"R": "rook",
	"B": "bishop",
	"N": "knight",
}

var __checksTranslationTable = map[string]string{
	"+": "check",
	"#": "checkmate",
}

var __resultTranslationTable = map[string]string{
	"1/2-1/2": "draw",
	"1-0":     "white wins",
	"0-1":     "black wins",
}

func init() {
	pattern := "^([KQRBNP])?([a-h][0-9]?)?(x)?([a-h][0-9]|O-O|O-O-O)(=?[QRBN])?([+#])?[!?]*$"
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	__sanRegexp = r
}

func Do(input io.Reader) (string, error) {
	g, err := pgn.Parse(input)
	if err != nil {
		return "", err
	}

	text := make([]string, 0)
	for _, m := range g.Movetext.Moves {
		text = append(text, Translate(m.White))
		if m.Black != "" {
			text = append(text, Translate(m.Black))
		}
	}

	text = append(text, __resultTranslationTable[string(g.Movetext.Result)])

	return strings.Join(text, ", "), nil
}

func Translate(san string) string {
	san = strings.TrimSpace(san)
	matches := __sanRegexp.FindStringSubmatch(san)

	piece := matches[1]
	//	from := matches[2]
	captures := matches[3]
	to := matches[4]
	promotes := matches[5]
	checks := matches[6]
	//	engine := matches[7]

	result := make([]string, 0)

	if piece != "" {
		result = append(result, __piecesTranslationTable[piece])
		if captures == "" {
			result = append(result, "to")
		}
	}

	if captures != "" {
		if piece == "" {
			result = append(result, "pawn")
		}
		result = append(result, "takes on")
	}

	if to == "O-O" {
		result = append(result, "castles kingside")
	} else if to == "O-O-O" {
		result = append(result, "castles queenside")
	} else {
		result = append(result, to)
	}

	if promotes != "" {
		result = append(result, "and promotes to "+__promotionTranslationTable[promotes])
	}

	if checks != "" {
		result = append(result, __checksTranslationTable[checks])
	}

	return strings.Join(result, " ")
}
