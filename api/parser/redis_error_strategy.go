package parser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type RedisErrorStrategy struct {
	antlr.DefaultErrorStrategy
	errorRecoveryMode bool
	Expects           []string
}

func (s *RedisErrorStrategy) beginErrorCondition(recognizer antlr.Parser) {
	s.errorRecoveryMode = true
}

func (s *RedisErrorStrategy) ReportError(recognizer antlr.Parser, e antlr.RecognitionException) {
	// if we've already Reported an error and have not Matched a token
	// yet successfully, don't Report any errors.
	if s.InErrorRecoveryMode(recognizer) {
		return // don't Report spurious errors
	}
	s.beginErrorCondition(recognizer)

	switch t := e.(type) {
	default:
		fmt.Println("unknown recognition error type: " + reflect.TypeOf(e).Name())
		//            fmt.Println(e.stack)
		recognizer.NotifyErrorListeners(e.GetMessage(), e.GetOffendingToken(), e)
	case *antlr.NoViableAltException:
		s.ReportNoViableAlternative(recognizer, t)
	case *antlr.InputMisMatchException:
		s.ReportInputMisMatch(recognizer, t)
	case *antlr.FailedPredicateException:
		s.ReportFailedPredicate(recognizer, t)
	}
}

func (s *RedisErrorStrategy) ReportInputMisMatch(recognizer antlr.Parser, e *antlr.InputMisMatchException) {
	s.setExpects(recognizer)
}

func (s *RedisErrorStrategy) ReportMissingToken(recognizer antlr.Parser) {
	if s.InErrorRecoveryMode(recognizer) {
		return
	}
	s.beginErrorCondition(recognizer)
	s.setExpects(recognizer)
}

func (s *RedisErrorStrategy) setExpects(recognizer antlr.Parser) {
	expecting := s.GetExpectedTokens(recognizer)
	expectStr := expecting.StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), true)
	if len(expectStr) > 0 {
		expectStr = strings.ReplaceAll(strings.ReplaceAll(expectStr, "{", ""), "}", "")
	}
	s.Expects = strings.Split(expectStr, ",")
}
