package parser

type SuggestionRedisParserVisitor struct {
	BaseRedisParserVisitor
}

func (s *SuggestionRedisParserVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return s.VisitChildren(ctx)
}
