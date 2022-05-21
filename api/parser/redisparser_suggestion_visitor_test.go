package parser

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/suite"
)

type SuggestionRedisParserVisitorSuite struct {
	suite.Suite
}

func (sut *SuggestionRedisParserVisitorSuite) SetupTest() {

}

func (sut *SuggestionRedisParserVisitorSuite) TestVisitCommand() {
	input := antlr.NewInputStream("")
	lexer := NewRedisLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRedisParser(stream)
	p.BuildParseTrees = true

	redisErrorStrategy := &RedisErrorStrategy{}
	p.SetErrorHandler(redisErrorStrategy)
	p.Command()

	sut.Assert().GreaterOrEqual(len(redisErrorStrategy.Expects), 1)
	fmt.Println(redisErrorStrategy.Expects)
}

func (sut *SuggestionRedisParserVisitorSuite) TestVisitClientCommand() {
	input := antlr.NewInputStream("CLIENT")
	lexer := NewRedisLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRedisParser(stream)
	p.BuildParseTrees = true

	//redisErrorStrategy := &RedisErrorStrategy{}
	//p.SetErrorHandler(redisErrorStrategy)
	redisErrorListener := &SuggestionRedisErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(redisErrorListener)
	p.Command()

	sut.Assert().GreaterOrEqual(len(redisErrorListener.ExpectedTokens), 1)
	fmt.Println(redisErrorListener.ExpectedTokens)
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(SuggestionRedisParserVisitorSuite))
}
