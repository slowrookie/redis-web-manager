package parser

// SuggestionRedisParserVisitor .
// TODO: 根据命令做对应的查询，需要先实现不同数据库对应不同连接的问题（多连接管理）。
type SuggestionRedisParserVisitor struct {
	BaseRedisParserVisitor
}

func (s *SuggestionRedisParserVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return s.VisitChildren(ctx)
}
