package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-redis/redis/v9"
	"github.com/samber/lo"
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
		ret := s.Client.ConfigGet(ctx, "databases").Val()
		databases, _ := strconv.Atoi(ret["databases"])
		s.Expects = lo.Map[int, string](lo.Range(databases), func(v int, _ int) string {
			return strconv.Itoa(v)
		})
	}
	return s.VisitChildren(ctx)
}
