package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"strconv"
)

// SuggestionRedisParserVisitor .
type SuggestionRedisParserVisitor struct {
	BaseRedisParserVisitor
	ExpectedTokens []string
	Expects        []string
	DB             *int
	Client         redis.UniversalClient
}

func (s *SuggestionRedisParserVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(s)
}

func (s *SuggestionRedisParserVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		return child.(antlr.ParseTree).Accept(s)
	}
	return nil
}

func (s *SuggestionRedisParserVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return s.VisitChildren(ctx)
}

func (s *SuggestionRedisParserVisitor) VisitConnectionManagement(ctx *ConnectionManagementContext) interface{} {
	if len(s.ExpectedTokens) <= 0 {
		return s.VisitChildren(ctx)
	}

	if s.ExpectedTokens[0] == ctx.parser.GetSymbolicNames()[RedisLexerDB] {
		ctx := context.Background()
		ret, _ := s.Client.ConfigGet(ctx, "databases").Result()
		if len(ret) >= 1 {
			databases, _ := strconv.Atoi((ret[1]).(string))
			for i := 0; i < databases; i++ {
				s.Expects = append(s.Expects, strconv.Itoa(i))
			}
		}
	}
	return s.VisitChildren(ctx)
}
