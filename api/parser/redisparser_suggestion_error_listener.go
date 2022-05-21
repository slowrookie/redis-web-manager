package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
	"strings"
)

type SuggestionRedisErrorListener struct {
	*antlr.DefaultErrorListener
	ExpectedTokens []string
}

func (d *SuggestionRedisErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	expecting := (recognizer.(antlr.Parser)).GetExpectedTokens()
	expectStr := expecting.StringVerbose([]string{}, recognizer.GetSymbolicNames(), true)
	reg := regexp.MustCompile("[{} ]")
	expectStr = reg.ReplaceAllString(expectStr, "")
	d.ExpectedTokens = strings.Split(expectStr, ",")
}
