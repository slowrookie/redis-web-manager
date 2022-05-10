package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Suggestions command
func Suggestions(command string) []string {
	input := antlr.NewInputStream("")
	lexer := NewRedisLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRedisParser(stream)
	p.BuildParseTrees = true

	redisErrorStrategy := &RedisErrorStrategy{}
	p.SetErrorHandler(redisErrorStrategy)
	p.Command()

	return redisErrorStrategy.Expects
}
