package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Suggestions command
func Suggestions(command string) []string {
	input := antlr.NewInputStream(command)
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

	return redisErrorListener.Expects
}
