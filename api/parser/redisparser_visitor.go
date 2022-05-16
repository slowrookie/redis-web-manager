// Code generated from RedisParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RedisParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RedisParser.
type RedisParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RedisParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by RedisParser#bitmap.
	VisitBitmap(ctx *BitmapContext) interface{}

	// Visit a parse tree produced by RedisParser#clusterManagement.
	VisitClusterManagement(ctx *ClusterManagementContext) interface{}

	// Visit a parse tree produced by RedisParser#clusterSub.
	VisitClusterSub(ctx *ClusterSubContext) interface{}

	// Visit a parse tree produced by RedisParser#connectionManagement.
	VisitConnectionManagement(ctx *ConnectionManagementContext) interface{}

	// Visit a parse tree produced by RedisParser#clientSub.
	VisitClientSub(ctx *ClientSubContext) interface{}

	// Visit a parse tree produced by RedisParser#generic.
	VisitGeneric(ctx *GenericContext) interface{}

	// Visit a parse tree produced by RedisParser#objectSub.
	VisitObjectSub(ctx *ObjectSubContext) interface{}

	// Visit a parse tree produced by RedisParser#geospatialIndices.
	VisitGeospatialIndices(ctx *GeospatialIndicesContext) interface{}

	// Visit a parse tree produced by RedisParser#hash.
	VisitHash(ctx *HashContext) interface{}

	// Visit a parse tree produced by RedisParser#hyperloglog.
	VisitHyperloglog(ctx *HyperloglogContext) interface{}

	// Visit a parse tree produced by RedisParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by RedisParser#pubSub.
	VisitPubSub(ctx *PubSubContext) interface{}

	// Visit a parse tree produced by RedisParser#pubSubSub.
	VisitPubSubSub(ctx *PubSubSubContext) interface{}

	// Visit a parse tree produced by RedisParser#scriptingAndFunctions.
	VisitScriptingAndFunctions(ctx *ScriptingAndFunctionsContext) interface{}

	// Visit a parse tree produced by RedisParser#functionSub.
	VisitFunctionSub(ctx *FunctionSubContext) interface{}

	// Visit a parse tree produced by RedisParser#scriptSub.
	VisitScriptSub(ctx *ScriptSubContext) interface{}

	// Visit a parse tree produced by RedisParser#serverManagement.
	VisitServerManagement(ctx *ServerManagementContext) interface{}

	// Visit a parse tree produced by RedisParser#aclSub.
	VisitAclSub(ctx *AclSubContext) interface{}

	// Visit a parse tree produced by RedisParser#commandSub.
	VisitCommandSub(ctx *CommandSubContext) interface{}

	// Visit a parse tree produced by RedisParser#configSub.
	VisitConfigSub(ctx *ConfigSubContext) interface{}

	// Visit a parse tree produced by RedisParser#latencySub.
	VisitLatencySub(ctx *LatencySubContext) interface{}

	// Visit a parse tree produced by RedisParser#memorySub.
	VisitMemorySub(ctx *MemorySubContext) interface{}

	// Visit a parse tree produced by RedisParser#moduleSub.
	VisitModuleSub(ctx *ModuleSubContext) interface{}

	// Visit a parse tree produced by RedisParser#slowlogSub.
	VisitSlowlogSub(ctx *SlowlogSubContext) interface{}

	// Visit a parse tree produced by RedisParser#set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by RedisParser#sortedSet.
	VisitSortedSet(ctx *SortedSetContext) interface{}

	// Visit a parse tree produced by RedisParser#stream.
	VisitStream(ctx *StreamContext) interface{}

	// Visit a parse tree produced by RedisParser#xgroupSub.
	VisitXgroupSub(ctx *XgroupSubContext) interface{}

	// Visit a parse tree produced by RedisParser#xinfoSub.
	VisitXinfoSub(ctx *XinfoSubContext) interface{}

	// Visit a parse tree produced by RedisParser#stringCmd.
	VisitStringCmd(ctx *StringCmdContext) interface{}

	// Visit a parse tree produced by RedisParser#transactions.
	VisitTransactions(ctx *TransactionsContext) interface{}
}
