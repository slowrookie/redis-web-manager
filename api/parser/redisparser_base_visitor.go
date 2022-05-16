// Code generated from RedisParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RedisParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRedisParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRedisParserVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitBitmap(ctx *BitmapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitClusterManagement(ctx *ClusterManagementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitClusterSub(ctx *ClusterSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitConnectionManagement(ctx *ConnectionManagementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitClientSub(ctx *ClientSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitGeneric(ctx *GenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitObjectSub(ctx *ObjectSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitGeospatialIndices(ctx *GeospatialIndicesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitHash(ctx *HashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitHyperloglog(ctx *HyperloglogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitPubSub(ctx *PubSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitPubSubSub(ctx *PubSubSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitScriptingAndFunctions(ctx *ScriptingAndFunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitFunctionSub(ctx *FunctionSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitScriptSub(ctx *ScriptSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitServerManagement(ctx *ServerManagementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitAclSub(ctx *AclSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitCommandSub(ctx *CommandSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitConfigSub(ctx *ConfigSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitLatencySub(ctx *LatencySubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitMemorySub(ctx *MemorySubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitModuleSub(ctx *ModuleSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitSlowlogSub(ctx *SlowlogSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitSet(ctx *SetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitSortedSet(ctx *SortedSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitStream(ctx *StreamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitXgroupSub(ctx *XgroupSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitXinfoSub(ctx *XinfoSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitStringCmd(ctx *StringCmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedisParserVisitor) VisitTransactions(ctx *TransactionsContext) interface{} {
	return v.VisitChildren(ctx)
}
