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

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(SuggestionRedisParserVisitorSuite))
}
