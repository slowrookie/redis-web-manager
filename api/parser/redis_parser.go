// Code generated from ./RedisParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RedisParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RedisParser struct {
	*antlr.BaseParser
}

var redisparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func redisparserParserInit() {
	staticData := &redisparserParserStaticData
	staticData.literalNames = []string{
		"", "'BITCOUNT'", "'BITFIELD'", "'BITFIELD_RO'", "'BITOP'", "'BITPOS'",
		"'GETBIT'", "'SETBIT'", "'ASKING'", "'CLUSTER'", "'ADDSLOTS'", "'ADDSLOTSRANGE'",
		"'BUMPEPOCH'", "'COUNT-FAILURE-REPORTS'", "'COUNTKEYSINSLOT'", "'DELSLOTS'",
		"'DELSLOTSRANGE'", "'FAILOVER'", "'FLUSHSLOTS'", "'FORGET'", "'GETKEYSINSLOT'",
		"'KEYSLOT'", "'LINKS'", "'MEET'", "'MYID'", "'NODES'", "'REPLICAS'",
		"'REPLICATE'", "'RESET'", "'SAVECONFIG'", "'SET-CONFIG-EPOCH'", "'SETSLOT'",
		"'SHARDS'", "'SLAVES'", "'SLOTS'", "'READONLY'", "'READWRITE'", "'AUTH'",
		"'CLIENT'", "'CACHING'", "'GETNAME'", "'GETREDIR'", "'ID'", "'INFO'",
		"'KILL'", "'NOEVICT'", "'PAUSE'", "'REPLY'", "'SETNAME'", "'TRACKING'",
		"'TRACKINGINFO'", "'UNBLOCK'", "'UNPAUSE'", "'ECHO'", "'HELLO'", "'PING'",
		"'QUIT'", "'SELECT'", "'COPY'", "'DEL'", "'DUMP'", "'EXISTS'", "'EXPIRE'",
		"'EXPIREAT'", "'EXPIRETIME'", "'KEYS'", "'MIGRATE'", "'MOVE'", "'OBJECT'",
		"'ENCODING'", "'FREQ'", "'IDLETIME'", "'REFCOUNT'", "'PERSIST'", "'PEXPIRE'",
		"'PEXPIREAT'", "'PEXPIRETIME'", "'PTTL'", "'RANDOMKEY'", "'RENAME'",
		"'RENAMENX'", "'RESTORE'", "'SCAN'", "'SORT'", "'SORT_RO'", "'TOUCH'",
		"'TTL'", "'TYPE'", "'UNLINK'", "'WAIT'", "'GEOADD'", "'GEODIST'", "'GEOHASH'",
		"'GEOPOS'", "'GEORADIUS'", "'GEORADIUSBYMEMBER'", "'GEORADIUSBYMEMBER_RO'",
		"'GEORADIUS_RO'", "'GEOSEARCH'", "'GEOSEARCHSTORE'", "'HDEL'", "'HEXISTS'",
		"'HGET'", "'HGETALL'", "'HINCRBY'", "'HINCRBYFLOAT'", "'HKEYS'", "'HLEN'",
		"'HMGET'", "'HMSET'", "'HRANDFIELD'", "'HSCAN'", "'HSET'", "'HSETNX'",
		"'HSTRLEN'", "'HVALS'", "'PFADD'", "'PFCOUNT'", "'PFDEBUG'", "'PFMERGE'",
		"'PFSELFTEST'", "'BLMOVE'", "'BLMPOP'", "'BLPOP'", "'BRPOP'", "'BRPOPLPUSH'",
		"'LINDEX'", "'LINSERT'", "'LLEN'", "'LMOVE'", "'LMPOP'", "'LPOP'", "'LPOS'",
		"'LPUSH'", "'LPUSHX'", "'LRANGE'", "'LREM'", "'LSET'", "'LTRIM'", "'RPOP'",
		"'RPOPLPUSH'", "'RPUSH'", "'RPUSHX'", "'PSUBSCRIBE'", "'PUBLISH'", "'PUBSUB'",
		"'CHANNELS'", "'NUMPAT'", "'NUMSUB'", "'SHARDCHANNELS'", "'SHARDNUMSUB'",
		"'PUNSUBSCRIBE'", "'SPUBLISH'", "'SSUBSCRIBE'", "'SUBSCRIBE'", "'SUNSUBSCRIBE'",
		"'UNSUBSCRIBE'", "'EVAL'", "'EVALSHA'", "'EVALSHA_RO'", "'EVAL_RO'",
		"'FCALL'", "'FCALL_RO'", "'FUNCTION'", "'DELETE'", "'FLUSH'", "'LIST'",
		"'LOAD'", "'STATS'", "'SCRIPT'", "'DEBUG'", "'ACL'", "'CAT'", "'DELUSER'",
		"'DRYRUN'", "'GETUSER'", "'HELP'", "'LOG'", "'SAVE'", "'SETUSER'", "'USERS'",
		"'WHOAMI'", "'BGREWRITEAOF'", "'BGSAVE'", "'COMMAND'", "'COUNT'", "'DOCS'",
		"'GETKEYS'", "'GETKEYSANDFLAGS'", "'CONFIG'", "'GET'", "'RESETSTAT'",
		"'REWRITE'", "'SET'", "'DBSIZE'", "'FLUSHALL'", "'FLUSHDB'", "'LASTSAVE'",
		"'LATENCY'", "'DOCTOR'", "'GRAPH'", "'HISTOGRAM'", "'HISTORY'", "'LATEST'",
		"'LOLWUT'", "'MEMORY'", "'MALLOC-STATS'", "'PURGE'", "'USAGE'", "'MODULE'",
		"'LOADEX'", "'UNLOAD'", "'MONITOR'", "'PSYNC'", "'REPLCONF'", "'REPLICAOF'",
		"'RESTORE-ASKING'", "'ROLE'", "'SHUTDOWN'", "'SLAVEOF'", "'SLOWLOG'",
		"'LEN'", "'SWAPDB'", "'SYNC'", "'TIME'", "'SADD'", "'SCARD'", "'SDIFF'",
		"'SDIFFSTORE'", "'SINTER'", "'SINTERCARD'", "'SINTERSTORE'", "'SISMEMBER'",
		"'SMEMBERS'", "'SMISMEMBER'", "'SMOVE'", "'SPOP'", "'SRANDMEMBER'",
		"'SREM'", "'SSCAN'", "'SUNION'", "'SUNIONSTORE'", "'BZMPOP'", "'BZPOPMAX'",
		"'BZPOPMIN'", "'ZADD'", "'ZCARD'", "'ZCOUNT'", "'ZDIFF'", "'ZDIFFSTORE'",
		"'ZINCRBY'", "'ZINTER'", "'ZINTERCARD'", "'ZINTERSTORE'", "'ZLEXCOUNT'",
		"'ZMPOP'", "'ZMSCORE'", "'ZPOPMAX'", "'ZPOPMIN'", "'ZRANDMEMBER'", "'ZRANGE'",
		"'ZRANGEBYLEX'", "'ZRANGEBYSCORE'", "'ZRANGESTORE'", "'ZRANK'", "'ZREM'",
		"'ZREMRANGEBYLEX'", "'ZREMRANGEBYRANK'", "'ZREMRANGEBYSCORE'", "'ZREVRANGE'",
		"'ZREVRANGEBYLEX'", "'ZREVRANGEBYSCORE'", "'ZREVRANK'", "'ZSCAN'", "'ZSCORE'",
		"'ZUNION'", "'ZUNIONSTORE'", "'XACK'", "'XADD'", "'XAUTOCLAIM'", "'XCLAIM'",
		"'XDEL'", "'XGROUP'", "'CREATE'", "'CREATECONSUMER'", "'DELCONSUMER'",
		"'DESTROY'", "'SETID'", "'XINFO'", "'CONSUMERS'", "'GROUPS'", "'STREAM'",
		"'XLEN'", "'XPENDING'", "'XRANGE'", "'XREAD'", "'XREADGROUP'", "'XREVRANGE'",
		"'XSETID'", "'XTRIM'", "'APPEND'", "'DECR'", "'DECRBY'", "'GETDEL'",
		"'GETEX'", "'GETRANGE'", "'GETSET'", "'INCR'", "'INCRBY'", "'INCRBYFLOAT'",
		"'LCS'", "'MGET'", "'MSET'", "'MSETNX'", "'PSETEX'", "'SETEX'", "'SETNX'",
		"'SETRANGE'", "'STRLEN'", "'SUBSTR'", "'DISCARD'", "'EXEC'", "'MULTI'",
		"'UNWATCH'", "'WATCH'", "", "' '",
	}
	staticData.symbolicNames = []string{
		"", "BITCOUNT", "BITFIELD", "BITFIELD_RO", "BITOP", "BITPOS", "GETBIT",
		"SETBIT", "ASKING", "CLUSTER", "ADDSLOTS", "ADDSLOTSRANGE", "BUMPEPOCH",
		"COUNTFAILUREREPORTS", "COUNTKEYSINSLOT", "DELSLOTS", "DELSLOTSRANGE",
		"FAILOVER", "FLUSHSLOTS", "FORGET", "GETKEYSINSLOT", "KEYSLOT", "LINKS",
		"MEET", "MYID", "NODES", "REPLICAS", "REPLICATE", "RESET", "SAVECONFIG",
		"SETCONFIGEPOCH", "SETSLOT", "SHARDS", "SLAVES", "SLOTS", "READONLY",
		"READWRITE", "AUTH", "CLIENT", "CACHING", "GETNAME", "GETREDIR", "ID",
		"INFO", "KILL", "NOEVICT", "PAUSE", "REPLY", "SETNAME", "TRACKING",
		"TRACKINGINFO", "UNBLOCK", "UNPAUSE", "ECHO", "HELLO", "PING", "QUIT",
		"SELECT", "COPY", "DEL", "DUMP", "EXISTS", "EXPIRE", "EXPIREAT", "EXPIRETIME",
		"KEYS", "MIGRATE", "MOVE", "OBJECT", "ENCODING", "FREQ", "IDLETIME",
		"REFCOUNT", "PERSIST", "PEXPIRE", "PEXPIREAT", "PEXPIRETIME", "PTTL",
		"RANDOMKEY", "RENAME", "RENAMENX", "RESTORE", "SCAN", "SORT", "SORT_RO",
		"TOUCH", "TTL", "TYPE", "UNLINK", "WAIT", "GEOADD", "GEODIST", "GEOHASH",
		"GEOPOS", "GEORADIUS", "GEORADIUSBYMEMBER", "GEORADIUSBYMEMBER_RO",
		"GEORADIUS_RO", "GEOSEARCH", "GEOSEARCHSTORE", "HDEL", "HEXISTS", "HGET",
		"HGETALL", "HINCRBY", "HINCRBYFLOAT", "HKEYS", "HLEN", "HMGET", "HMSET",
		"HRANDFIELD", "HSCAN", "HSET", "HSETNX", "HSTRLEN", "HVALS", "PFADD",
		"PFCOUNT", "PFDEBUG", "PFMERGE", "PFSELFTEST", "BLMOVE", "BLMPOP", "BLPOP",
		"BRPOP", "BRPOPLPUSH", "LINDEX", "LINSERT", "LLEN", "LMOVE", "LMPOP",
		"LPOP", "LPOS", "LPUSH", "LPUSHX", "LRANGE", "LREM", "LSET", "LTRIM",
		"RPOP", "RPOPLPUSH", "RPUSH", "RPUSHX", "PSUBSCRIBE", "PUBLISH", "PUBSUB",
		"CHANNELS", "NUMPAT", "NUMSUB", "SHARDCHANNELS", "SHARDNUMSUB", "PUNSUBSCRIBE",
		"SPUBLISH", "SSUBSCRIBE", "SUBSCRIBE", "SUNSUBSCRIBE", "UNSUBSCRIBE",
		"EVAL", "EVALSHA", "EVALSHA_RO", "EVAL_RO", "FCALL", "FCALL_RO", "FUNCTION",
		"DELETE", "FLUSH", "LIST", "LOAD", "STATS", "SCRIPT", "DEBUG", "ACL",
		"CAT", "DELUSER", "DRYRUN", "GETUSER", "HELP", "LOG", "SAVE", "SETUSER",
		"USERS", "WHOAMI", "BGREWRITEAOF", "BGSAVE", "COMMAND", "COUNT", "DOCS",
		"GETKEYS", "GETKEYSANDFLAGS", "CONFIG", "GET", "RESETSTAT", "REWRITE",
		"SET", "DBSIZE", "FLUSHALL", "FLUSHDB", "LASTSAVE", "LATENCY", "DOCTOR",
		"GRAPH", "HISTOGRAM", "HISTORY", "LATEST", "LOLWUT", "MEMORY", "MALLOCSTATS",
		"PURGE", "USAGE", "MODULE", "LOADEX", "UNLOAD", "MONITOR", "PSYNC",
		"REPLCONF", "REPLICAOF", "RESTOREASKING", "ROLE", "SHUTDOWN", "SLAVEOF",
		"SLOWLOG", "LEN", "SWAPDB", "SYNC", "TIME", "SADD", "SCARD", "SDIFF",
		"SDIFFSTORE", "SINTER", "SINTERCARD", "SINTERSTORE", "SISMEMBER", "SMEMBERS",
		"SMISMEMBER", "SMOVE", "SPOP", "SRANDMEMBER", "SREM", "SSCAN", "SUNION",
		"SUNIONSTORE", "BZMPOP", "BZPOPMAX", "BZPOPMIN", "ZADD", "ZCARD", "ZCOUNT",
		"ZDIFF", "ZDIFFSTORE", "ZINCRBY", "ZINTER", "ZINTERCARD", "ZINTERSTORE",
		"ZLEXCOUNT", "ZMPOP", "ZMSCORE", "ZPOPMAX", "ZPOPMIN", "ZRANDMEMBER",
		"ZRANGE", "ZRANGEBYLEX", "ZRANGEBYSCORE", "ZRANGESTORE", "ZRANK", "ZREM",
		"ZREMRANGEBYLEX", "ZREMRANGEBYRANK", "ZREMRANGEBYSCORE", "ZREVRANGE",
		"ZREVRANGEBYLEX", "ZREVRANGEBYSCORE", "ZREVRANK", "ZSCAN", "ZSCORE",
		"ZUNION", "ZUNIONSTORE", "XACK", "XADD", "XAUTOCLAIM", "XCLAIM", "XDEL",
		"XGROUP", "CREATE", "CREATECONSUMER", "DELCONSUMER", "DESTROY", "SETID",
		"XINFO", "CONSUMERS", "GROUPS", "STREAM", "XLEN", "XPENDING", "XRANGE",
		"XREAD", "XREADGROUP", "XREVRANGE", "XSETID", "XTRIM", "APPEND", "DECR",
		"DECRBY", "GETDEL", "GETEX", "GETRANGE", "GETSET", "INCR", "INCRBY",
		"INCRBYFLOAT", "LCS", "MGET", "MSET", "MSETNX", "PSETEX", "SETEX", "SETNX",
		"SETRANGE", "STRLEN", "SUBSTR", "DISCARD", "EXEC", "MULTI", "UNWATCH",
		"WATCH", "DB", "SPACE", "NL",
	}
	staticData.ruleNames = []string{
		"command", "bitmap", "clusterManagement", "clusterSub", "connectionManagement",
		"clientSub", "generic", "objectSub", "geospatialIndices", "hash", "hyperloglog",
		"list", "pubSub", "pubSubSub", "scriptingAndFunctions", "functionSub",
		"scriptSub", "serverManagement", "aclSub", "commandSub", "configSub",
		"latencySub", "memorySub", "moduleSub", "slowlogSub", "set", "sortedSet",
		"stream", "xgroupSub", "xinfoSub", "stringCmd", "transactions",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 327, 278, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 81, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 105, 8, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 139, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 162, 8, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 178, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 192, 8, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 229, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 268, 8, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 0, 0, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 0, 24, 1, 0, 1, 7, 3, 0, 10, 36, 43, 43, 176, 176, 3, 0, 39, 52, 166,
		166, 176, 176, 2, 0, 69, 72, 176, 176, 1, 0, 90, 99, 1, 0, 100, 115, 1,
		0, 116, 120, 1, 0, 121, 142, 2, 0, 146, 150, 176, 176, 5, 0, 44, 44, 60,
		60, 81, 81, 164, 168, 176, 176, 6, 0, 44, 44, 61, 61, 165, 165, 167, 167,
		170, 170, 176, 176, 2, 0, 166, 167, 172, 181, 4, 0, 43, 43, 166, 166, 176,
		176, 185, 188, 2, 0, 176, 176, 190, 193, 3, 0, 28, 28, 176, 176, 199, 203,
		4, 0, 168, 168, 176, 176, 199, 199, 206, 208, 3, 0, 166, 167, 176, 176,
		210, 211, 4, 0, 28, 28, 176, 176, 190, 190, 221, 221, 1, 0, 225, 241, 1,
		0, 242, 276, 2, 0, 176, 176, 283, 287, 2, 0, 176, 176, 289, 291, 3, 0,
		190, 190, 193, 193, 300, 319, 1, 0, 320, 324, 353, 0, 80, 1, 0, 0, 0, 2,
		82, 1, 0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 104, 1, 0, 0,
		0, 10, 106, 1, 0, 0, 0, 12, 138, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 142,
		1, 0, 0, 0, 18, 144, 1, 0, 0, 0, 20, 146, 1, 0, 0, 0, 22, 148, 1, 0, 0,
		0, 24, 161, 1, 0, 0, 0, 26, 163, 1, 0, 0, 0, 28, 177, 1, 0, 0, 0, 30, 179,
		1, 0, 0, 0, 32, 181, 1, 0, 0, 0, 34, 228, 1, 0, 0, 0, 36, 230, 1, 0, 0,
		0, 38, 232, 1, 0, 0, 0, 40, 234, 1, 0, 0, 0, 42, 236, 1, 0, 0, 0, 44, 238,
		1, 0, 0, 0, 46, 240, 1, 0, 0, 0, 48, 242, 1, 0, 0, 0, 50, 244, 1, 0, 0,
		0, 52, 246, 1, 0, 0, 0, 54, 267, 1, 0, 0, 0, 56, 269, 1, 0, 0, 0, 58, 271,
		1, 0, 0, 0, 60, 273, 1, 0, 0, 0, 62, 275, 1, 0, 0, 0, 64, 81, 3, 2, 1,
		0, 65, 81, 3, 4, 2, 0, 66, 81, 3, 8, 4, 0, 67, 81, 3, 12, 6, 0, 68, 81,
		3, 16, 8, 0, 69, 81, 3, 18, 9, 0, 70, 81, 3, 20, 10, 0, 71, 81, 3, 22,
		11, 0, 72, 81, 3, 24, 12, 0, 73, 81, 3, 28, 14, 0, 74, 81, 3, 34, 17, 0,
		75, 81, 3, 50, 25, 0, 76, 81, 3, 52, 26, 0, 77, 81, 3, 54, 27, 0, 78, 81,
		3, 60, 30, 0, 79, 81, 3, 62, 31, 0, 80, 64, 1, 0, 0, 0, 80, 65, 1, 0, 0,
		0, 80, 66, 1, 0, 0, 0, 80, 67, 1, 0, 0, 0, 80, 68, 1, 0, 0, 0, 80, 69,
		1, 0, 0, 0, 80, 70, 1, 0, 0, 0, 80, 71, 1, 0, 0, 0, 80, 72, 1, 0, 0, 0,
		80, 73, 1, 0, 0, 0, 80, 74, 1, 0, 0, 0, 80, 75, 1, 0, 0, 0, 80, 76, 1,
		0, 0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81,
		1, 1, 0, 0, 0, 82, 83, 7, 0, 0, 0, 83, 3, 1, 0, 0, 0, 84, 89, 5, 8, 0,
		0, 85, 86, 5, 9, 0, 0, 86, 87, 5, 326, 0, 0, 87, 89, 3, 6, 3, 0, 88, 84,
		1, 0, 0, 0, 88, 85, 1, 0, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91, 7, 1, 0, 0,
		91, 7, 1, 0, 0, 0, 92, 105, 5, 37, 0, 0, 93, 94, 5, 38, 0, 0, 94, 95, 5,
		326, 0, 0, 95, 105, 3, 10, 5, 0, 96, 105, 5, 53, 0, 0, 97, 105, 5, 54,
		0, 0, 98, 105, 5, 55, 0, 0, 99, 105, 5, 56, 0, 0, 100, 105, 5, 28, 0, 0,
		101, 102, 5, 57, 0, 0, 102, 103, 5, 326, 0, 0, 103, 105, 5, 325, 0, 0,
		104, 92, 1, 0, 0, 0, 104, 93, 1, 0, 0, 0, 104, 96, 1, 0, 0, 0, 104, 97,
		1, 0, 0, 0, 104, 98, 1, 0, 0, 0, 104, 99, 1, 0, 0, 0, 104, 100, 1, 0, 0,
		0, 104, 101, 1, 0, 0, 0, 105, 9, 1, 0, 0, 0, 106, 107, 7, 2, 0, 0, 107,
		11, 1, 0, 0, 0, 108, 139, 5, 58, 0, 0, 109, 139, 5, 59, 0, 0, 110, 139,
		5, 60, 0, 0, 111, 139, 5, 61, 0, 0, 112, 139, 5, 62, 0, 0, 113, 139, 5,
		63, 0, 0, 114, 139, 5, 64, 0, 0, 115, 139, 5, 65, 0, 0, 116, 139, 5, 66,
		0, 0, 117, 139, 5, 67, 0, 0, 118, 119, 5, 68, 0, 0, 119, 120, 5, 326, 0,
		0, 120, 139, 3, 14, 7, 0, 121, 139, 5, 73, 0, 0, 122, 139, 5, 74, 0, 0,
		123, 139, 5, 75, 0, 0, 124, 139, 5, 76, 0, 0, 125, 139, 5, 77, 0, 0, 126,
		139, 5, 78, 0, 0, 127, 139, 5, 79, 0, 0, 128, 139, 5, 80, 0, 0, 129, 139,
		5, 81, 0, 0, 130, 139, 5, 82, 0, 0, 131, 139, 5, 83, 0, 0, 132, 139, 5,
		84, 0, 0, 133, 139, 5, 85, 0, 0, 134, 139, 5, 86, 0, 0, 135, 139, 5, 87,
		0, 0, 136, 139, 5, 88, 0, 0, 137, 139, 5, 89, 0, 0, 138, 108, 1, 0, 0,
		0, 138, 109, 1, 0, 0, 0, 138, 110, 1, 0, 0, 0, 138, 111, 1, 0, 0, 0, 138,
		112, 1, 0, 0, 0, 138, 113, 1, 0, 0, 0, 138, 114, 1, 0, 0, 0, 138, 115,
		1, 0, 0, 0, 138, 116, 1, 0, 0, 0, 138, 117, 1, 0, 0, 0, 138, 118, 1, 0,
		0, 0, 138, 121, 1, 0, 0, 0, 138, 122, 1, 0, 0, 0, 138, 123, 1, 0, 0, 0,
		138, 124, 1, 0, 0, 0, 138, 125, 1, 0, 0, 0, 138, 126, 1, 0, 0, 0, 138,
		127, 1, 0, 0, 0, 138, 128, 1, 0, 0, 0, 138, 129, 1, 0, 0, 0, 138, 130,
		1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138, 133, 1, 0,
		0, 0, 138, 134, 1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		138, 137, 1, 0, 0, 0, 139, 13, 1, 0, 0, 0, 140, 141, 7, 3, 0, 0, 141, 15,
		1, 0, 0, 0, 142, 143, 7, 4, 0, 0, 143, 17, 1, 0, 0, 0, 144, 145, 7, 5,
		0, 0, 145, 19, 1, 0, 0, 0, 146, 147, 7, 6, 0, 0, 147, 21, 1, 0, 0, 0, 148,
		149, 7, 7, 0, 0, 149, 23, 1, 0, 0, 0, 150, 162, 5, 143, 0, 0, 151, 162,
		5, 144, 0, 0, 152, 153, 5, 145, 0, 0, 153, 154, 5, 326, 0, 0, 154, 162,
		3, 26, 13, 0, 155, 162, 5, 151, 0, 0, 156, 162, 5, 152, 0, 0, 157, 162,
		5, 153, 0, 0, 158, 162, 5, 154, 0, 0, 159, 162, 5, 155, 0, 0, 160, 162,
		5, 156, 0, 0, 161, 150, 1, 0, 0, 0, 161, 151, 1, 0, 0, 0, 161, 152, 1,
		0, 0, 0, 161, 155, 1, 0, 0, 0, 161, 156, 1, 0, 0, 0, 161, 157, 1, 0, 0,
		0, 161, 158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162,
		25, 1, 0, 0, 0, 163, 164, 7, 8, 0, 0, 164, 27, 1, 0, 0, 0, 165, 178, 5,
		157, 0, 0, 166, 178, 5, 158, 0, 0, 167, 178, 5, 159, 0, 0, 168, 178, 5,
		160, 0, 0, 169, 178, 5, 161, 0, 0, 170, 178, 5, 162, 0, 0, 171, 172, 5,
		163, 0, 0, 172, 173, 5, 326, 0, 0, 173, 178, 3, 30, 15, 0, 174, 175, 5,
		169, 0, 0, 175, 176, 5, 326, 0, 0, 176, 178, 3, 32, 16, 0, 177, 165, 1,
		0, 0, 0, 177, 166, 1, 0, 0, 0, 177, 167, 1, 0, 0, 0, 177, 168, 1, 0, 0,
		0, 177, 169, 1, 0, 0, 0, 177, 170, 1, 0, 0, 0, 177, 171, 1, 0, 0, 0, 177,
		174, 1, 0, 0, 0, 178, 29, 1, 0, 0, 0, 179, 180, 7, 9, 0, 0, 180, 31, 1,
		0, 0, 0, 181, 182, 7, 10, 0, 0, 182, 33, 1, 0, 0, 0, 183, 184, 5, 171,
		0, 0, 184, 185, 5, 326, 0, 0, 185, 229, 3, 36, 18, 0, 186, 229, 5, 182,
		0, 0, 187, 229, 5, 183, 0, 0, 188, 191, 5, 184, 0, 0, 189, 190, 5, 326,
		0, 0, 190, 192, 3, 38, 19, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0,
		0, 192, 229, 1, 0, 0, 0, 193, 194, 5, 189, 0, 0, 194, 195, 5, 326, 0, 0,
		195, 229, 3, 40, 20, 0, 196, 229, 5, 194, 0, 0, 197, 229, 5, 170, 0, 0,
		198, 229, 5, 17, 0, 0, 199, 229, 5, 195, 0, 0, 200, 229, 5, 196, 0, 0,
		201, 229, 5, 43, 0, 0, 202, 229, 5, 197, 0, 0, 203, 204, 5, 198, 0, 0,
		204, 205, 5, 326, 0, 0, 205, 229, 3, 42, 21, 0, 206, 229, 5, 204, 0, 0,
		207, 208, 5, 205, 0, 0, 208, 209, 5, 326, 0, 0, 209, 229, 3, 44, 22, 0,
		210, 211, 5, 209, 0, 0, 211, 212, 5, 326, 0, 0, 212, 229, 3, 46, 23, 0,
		213, 229, 5, 212, 0, 0, 214, 229, 5, 213, 0, 0, 215, 229, 5, 214, 0, 0,
		216, 229, 5, 215, 0, 0, 217, 229, 5, 216, 0, 0, 218, 229, 5, 217, 0, 0,
		219, 229, 5, 178, 0, 0, 220, 229, 5, 218, 0, 0, 221, 229, 5, 219, 0, 0,
		222, 223, 5, 220, 0, 0, 223, 224, 5, 326, 0, 0, 224, 229, 3, 48, 24, 0,
		225, 229, 5, 222, 0, 0, 226, 229, 5, 223, 0, 0, 227, 229, 5, 224, 0, 0,
		228, 183, 1, 0, 0, 0, 228, 186, 1, 0, 0, 0, 228, 187, 1, 0, 0, 0, 228,
		188, 1, 0, 0, 0, 228, 193, 1, 0, 0, 0, 228, 196, 1, 0, 0, 0, 228, 197,
		1, 0, 0, 0, 228, 198, 1, 0, 0, 0, 228, 199, 1, 0, 0, 0, 228, 200, 1, 0,
		0, 0, 228, 201, 1, 0, 0, 0, 228, 202, 1, 0, 0, 0, 228, 203, 1, 0, 0, 0,
		228, 206, 1, 0, 0, 0, 228, 207, 1, 0, 0, 0, 228, 210, 1, 0, 0, 0, 228,
		213, 1, 0, 0, 0, 228, 214, 1, 0, 0, 0, 228, 215, 1, 0, 0, 0, 228, 216,
		1, 0, 0, 0, 228, 217, 1, 0, 0, 0, 228, 218, 1, 0, 0, 0, 228, 219, 1, 0,
		0, 0, 228, 220, 1, 0, 0, 0, 228, 221, 1, 0, 0, 0, 228, 222, 1, 0, 0, 0,
		228, 225, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229,
		35, 1, 0, 0, 0, 230, 231, 7, 11, 0, 0, 231, 37, 1, 0, 0, 0, 232, 233, 7,
		12, 0, 0, 233, 39, 1, 0, 0, 0, 234, 235, 7, 13, 0, 0, 235, 41, 1, 0, 0,
		0, 236, 237, 7, 14, 0, 0, 237, 43, 1, 0, 0, 0, 238, 239, 7, 15, 0, 0, 239,
		45, 1, 0, 0, 0, 240, 241, 7, 16, 0, 0, 241, 47, 1, 0, 0, 0, 242, 243, 7,
		17, 0, 0, 243, 49, 1, 0, 0, 0, 244, 245, 7, 18, 0, 0, 245, 51, 1, 0, 0,
		0, 246, 247, 7, 19, 0, 0, 247, 53, 1, 0, 0, 0, 248, 268, 5, 277, 0, 0,
		249, 268, 5, 278, 0, 0, 250, 268, 5, 279, 0, 0, 251, 268, 5, 280, 0, 0,
		252, 268, 5, 281, 0, 0, 253, 254, 5, 282, 0, 0, 254, 255, 5, 326, 0, 0,
		255, 268, 3, 56, 28, 0, 256, 257, 5, 288, 0, 0, 257, 258, 5, 326, 0, 0,
		258, 268, 3, 58, 29, 0, 259, 268, 5, 292, 0, 0, 260, 268, 5, 293, 0, 0,
		261, 268, 5, 294, 0, 0, 262, 268, 5, 295, 0, 0, 263, 268, 5, 296, 0, 0,
		264, 268, 5, 297, 0, 0, 265, 268, 5, 298, 0, 0, 266, 268, 5, 299, 0, 0,
		267, 248, 1, 0, 0, 0, 267, 249, 1, 0, 0, 0, 267, 250, 1, 0, 0, 0, 267,
		251, 1, 0, 0, 0, 267, 252, 1, 0, 0, 0, 267, 253, 1, 0, 0, 0, 267, 256,
		1, 0, 0, 0, 267, 259, 1, 0, 0, 0, 267, 260, 1, 0, 0, 0, 267, 261, 1, 0,
		0, 0, 267, 262, 1, 0, 0, 0, 267, 263, 1, 0, 0, 0, 267, 264, 1, 0, 0, 0,
		267, 265, 1, 0, 0, 0, 267, 266, 1, 0, 0, 0, 268, 55, 1, 0, 0, 0, 269, 270,
		7, 20, 0, 0, 270, 57, 1, 0, 0, 0, 271, 272, 7, 21, 0, 0, 272, 59, 1, 0,
		0, 0, 273, 274, 7, 22, 0, 0, 274, 61, 1, 0, 0, 0, 275, 276, 7, 23, 0, 0,
		276, 63, 1, 0, 0, 0, 9, 80, 88, 104, 138, 161, 177, 191, 228, 267,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RedisParserInit initializes any static state used to implement RedisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRedisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RedisParserInit() {
	staticData := &redisparserParserStaticData
	staticData.once.Do(redisparserParserInit)
}

// NewRedisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRedisParser(input antlr.TokenStream) *RedisParser {
	RedisParserInit()
	this := new(RedisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &redisparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RedisParser.g4"

	return this
}

// RedisParser tokens.
const (
	RedisParserEOF                  = antlr.TokenEOF
	RedisParserBITCOUNT             = 1
	RedisParserBITFIELD             = 2
	RedisParserBITFIELD_RO          = 3
	RedisParserBITOP                = 4
	RedisParserBITPOS               = 5
	RedisParserGETBIT               = 6
	RedisParserSETBIT               = 7
	RedisParserASKING               = 8
	RedisParserCLUSTER              = 9
	RedisParserADDSLOTS             = 10
	RedisParserADDSLOTSRANGE        = 11
	RedisParserBUMPEPOCH            = 12
	RedisParserCOUNTFAILUREREPORTS  = 13
	RedisParserCOUNTKEYSINSLOT      = 14
	RedisParserDELSLOTS             = 15
	RedisParserDELSLOTSRANGE        = 16
	RedisParserFAILOVER             = 17
	RedisParserFLUSHSLOTS           = 18
	RedisParserFORGET               = 19
	RedisParserGETKEYSINSLOT        = 20
	RedisParserKEYSLOT              = 21
	RedisParserLINKS                = 22
	RedisParserMEET                 = 23
	RedisParserMYID                 = 24
	RedisParserNODES                = 25
	RedisParserREPLICAS             = 26
	RedisParserREPLICATE            = 27
	RedisParserRESET                = 28
	RedisParserSAVECONFIG           = 29
	RedisParserSETCONFIGEPOCH       = 30
	RedisParserSETSLOT              = 31
	RedisParserSHARDS               = 32
	RedisParserSLAVES               = 33
	RedisParserSLOTS                = 34
	RedisParserREADONLY             = 35
	RedisParserREADWRITE            = 36
	RedisParserAUTH                 = 37
	RedisParserCLIENT               = 38
	RedisParserCACHING              = 39
	RedisParserGETNAME              = 40
	RedisParserGETREDIR             = 41
	RedisParserID                   = 42
	RedisParserINFO                 = 43
	RedisParserKILL                 = 44
	RedisParserNOEVICT              = 45
	RedisParserPAUSE                = 46
	RedisParserREPLY                = 47
	RedisParserSETNAME              = 48
	RedisParserTRACKING             = 49
	RedisParserTRACKINGINFO         = 50
	RedisParserUNBLOCK              = 51
	RedisParserUNPAUSE              = 52
	RedisParserECHO                 = 53
	RedisParserHELLO                = 54
	RedisParserPING                 = 55
	RedisParserQUIT                 = 56
	RedisParserSELECT               = 57
	RedisParserCOPY                 = 58
	RedisParserDEL                  = 59
	RedisParserDUMP                 = 60
	RedisParserEXISTS               = 61
	RedisParserEXPIRE               = 62
	RedisParserEXPIREAT             = 63
	RedisParserEXPIRETIME           = 64
	RedisParserKEYS                 = 65
	RedisParserMIGRATE              = 66
	RedisParserMOVE                 = 67
	RedisParserOBJECT               = 68
	RedisParserENCODING             = 69
	RedisParserFREQ                 = 70
	RedisParserIDLETIME             = 71
	RedisParserREFCOUNT             = 72
	RedisParserPERSIST              = 73
	RedisParserPEXPIRE              = 74
	RedisParserPEXPIREAT            = 75
	RedisParserPEXPIRETIME          = 76
	RedisParserPTTL                 = 77
	RedisParserRANDOMKEY            = 78
	RedisParserRENAME               = 79
	RedisParserRENAMENX             = 80
	RedisParserRESTORE              = 81
	RedisParserSCAN                 = 82
	RedisParserSORT                 = 83
	RedisParserSORT_RO              = 84
	RedisParserTOUCH                = 85
	RedisParserTTL                  = 86
	RedisParserTYPE                 = 87
	RedisParserUNLINK               = 88
	RedisParserWAIT                 = 89
	RedisParserGEOADD               = 90
	RedisParserGEODIST              = 91
	RedisParserGEOHASH              = 92
	RedisParserGEOPOS               = 93
	RedisParserGEORADIUS            = 94
	RedisParserGEORADIUSBYMEMBER    = 95
	RedisParserGEORADIUSBYMEMBER_RO = 96
	RedisParserGEORADIUS_RO         = 97
	RedisParserGEOSEARCH            = 98
	RedisParserGEOSEARCHSTORE       = 99
	RedisParserHDEL                 = 100
	RedisParserHEXISTS              = 101
	RedisParserHGET                 = 102
	RedisParserHGETALL              = 103
	RedisParserHINCRBY              = 104
	RedisParserHINCRBYFLOAT         = 105
	RedisParserHKEYS                = 106
	RedisParserHLEN                 = 107
	RedisParserHMGET                = 108
	RedisParserHMSET                = 109
	RedisParserHRANDFIELD           = 110
	RedisParserHSCAN                = 111
	RedisParserHSET                 = 112
	RedisParserHSETNX               = 113
	RedisParserHSTRLEN              = 114
	RedisParserHVALS                = 115
	RedisParserPFADD                = 116
	RedisParserPFCOUNT              = 117
	RedisParserPFDEBUG              = 118
	RedisParserPFMERGE              = 119
	RedisParserPFSELFTEST           = 120
	RedisParserBLMOVE               = 121
	RedisParserBLMPOP               = 122
	RedisParserBLPOP                = 123
	RedisParserBRPOP                = 124
	RedisParserBRPOPLPUSH           = 125
	RedisParserLINDEX               = 126
	RedisParserLINSERT              = 127
	RedisParserLLEN                 = 128
	RedisParserLMOVE                = 129
	RedisParserLMPOP                = 130
	RedisParserLPOP                 = 131
	RedisParserLPOS                 = 132
	RedisParserLPUSH                = 133
	RedisParserLPUSHX               = 134
	RedisParserLRANGE               = 135
	RedisParserLREM                 = 136
	RedisParserLSET                 = 137
	RedisParserLTRIM                = 138
	RedisParserRPOP                 = 139
	RedisParserRPOPLPUSH            = 140
	RedisParserRPUSH                = 141
	RedisParserRPUSHX               = 142
	RedisParserPSUBSCRIBE           = 143
	RedisParserPUBLISH              = 144
	RedisParserPUBSUB               = 145
	RedisParserCHANNELS             = 146
	RedisParserNUMPAT               = 147
	RedisParserNUMSUB               = 148
	RedisParserSHARDCHANNELS        = 149
	RedisParserSHARDNUMSUB          = 150
	RedisParserPUNSUBSCRIBE         = 151
	RedisParserSPUBLISH             = 152
	RedisParserSSUBSCRIBE           = 153
	RedisParserSUBSCRIBE            = 154
	RedisParserSUNSUBSCRIBE         = 155
	RedisParserUNSUBSCRIBE          = 156
	RedisParserEVAL                 = 157
	RedisParserEVALSHA              = 158
	RedisParserEVALSHA_RO           = 159
	RedisParserEVAL_RO              = 160
	RedisParserFCALL                = 161
	RedisParserFCALL_RO             = 162
	RedisParserFUNCTION             = 163
	RedisParserDELETE               = 164
	RedisParserFLUSH                = 165
	RedisParserLIST                 = 166
	RedisParserLOAD                 = 167
	RedisParserSTATS                = 168
	RedisParserSCRIPT               = 169
	RedisParserDEBUG                = 170
	RedisParserACL                  = 171
	RedisParserCAT                  = 172
	RedisParserDELUSER              = 173
	RedisParserDRYRUN               = 174
	RedisParserGETUSER              = 175
	RedisParserHELP                 = 176
	RedisParserLOG                  = 177
	RedisParserSAVE                 = 178
	RedisParserSETUSER              = 179
	RedisParserUSERS                = 180
	RedisParserWHOAMI               = 181
	RedisParserBGREWRITEAOF         = 182
	RedisParserBGSAVE               = 183
	RedisParserCOMMAND              = 184
	RedisParserCOUNT                = 185
	RedisParserDOCS                 = 186
	RedisParserGETKEYS              = 187
	RedisParserGETKEYSANDFLAGS      = 188
	RedisParserCONFIG               = 189
	RedisParserGET                  = 190
	RedisParserRESETSTAT            = 191
	RedisParserREWRITE              = 192
	RedisParserSET                  = 193
	RedisParserDBSIZE               = 194
	RedisParserFLUSHALL             = 195
	RedisParserFLUSHDB              = 196
	RedisParserLASTSAVE             = 197
	RedisParserLATENCY              = 198
	RedisParserDOCTOR               = 199
	RedisParserGRAPH                = 200
	RedisParserHISTOGRAM            = 201
	RedisParserHISTORY              = 202
	RedisParserLATEST               = 203
	RedisParserLOLWUT               = 204
	RedisParserMEMORY               = 205
	RedisParserMALLOCSTATS          = 206
	RedisParserPURGE                = 207
	RedisParserUSAGE                = 208
	RedisParserMODULE               = 209
	RedisParserLOADEX               = 210
	RedisParserUNLOAD               = 211
	RedisParserMONITOR              = 212
	RedisParserPSYNC                = 213
	RedisParserREPLCONF             = 214
	RedisParserREPLICAOF            = 215
	RedisParserRESTOREASKING        = 216
	RedisParserROLE                 = 217
	RedisParserSHUTDOWN             = 218
	RedisParserSLAVEOF              = 219
	RedisParserSLOWLOG              = 220
	RedisParserLEN                  = 221
	RedisParserSWAPDB               = 222
	RedisParserSYNC                 = 223
	RedisParserTIME                 = 224
	RedisParserSADD                 = 225
	RedisParserSCARD                = 226
	RedisParserSDIFF                = 227
	RedisParserSDIFFSTORE           = 228
	RedisParserSINTER               = 229
	RedisParserSINTERCARD           = 230
	RedisParserSINTERSTORE          = 231
	RedisParserSISMEMBER            = 232
	RedisParserSMEMBERS             = 233
	RedisParserSMISMEMBER           = 234
	RedisParserSMOVE                = 235
	RedisParserSPOP                 = 236
	RedisParserSRANDMEMBER          = 237
	RedisParserSREM                 = 238
	RedisParserSSCAN                = 239
	RedisParserSUNION               = 240
	RedisParserSUNIONSTORE          = 241
	RedisParserBZMPOP               = 242
	RedisParserBZPOPMAX             = 243
	RedisParserBZPOPMIN             = 244
	RedisParserZADD                 = 245
	RedisParserZCARD                = 246
	RedisParserZCOUNT               = 247
	RedisParserZDIFF                = 248
	RedisParserZDIFFSTORE           = 249
	RedisParserZINCRBY              = 250
	RedisParserZINTER               = 251
	RedisParserZINTERCARD           = 252
	RedisParserZINTERSTORE          = 253
	RedisParserZLEXCOUNT            = 254
	RedisParserZMPOP                = 255
	RedisParserZMSCORE              = 256
	RedisParserZPOPMAX              = 257
	RedisParserZPOPMIN              = 258
	RedisParserZRANDMEMBER          = 259
	RedisParserZRANGE               = 260
	RedisParserZRANGEBYLEX          = 261
	RedisParserZRANGEBYSCORE        = 262
	RedisParserZRANGESTORE          = 263
	RedisParserZRANK                = 264
	RedisParserZREM                 = 265
	RedisParserZREMRANGEBYLEX       = 266
	RedisParserZREMRANGEBYRANK      = 267
	RedisParserZREMRANGEBYSCORE     = 268
	RedisParserZREVRANGE            = 269
	RedisParserZREVRANGEBYLEX       = 270
	RedisParserZREVRANGEBYSCORE     = 271
	RedisParserZREVRANK             = 272
	RedisParserZSCAN                = 273
	RedisParserZSCORE               = 274
	RedisParserZUNION               = 275
	RedisParserZUNIONSTORE          = 276
	RedisParserXACK                 = 277
	RedisParserXADD                 = 278
	RedisParserXAUTOCLAIM           = 279
	RedisParserXCLAIM               = 280
	RedisParserXDEL                 = 281
	RedisParserXGROUP               = 282
	RedisParserCREATE               = 283
	RedisParserCREATECONSUMER       = 284
	RedisParserDELCONSUMER          = 285
	RedisParserDESTROY              = 286
	RedisParserSETID                = 287
	RedisParserXINFO                = 288
	RedisParserCONSUMERS            = 289
	RedisParserGROUPS               = 290
	RedisParserSTREAM               = 291
	RedisParserXLEN                 = 292
	RedisParserXPENDING             = 293
	RedisParserXRANGE               = 294
	RedisParserXREAD                = 295
	RedisParserXREADGROUP           = 296
	RedisParserXREVRANGE            = 297
	RedisParserXSETID               = 298
	RedisParserXTRIM                = 299
	RedisParserAPPEND               = 300
	RedisParserDECR                 = 301
	RedisParserDECRBY               = 302
	RedisParserGETDEL               = 303
	RedisParserGETEX                = 304
	RedisParserGETRANGE             = 305
	RedisParserGETSET               = 306
	RedisParserINCR                 = 307
	RedisParserINCRBY               = 308
	RedisParserINCRBYFLOAT          = 309
	RedisParserLCS                  = 310
	RedisParserMGET                 = 311
	RedisParserMSET                 = 312
	RedisParserMSETNX               = 313
	RedisParserPSETEX               = 314
	RedisParserSETEX                = 315
	RedisParserSETNX                = 316
	RedisParserSETRANGE             = 317
	RedisParserSTRLEN               = 318
	RedisParserSUBSTR               = 319
	RedisParserDISCARD              = 320
	RedisParserEXEC                 = 321
	RedisParserMULTI                = 322
	RedisParserUNWATCH              = 323
	RedisParserWATCH                = 324
	RedisParserDB                   = 325
	RedisParserSPACE                = 326
	RedisParserNL                   = 327
)

// RedisParser rules.
const (
	RedisParserRULE_command               = 0
	RedisParserRULE_bitmap                = 1
	RedisParserRULE_clusterManagement     = 2
	RedisParserRULE_clusterSub            = 3
	RedisParserRULE_connectionManagement  = 4
	RedisParserRULE_clientSub             = 5
	RedisParserRULE_generic               = 6
	RedisParserRULE_objectSub             = 7
	RedisParserRULE_geospatialIndices     = 8
	RedisParserRULE_hash                  = 9
	RedisParserRULE_hyperloglog           = 10
	RedisParserRULE_list                  = 11
	RedisParserRULE_pubSub                = 12
	RedisParserRULE_pubSubSub             = 13
	RedisParserRULE_scriptingAndFunctions = 14
	RedisParserRULE_functionSub           = 15
	RedisParserRULE_scriptSub             = 16
	RedisParserRULE_serverManagement      = 17
	RedisParserRULE_aclSub                = 18
	RedisParserRULE_commandSub            = 19
	RedisParserRULE_configSub             = 20
	RedisParserRULE_latencySub            = 21
	RedisParserRULE_memorySub             = 22
	RedisParserRULE_moduleSub             = 23
	RedisParserRULE_slowlogSub            = 24
	RedisParserRULE_set                   = 25
	RedisParserRULE_sortedSet             = 26
	RedisParserRULE_stream                = 27
	RedisParserRULE_xgroupSub             = 28
	RedisParserRULE_xinfoSub              = 29
	RedisParserRULE_stringCmd             = 30
	RedisParserRULE_transactions          = 31
)

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Bitmap() IBitmapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitmapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitmapContext)
}

func (s *CommandContext) ClusterManagement() IClusterManagementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClusterManagementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClusterManagementContext)
}

func (s *CommandContext) ConnectionManagement() IConnectionManagementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectionManagementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectionManagementContext)
}

func (s *CommandContext) Generic() IGenericContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericContext)
}

func (s *CommandContext) GeospatialIndices() IGeospatialIndicesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeospatialIndicesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeospatialIndicesContext)
}

func (s *CommandContext) Hash() IHashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHashContext)
}

func (s *CommandContext) Hyperloglog() IHyperloglogContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHyperloglogContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHyperloglogContext)
}

func (s *CommandContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *CommandContext) PubSub() IPubSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPubSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPubSubContext)
}

func (s *CommandContext) ScriptingAndFunctions() IScriptingAndFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScriptingAndFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScriptingAndFunctionsContext)
}

func (s *CommandContext) ServerManagement() IServerManagementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServerManagementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServerManagementContext)
}

func (s *CommandContext) Set() ISetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *CommandContext) SortedSet() ISortedSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortedSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortedSetContext)
}

func (s *CommandContext) Stream() IStreamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStreamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStreamContext)
}

func (s *CommandContext) StringCmd() IStringCmdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringCmdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringCmdContext)
}

func (s *CommandContext) Transactions() ITransactionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransactionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransactionsContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RedisParserRULE_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserBITCOUNT, RedisParserBITFIELD, RedisParserBITFIELD_RO, RedisParserBITOP, RedisParserBITPOS, RedisParserGETBIT, RedisParserSETBIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Bitmap()
		}

	case RedisParserASKING, RedisParserCLUSTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.ClusterManagement()
		}

	case RedisParserRESET, RedisParserAUTH, RedisParserCLIENT, RedisParserECHO, RedisParserHELLO, RedisParserPING, RedisParserQUIT, RedisParserSELECT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.ConnectionManagement()
		}

	case RedisParserCOPY, RedisParserDEL, RedisParserDUMP, RedisParserEXISTS, RedisParserEXPIRE, RedisParserEXPIREAT, RedisParserEXPIRETIME, RedisParserKEYS, RedisParserMIGRATE, RedisParserMOVE, RedisParserOBJECT, RedisParserPERSIST, RedisParserPEXPIRE, RedisParserPEXPIREAT, RedisParserPEXPIRETIME, RedisParserPTTL, RedisParserRANDOMKEY, RedisParserRENAME, RedisParserRENAMENX, RedisParserRESTORE, RedisParserSCAN, RedisParserSORT, RedisParserSORT_RO, RedisParserTOUCH, RedisParserTTL, RedisParserTYPE, RedisParserUNLINK, RedisParserWAIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Generic()
		}

	case RedisParserGEOADD, RedisParserGEODIST, RedisParserGEOHASH, RedisParserGEOPOS, RedisParserGEORADIUS, RedisParserGEORADIUSBYMEMBER, RedisParserGEORADIUSBYMEMBER_RO, RedisParserGEORADIUS_RO, RedisParserGEOSEARCH, RedisParserGEOSEARCHSTORE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.GeospatialIndices()
		}

	case RedisParserHDEL, RedisParserHEXISTS, RedisParserHGET, RedisParserHGETALL, RedisParserHINCRBY, RedisParserHINCRBYFLOAT, RedisParserHKEYS, RedisParserHLEN, RedisParserHMGET, RedisParserHMSET, RedisParserHRANDFIELD, RedisParserHSCAN, RedisParserHSET, RedisParserHSETNX, RedisParserHSTRLEN, RedisParserHVALS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.Hash()
		}

	case RedisParserPFADD, RedisParserPFCOUNT, RedisParserPFDEBUG, RedisParserPFMERGE, RedisParserPFSELFTEST:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.Hyperloglog()
		}

	case RedisParserBLMOVE, RedisParserBLMPOP, RedisParserBLPOP, RedisParserBRPOP, RedisParserBRPOPLPUSH, RedisParserLINDEX, RedisParserLINSERT, RedisParserLLEN, RedisParserLMOVE, RedisParserLMPOP, RedisParserLPOP, RedisParserLPOS, RedisParserLPUSH, RedisParserLPUSHX, RedisParserLRANGE, RedisParserLREM, RedisParserLSET, RedisParserLTRIM, RedisParserRPOP, RedisParserRPOPLPUSH, RedisParserRPUSH, RedisParserRPUSHX:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.List()
		}

	case RedisParserPSUBSCRIBE, RedisParserPUBLISH, RedisParserPUBSUB, RedisParserPUNSUBSCRIBE, RedisParserSPUBLISH, RedisParserSSUBSCRIBE, RedisParserSUBSCRIBE, RedisParserSUNSUBSCRIBE, RedisParserUNSUBSCRIBE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(72)
			p.PubSub()
		}

	case RedisParserEVAL, RedisParserEVALSHA, RedisParserEVALSHA_RO, RedisParserEVAL_RO, RedisParserFCALL, RedisParserFCALL_RO, RedisParserFUNCTION, RedisParserSCRIPT:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(73)
			p.ScriptingAndFunctions()
		}

	case RedisParserFAILOVER, RedisParserINFO, RedisParserDEBUG, RedisParserACL, RedisParserSAVE, RedisParserBGREWRITEAOF, RedisParserBGSAVE, RedisParserCOMMAND, RedisParserCONFIG, RedisParserDBSIZE, RedisParserFLUSHALL, RedisParserFLUSHDB, RedisParserLASTSAVE, RedisParserLATENCY, RedisParserLOLWUT, RedisParserMEMORY, RedisParserMODULE, RedisParserMONITOR, RedisParserPSYNC, RedisParserREPLCONF, RedisParserREPLICAOF, RedisParserRESTOREASKING, RedisParserROLE, RedisParserSHUTDOWN, RedisParserSLAVEOF, RedisParserSLOWLOG, RedisParserSWAPDB, RedisParserSYNC, RedisParserTIME:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(74)
			p.ServerManagement()
		}

	case RedisParserSADD, RedisParserSCARD, RedisParserSDIFF, RedisParserSDIFFSTORE, RedisParserSINTER, RedisParserSINTERCARD, RedisParserSINTERSTORE, RedisParserSISMEMBER, RedisParserSMEMBERS, RedisParserSMISMEMBER, RedisParserSMOVE, RedisParserSPOP, RedisParserSRANDMEMBER, RedisParserSREM, RedisParserSSCAN, RedisParserSUNION, RedisParserSUNIONSTORE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(75)
			p.Set()
		}

	case RedisParserBZMPOP, RedisParserBZPOPMAX, RedisParserBZPOPMIN, RedisParserZADD, RedisParserZCARD, RedisParserZCOUNT, RedisParserZDIFF, RedisParserZDIFFSTORE, RedisParserZINCRBY, RedisParserZINTER, RedisParserZINTERCARD, RedisParserZINTERSTORE, RedisParserZLEXCOUNT, RedisParserZMPOP, RedisParserZMSCORE, RedisParserZPOPMAX, RedisParserZPOPMIN, RedisParserZRANDMEMBER, RedisParserZRANGE, RedisParserZRANGEBYLEX, RedisParserZRANGEBYSCORE, RedisParserZRANGESTORE, RedisParserZRANK, RedisParserZREM, RedisParserZREMRANGEBYLEX, RedisParserZREMRANGEBYRANK, RedisParserZREMRANGEBYSCORE, RedisParserZREVRANGE, RedisParserZREVRANGEBYLEX, RedisParserZREVRANGEBYSCORE, RedisParserZREVRANK, RedisParserZSCAN, RedisParserZSCORE, RedisParserZUNION, RedisParserZUNIONSTORE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(76)
			p.SortedSet()
		}

	case RedisParserXACK, RedisParserXADD, RedisParserXAUTOCLAIM, RedisParserXCLAIM, RedisParserXDEL, RedisParserXGROUP, RedisParserXINFO, RedisParserXLEN, RedisParserXPENDING, RedisParserXRANGE, RedisParserXREAD, RedisParserXREADGROUP, RedisParserXREVRANGE, RedisParserXSETID, RedisParserXTRIM:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(77)
			p.Stream()
		}

	case RedisParserGET, RedisParserSET, RedisParserAPPEND, RedisParserDECR, RedisParserDECRBY, RedisParserGETDEL, RedisParserGETEX, RedisParserGETRANGE, RedisParserGETSET, RedisParserINCR, RedisParserINCRBY, RedisParserINCRBYFLOAT, RedisParserLCS, RedisParserMGET, RedisParserMSET, RedisParserMSETNX, RedisParserPSETEX, RedisParserSETEX, RedisParserSETNX, RedisParserSETRANGE, RedisParserSTRLEN, RedisParserSUBSTR:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(78)
			p.StringCmd()
		}

	case RedisParserDISCARD, RedisParserEXEC, RedisParserMULTI, RedisParserUNWATCH, RedisParserWATCH:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(79)
			p.Transactions()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBitmapContext is an interface to support dynamic dispatch.
type IBitmapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitmapContext differentiates from other interfaces.
	IsBitmapContext()
}

type BitmapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitmapContext() *BitmapContext {
	var p = new(BitmapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_bitmap
	return p
}

func (*BitmapContext) IsBitmapContext() {}

func NewBitmapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitmapContext {
	var p = new(BitmapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_bitmap

	return p
}

func (s *BitmapContext) GetParser() antlr.Parser { return s.parser }

func (s *BitmapContext) BITCOUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserBITCOUNT, 0)
}

func (s *BitmapContext) BITFIELD() antlr.TerminalNode {
	return s.GetToken(RedisParserBITFIELD, 0)
}

func (s *BitmapContext) BITFIELD_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserBITFIELD_RO, 0)
}

func (s *BitmapContext) BITOP() antlr.TerminalNode {
	return s.GetToken(RedisParserBITOP, 0)
}

func (s *BitmapContext) BITPOS() antlr.TerminalNode {
	return s.GetToken(RedisParserBITPOS, 0)
}

func (s *BitmapContext) GETBIT() antlr.TerminalNode {
	return s.GetToken(RedisParserGETBIT, 0)
}

func (s *BitmapContext) SETBIT() antlr.TerminalNode {
	return s.GetToken(RedisParserSETBIT, 0)
}

func (s *BitmapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitmapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitmapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitBitmap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Bitmap() (localctx IBitmapContext) {
	this := p
	_ = this

	localctx = NewBitmapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RedisParserRULE_bitmap)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RedisParserBITCOUNT)|(1<<RedisParserBITFIELD)|(1<<RedisParserBITFIELD_RO)|(1<<RedisParserBITOP)|(1<<RedisParserBITPOS)|(1<<RedisParserGETBIT)|(1<<RedisParserSETBIT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IClusterManagementContext is an interface to support dynamic dispatch.
type IClusterManagementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClusterManagementContext differentiates from other interfaces.
	IsClusterManagementContext()
}

type ClusterManagementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusterManagementContext() *ClusterManagementContext {
	var p = new(ClusterManagementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_clusterManagement
	return p
}

func (*ClusterManagementContext) IsClusterManagementContext() {}

func NewClusterManagementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusterManagementContext {
	var p = new(ClusterManagementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_clusterManagement

	return p
}

func (s *ClusterManagementContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusterManagementContext) ASKING() antlr.TerminalNode {
	return s.GetToken(RedisParserASKING, 0)
}

func (s *ClusterManagementContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(RedisParserCLUSTER, 0)
}

func (s *ClusterManagementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *ClusterManagementContext) ClusterSub() IClusterSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClusterSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClusterSubContext)
}

func (s *ClusterManagementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusterManagementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusterManagementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitClusterManagement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ClusterManagement() (localctx IClusterManagementContext) {
	this := p
	_ = this

	localctx = NewClusterManagementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RedisParserRULE_clusterManagement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserASKING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(RedisParserASKING)
		}

	case RedisParserCLUSTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(RedisParserCLUSTER)
		}
		{
			p.SetState(86)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(87)
			p.ClusterSub()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClusterSubContext is an interface to support dynamic dispatch.
type IClusterSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClusterSubContext differentiates from other interfaces.
	IsClusterSubContext()
}

type ClusterSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusterSubContext() *ClusterSubContext {
	var p = new(ClusterSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_clusterSub
	return p
}

func (*ClusterSubContext) IsClusterSubContext() {}

func NewClusterSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusterSubContext {
	var p = new(ClusterSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_clusterSub

	return p
}

func (s *ClusterSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusterSubContext) ADDSLOTS() antlr.TerminalNode {
	return s.GetToken(RedisParserADDSLOTS, 0)
}

func (s *ClusterSubContext) ADDSLOTSRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserADDSLOTSRANGE, 0)
}

func (s *ClusterSubContext) BUMPEPOCH() antlr.TerminalNode {
	return s.GetToken(RedisParserBUMPEPOCH, 0)
}

func (s *ClusterSubContext) COUNTFAILUREREPORTS() antlr.TerminalNode {
	return s.GetToken(RedisParserCOUNTFAILUREREPORTS, 0)
}

func (s *ClusterSubContext) COUNTKEYSINSLOT() antlr.TerminalNode {
	return s.GetToken(RedisParserCOUNTKEYSINSLOT, 0)
}

func (s *ClusterSubContext) DELSLOTS() antlr.TerminalNode {
	return s.GetToken(RedisParserDELSLOTS, 0)
}

func (s *ClusterSubContext) DELSLOTSRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserDELSLOTSRANGE, 0)
}

func (s *ClusterSubContext) FAILOVER() antlr.TerminalNode {
	return s.GetToken(RedisParserFAILOVER, 0)
}

func (s *ClusterSubContext) FLUSHSLOTS() antlr.TerminalNode {
	return s.GetToken(RedisParserFLUSHSLOTS, 0)
}

func (s *ClusterSubContext) FORGET() antlr.TerminalNode {
	return s.GetToken(RedisParserFORGET, 0)
}

func (s *ClusterSubContext) GETKEYSINSLOT() antlr.TerminalNode {
	return s.GetToken(RedisParserGETKEYSINSLOT, 0)
}

func (s *ClusterSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ClusterSubContext) INFO() antlr.TerminalNode {
	return s.GetToken(RedisParserINFO, 0)
}

func (s *ClusterSubContext) KEYSLOT() antlr.TerminalNode {
	return s.GetToken(RedisParserKEYSLOT, 0)
}

func (s *ClusterSubContext) LINKS() antlr.TerminalNode {
	return s.GetToken(RedisParserLINKS, 0)
}

func (s *ClusterSubContext) MEET() antlr.TerminalNode {
	return s.GetToken(RedisParserMEET, 0)
}

func (s *ClusterSubContext) MYID() antlr.TerminalNode {
	return s.GetToken(RedisParserMYID, 0)
}

func (s *ClusterSubContext) NODES() antlr.TerminalNode {
	return s.GetToken(RedisParserNODES, 0)
}

func (s *ClusterSubContext) REPLICAS() antlr.TerminalNode {
	return s.GetToken(RedisParserREPLICAS, 0)
}

func (s *ClusterSubContext) REPLICATE() antlr.TerminalNode {
	return s.GetToken(RedisParserREPLICATE, 0)
}

func (s *ClusterSubContext) RESET() antlr.TerminalNode {
	return s.GetToken(RedisParserRESET, 0)
}

func (s *ClusterSubContext) SAVECONFIG() antlr.TerminalNode {
	return s.GetToken(RedisParserSAVECONFIG, 0)
}

func (s *ClusterSubContext) SETCONFIGEPOCH() antlr.TerminalNode {
	return s.GetToken(RedisParserSETCONFIGEPOCH, 0)
}

func (s *ClusterSubContext) SETSLOT() antlr.TerminalNode {
	return s.GetToken(RedisParserSETSLOT, 0)
}

func (s *ClusterSubContext) SHARDS() antlr.TerminalNode {
	return s.GetToken(RedisParserSHARDS, 0)
}

func (s *ClusterSubContext) SLAVES() antlr.TerminalNode {
	return s.GetToken(RedisParserSLAVES, 0)
}

func (s *ClusterSubContext) SLOTS() antlr.TerminalNode {
	return s.GetToken(RedisParserSLOTS, 0)
}

func (s *ClusterSubContext) READONLY() antlr.TerminalNode {
	return s.GetToken(RedisParserREADONLY, 0)
}

func (s *ClusterSubContext) READWRITE() antlr.TerminalNode {
	return s.GetToken(RedisParserREADWRITE, 0)
}

func (s *ClusterSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusterSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusterSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitClusterSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ClusterSub() (localctx IClusterSubContext) {
	this := p
	_ = this

	localctx = NewClusterSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RedisParserRULE_clusterSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RedisParserADDSLOTS)|(1<<RedisParserADDSLOTSRANGE)|(1<<RedisParserBUMPEPOCH)|(1<<RedisParserCOUNTFAILUREREPORTS)|(1<<RedisParserCOUNTKEYSINSLOT)|(1<<RedisParserDELSLOTS)|(1<<RedisParserDELSLOTSRANGE)|(1<<RedisParserFAILOVER)|(1<<RedisParserFLUSHSLOTS)|(1<<RedisParserFORGET)|(1<<RedisParserGETKEYSINSLOT)|(1<<RedisParserKEYSLOT)|(1<<RedisParserLINKS)|(1<<RedisParserMEET)|(1<<RedisParserMYID)|(1<<RedisParserNODES)|(1<<RedisParserREPLICAS)|(1<<RedisParserREPLICATE)|(1<<RedisParserRESET)|(1<<RedisParserSAVECONFIG)|(1<<RedisParserSETCONFIGEPOCH)|(1<<RedisParserSETSLOT))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(RedisParserSHARDS-32))|(1<<(RedisParserSLAVES-32))|(1<<(RedisParserSLOTS-32))|(1<<(RedisParserREADONLY-32))|(1<<(RedisParserREADWRITE-32))|(1<<(RedisParserINFO-32)))) != 0) || _la == RedisParserHELP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConnectionManagementContext is an interface to support dynamic dispatch.
type IConnectionManagementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectionManagementContext differentiates from other interfaces.
	IsConnectionManagementContext()
}

type ConnectionManagementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectionManagementContext() *ConnectionManagementContext {
	var p = new(ConnectionManagementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_connectionManagement
	return p
}

func (*ConnectionManagementContext) IsConnectionManagementContext() {}

func NewConnectionManagementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectionManagementContext {
	var p = new(ConnectionManagementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_connectionManagement

	return p
}

func (s *ConnectionManagementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectionManagementContext) AUTH() antlr.TerminalNode {
	return s.GetToken(RedisParserAUTH, 0)
}

func (s *ConnectionManagementContext) CLIENT() antlr.TerminalNode {
	return s.GetToken(RedisParserCLIENT, 0)
}

func (s *ConnectionManagementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *ConnectionManagementContext) ClientSub() IClientSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClientSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClientSubContext)
}

func (s *ConnectionManagementContext) ECHO() antlr.TerminalNode {
	return s.GetToken(RedisParserECHO, 0)
}

func (s *ConnectionManagementContext) HELLO() antlr.TerminalNode {
	return s.GetToken(RedisParserHELLO, 0)
}

func (s *ConnectionManagementContext) PING() antlr.TerminalNode {
	return s.GetToken(RedisParserPING, 0)
}

func (s *ConnectionManagementContext) QUIT() antlr.TerminalNode {
	return s.GetToken(RedisParserQUIT, 0)
}

func (s *ConnectionManagementContext) RESET() antlr.TerminalNode {
	return s.GetToken(RedisParserRESET, 0)
}

func (s *ConnectionManagementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(RedisParserSELECT, 0)
}

func (s *ConnectionManagementContext) DB() antlr.TerminalNode {
	return s.GetToken(RedisParserDB, 0)
}

func (s *ConnectionManagementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionManagementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectionManagementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitConnectionManagement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ConnectionManagement() (localctx IConnectionManagementContext) {
	this := p
	_ = this

	localctx = NewConnectionManagementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RedisParserRULE_connectionManagement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserAUTH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(RedisParserAUTH)
		}

	case RedisParserCLIENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(RedisParserCLIENT)
		}
		{
			p.SetState(94)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(95)
			p.ClientSub()
		}

	case RedisParserECHO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(RedisParserECHO)
		}

	case RedisParserHELLO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.Match(RedisParserHELLO)
		}

	case RedisParserPING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(RedisParserPING)
		}

	case RedisParserQUIT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.Match(RedisParserQUIT)
		}

	case RedisParserRESET:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(100)
			p.Match(RedisParserRESET)
		}

	case RedisParserSELECT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(101)
			p.Match(RedisParserSELECT)
		}
		{
			p.SetState(102)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(103)
			p.Match(RedisParserDB)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClientSubContext is an interface to support dynamic dispatch.
type IClientSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClientSubContext differentiates from other interfaces.
	IsClientSubContext()
}

type ClientSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClientSubContext() *ClientSubContext {
	var p = new(ClientSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_clientSub
	return p
}

func (*ClientSubContext) IsClientSubContext() {}

func NewClientSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClientSubContext {
	var p = new(ClientSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_clientSub

	return p
}

func (s *ClientSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ClientSubContext) CACHING() antlr.TerminalNode {
	return s.GetToken(RedisParserCACHING, 0)
}

func (s *ClientSubContext) GETNAME() antlr.TerminalNode {
	return s.GetToken(RedisParserGETNAME, 0)
}

func (s *ClientSubContext) GETREDIR() antlr.TerminalNode {
	return s.GetToken(RedisParserGETREDIR, 0)
}

func (s *ClientSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ClientSubContext) ID() antlr.TerminalNode {
	return s.GetToken(RedisParserID, 0)
}

func (s *ClientSubContext) INFO() antlr.TerminalNode {
	return s.GetToken(RedisParserINFO, 0)
}

func (s *ClientSubContext) KILL() antlr.TerminalNode {
	return s.GetToken(RedisParserKILL, 0)
}

func (s *ClientSubContext) LIST() antlr.TerminalNode {
	return s.GetToken(RedisParserLIST, 0)
}

func (s *ClientSubContext) NOEVICT() antlr.TerminalNode {
	return s.GetToken(RedisParserNOEVICT, 0)
}

func (s *ClientSubContext) PAUSE() antlr.TerminalNode {
	return s.GetToken(RedisParserPAUSE, 0)
}

func (s *ClientSubContext) REPLY() antlr.TerminalNode {
	return s.GetToken(RedisParserREPLY, 0)
}

func (s *ClientSubContext) SETNAME() antlr.TerminalNode {
	return s.GetToken(RedisParserSETNAME, 0)
}

func (s *ClientSubContext) TRACKING() antlr.TerminalNode {
	return s.GetToken(RedisParserTRACKING, 0)
}

func (s *ClientSubContext) TRACKINGINFO() antlr.TerminalNode {
	return s.GetToken(RedisParserTRACKINGINFO, 0)
}

func (s *ClientSubContext) UNBLOCK() antlr.TerminalNode {
	return s.GetToken(RedisParserUNBLOCK, 0)
}

func (s *ClientSubContext) UNPAUSE() antlr.TerminalNode {
	return s.GetToken(RedisParserUNPAUSE, 0)
}

func (s *ClientSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClientSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClientSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitClientSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ClientSub() (localctx IClientSubContext) {
	this := p
	_ = this

	localctx = NewClientSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RedisParserRULE_clientSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(RedisParserCACHING-39))|(1<<(RedisParserGETNAME-39))|(1<<(RedisParserGETREDIR-39))|(1<<(RedisParserID-39))|(1<<(RedisParserINFO-39))|(1<<(RedisParserKILL-39))|(1<<(RedisParserNOEVICT-39))|(1<<(RedisParserPAUSE-39))|(1<<(RedisParserREPLY-39))|(1<<(RedisParserSETNAME-39))|(1<<(RedisParserTRACKING-39))|(1<<(RedisParserTRACKINGINFO-39))|(1<<(RedisParserUNBLOCK-39))|(1<<(RedisParserUNPAUSE-39)))) != 0) || _la == RedisParserLIST || _la == RedisParserHELP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGenericContext is an interface to support dynamic dispatch.
type IGenericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericContext differentiates from other interfaces.
	IsGenericContext()
}

type GenericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericContext() *GenericContext {
	var p = new(GenericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_generic
	return p
}

func (*GenericContext) IsGenericContext() {}

func NewGenericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericContext {
	var p = new(GenericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_generic

	return p
}

func (s *GenericContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericContext) COPY() antlr.TerminalNode {
	return s.GetToken(RedisParserCOPY, 0)
}

func (s *GenericContext) DEL() antlr.TerminalNode {
	return s.GetToken(RedisParserDEL, 0)
}

func (s *GenericContext) DUMP() antlr.TerminalNode {
	return s.GetToken(RedisParserDUMP, 0)
}

func (s *GenericContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(RedisParserEXISTS, 0)
}

func (s *GenericContext) EXPIRE() antlr.TerminalNode {
	return s.GetToken(RedisParserEXPIRE, 0)
}

func (s *GenericContext) EXPIREAT() antlr.TerminalNode {
	return s.GetToken(RedisParserEXPIREAT, 0)
}

func (s *GenericContext) EXPIRETIME() antlr.TerminalNode {
	return s.GetToken(RedisParserEXPIRETIME, 0)
}

func (s *GenericContext) KEYS() antlr.TerminalNode {
	return s.GetToken(RedisParserKEYS, 0)
}

func (s *GenericContext) MIGRATE() antlr.TerminalNode {
	return s.GetToken(RedisParserMIGRATE, 0)
}

func (s *GenericContext) MOVE() antlr.TerminalNode {
	return s.GetToken(RedisParserMOVE, 0)
}

func (s *GenericContext) OBJECT() antlr.TerminalNode {
	return s.GetToken(RedisParserOBJECT, 0)
}

func (s *GenericContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *GenericContext) ObjectSub() IObjectSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectSubContext)
}

func (s *GenericContext) PERSIST() antlr.TerminalNode {
	return s.GetToken(RedisParserPERSIST, 0)
}

func (s *GenericContext) PEXPIRE() antlr.TerminalNode {
	return s.GetToken(RedisParserPEXPIRE, 0)
}

func (s *GenericContext) PEXPIREAT() antlr.TerminalNode {
	return s.GetToken(RedisParserPEXPIREAT, 0)
}

func (s *GenericContext) PEXPIRETIME() antlr.TerminalNode {
	return s.GetToken(RedisParserPEXPIRETIME, 0)
}

func (s *GenericContext) PTTL() antlr.TerminalNode {
	return s.GetToken(RedisParserPTTL, 0)
}

func (s *GenericContext) RANDOMKEY() antlr.TerminalNode {
	return s.GetToken(RedisParserRANDOMKEY, 0)
}

func (s *GenericContext) RENAME() antlr.TerminalNode {
	return s.GetToken(RedisParserRENAME, 0)
}

func (s *GenericContext) RENAMENX() antlr.TerminalNode {
	return s.GetToken(RedisParserRENAMENX, 0)
}

func (s *GenericContext) RESTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserRESTORE, 0)
}

func (s *GenericContext) SCAN() antlr.TerminalNode {
	return s.GetToken(RedisParserSCAN, 0)
}

func (s *GenericContext) SORT() antlr.TerminalNode {
	return s.GetToken(RedisParserSORT, 0)
}

func (s *GenericContext) SORT_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserSORT_RO, 0)
}

func (s *GenericContext) TOUCH() antlr.TerminalNode {
	return s.GetToken(RedisParserTOUCH, 0)
}

func (s *GenericContext) TTL() antlr.TerminalNode {
	return s.GetToken(RedisParserTTL, 0)
}

func (s *GenericContext) TYPE() antlr.TerminalNode {
	return s.GetToken(RedisParserTYPE, 0)
}

func (s *GenericContext) UNLINK() antlr.TerminalNode {
	return s.GetToken(RedisParserUNLINK, 0)
}

func (s *GenericContext) WAIT() antlr.TerminalNode {
	return s.GetToken(RedisParserWAIT, 0)
}

func (s *GenericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitGeneric(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Generic() (localctx IGenericContext) {
	this := p
	_ = this

	localctx = NewGenericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RedisParserRULE_generic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserCOPY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(RedisParserCOPY)
		}

	case RedisParserDEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(RedisParserDEL)
		}

	case RedisParserDUMP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(RedisParserDUMP)
		}

	case RedisParserEXISTS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Match(RedisParserEXISTS)
		}

	case RedisParserEXPIRE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Match(RedisParserEXPIRE)
		}

	case RedisParserEXPIREAT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Match(RedisParserEXPIREAT)
		}

	case RedisParserEXPIRETIME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(114)
			p.Match(RedisParserEXPIRETIME)
		}

	case RedisParserKEYS:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(115)
			p.Match(RedisParserKEYS)
		}

	case RedisParserMIGRATE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(116)
			p.Match(RedisParserMIGRATE)
		}

	case RedisParserMOVE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(117)
			p.Match(RedisParserMOVE)
		}

	case RedisParserOBJECT:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(118)
			p.Match(RedisParserOBJECT)
		}
		{
			p.SetState(119)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(120)
			p.ObjectSub()
		}

	case RedisParserPERSIST:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(121)
			p.Match(RedisParserPERSIST)
		}

	case RedisParserPEXPIRE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(122)
			p.Match(RedisParserPEXPIRE)
		}

	case RedisParserPEXPIREAT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(123)
			p.Match(RedisParserPEXPIREAT)
		}

	case RedisParserPEXPIRETIME:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(124)
			p.Match(RedisParserPEXPIRETIME)
		}

	case RedisParserPTTL:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(125)
			p.Match(RedisParserPTTL)
		}

	case RedisParserRANDOMKEY:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(126)
			p.Match(RedisParserRANDOMKEY)
		}

	case RedisParserRENAME:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(127)
			p.Match(RedisParserRENAME)
		}

	case RedisParserRENAMENX:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(128)
			p.Match(RedisParserRENAMENX)
		}

	case RedisParserRESTORE:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(129)
			p.Match(RedisParserRESTORE)
		}

	case RedisParserSCAN:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(130)
			p.Match(RedisParserSCAN)
		}

	case RedisParserSORT:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(131)
			p.Match(RedisParserSORT)
		}

	case RedisParserSORT_RO:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(132)
			p.Match(RedisParserSORT_RO)
		}

	case RedisParserTOUCH:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(133)
			p.Match(RedisParserTOUCH)
		}

	case RedisParserTTL:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(134)
			p.Match(RedisParserTTL)
		}

	case RedisParserTYPE:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(135)
			p.Match(RedisParserTYPE)
		}

	case RedisParserUNLINK:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(136)
			p.Match(RedisParserUNLINK)
		}

	case RedisParserWAIT:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(137)
			p.Match(RedisParserWAIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectSubContext is an interface to support dynamic dispatch.
type IObjectSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectSubContext differentiates from other interfaces.
	IsObjectSubContext()
}

type ObjectSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectSubContext() *ObjectSubContext {
	var p = new(ObjectSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_objectSub
	return p
}

func (*ObjectSubContext) IsObjectSubContext() {}

func NewObjectSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectSubContext {
	var p = new(ObjectSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_objectSub

	return p
}

func (s *ObjectSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectSubContext) ENCODING() antlr.TerminalNode {
	return s.GetToken(RedisParserENCODING, 0)
}

func (s *ObjectSubContext) FREQ() antlr.TerminalNode {
	return s.GetToken(RedisParserFREQ, 0)
}

func (s *ObjectSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ObjectSubContext) IDLETIME() antlr.TerminalNode {
	return s.GetToken(RedisParserIDLETIME, 0)
}

func (s *ObjectSubContext) REFCOUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserREFCOUNT, 0)
}

func (s *ObjectSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitObjectSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ObjectSub() (localctx IObjectSubContext) {
	this := p
	_ = this

	localctx = NewObjectSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RedisParserRULE_objectSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(RedisParserENCODING-69))|(1<<(RedisParserFREQ-69))|(1<<(RedisParserIDLETIME-69))|(1<<(RedisParserREFCOUNT-69)))) != 0) || _la == RedisParserHELP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGeospatialIndicesContext is an interface to support dynamic dispatch.
type IGeospatialIndicesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeospatialIndicesContext differentiates from other interfaces.
	IsGeospatialIndicesContext()
}

type GeospatialIndicesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeospatialIndicesContext() *GeospatialIndicesContext {
	var p = new(GeospatialIndicesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_geospatialIndices
	return p
}

func (*GeospatialIndicesContext) IsGeospatialIndicesContext() {}

func NewGeospatialIndicesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeospatialIndicesContext {
	var p = new(GeospatialIndicesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_geospatialIndices

	return p
}

func (s *GeospatialIndicesContext) GetParser() antlr.Parser { return s.parser }

func (s *GeospatialIndicesContext) GEOADD() antlr.TerminalNode {
	return s.GetToken(RedisParserGEOADD, 0)
}

func (s *GeospatialIndicesContext) GEODIST() antlr.TerminalNode {
	return s.GetToken(RedisParserGEODIST, 0)
}

func (s *GeospatialIndicesContext) GEOHASH() antlr.TerminalNode {
	return s.GetToken(RedisParserGEOHASH, 0)
}

func (s *GeospatialIndicesContext) GEOPOS() antlr.TerminalNode {
	return s.GetToken(RedisParserGEOPOS, 0)
}

func (s *GeospatialIndicesContext) GEORADIUS() antlr.TerminalNode {
	return s.GetToken(RedisParserGEORADIUS, 0)
}

func (s *GeospatialIndicesContext) GEORADIUSBYMEMBER() antlr.TerminalNode {
	return s.GetToken(RedisParserGEORADIUSBYMEMBER, 0)
}

func (s *GeospatialIndicesContext) GEORADIUSBYMEMBER_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserGEORADIUSBYMEMBER_RO, 0)
}

func (s *GeospatialIndicesContext) GEORADIUS_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserGEORADIUS_RO, 0)
}

func (s *GeospatialIndicesContext) GEOSEARCH() antlr.TerminalNode {
	return s.GetToken(RedisParserGEOSEARCH, 0)
}

func (s *GeospatialIndicesContext) GEOSEARCHSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserGEOSEARCHSTORE, 0)
}

func (s *GeospatialIndicesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeospatialIndicesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeospatialIndicesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitGeospatialIndices(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) GeospatialIndices() (localctx IGeospatialIndicesContext) {
	this := p
	_ = this

	localctx = NewGeospatialIndicesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RedisParserRULE_geospatialIndices)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(RedisParserGEOADD-90))|(1<<(RedisParserGEODIST-90))|(1<<(RedisParserGEOHASH-90))|(1<<(RedisParserGEOPOS-90))|(1<<(RedisParserGEORADIUS-90))|(1<<(RedisParserGEORADIUSBYMEMBER-90))|(1<<(RedisParserGEORADIUSBYMEMBER_RO-90))|(1<<(RedisParserGEORADIUS_RO-90))|(1<<(RedisParserGEOSEARCH-90))|(1<<(RedisParserGEOSEARCHSTORE-90)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHashContext is an interface to support dynamic dispatch.
type IHashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHashContext differentiates from other interfaces.
	IsHashContext()
}

type HashContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashContext() *HashContext {
	var p = new(HashContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_hash
	return p
}

func (*HashContext) IsHashContext() {}

func NewHashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashContext {
	var p = new(HashContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_hash

	return p
}

func (s *HashContext) GetParser() antlr.Parser { return s.parser }

func (s *HashContext) HDEL() antlr.TerminalNode {
	return s.GetToken(RedisParserHDEL, 0)
}

func (s *HashContext) HEXISTS() antlr.TerminalNode {
	return s.GetToken(RedisParserHEXISTS, 0)
}

func (s *HashContext) HGET() antlr.TerminalNode {
	return s.GetToken(RedisParserHGET, 0)
}

func (s *HashContext) HGETALL() antlr.TerminalNode {
	return s.GetToken(RedisParserHGETALL, 0)
}

func (s *HashContext) HINCRBY() antlr.TerminalNode {
	return s.GetToken(RedisParserHINCRBY, 0)
}

func (s *HashContext) HINCRBYFLOAT() antlr.TerminalNode {
	return s.GetToken(RedisParserHINCRBYFLOAT, 0)
}

func (s *HashContext) HKEYS() antlr.TerminalNode {
	return s.GetToken(RedisParserHKEYS, 0)
}

func (s *HashContext) HLEN() antlr.TerminalNode {
	return s.GetToken(RedisParserHLEN, 0)
}

func (s *HashContext) HMGET() antlr.TerminalNode {
	return s.GetToken(RedisParserHMGET, 0)
}

func (s *HashContext) HMSET() antlr.TerminalNode {
	return s.GetToken(RedisParserHMSET, 0)
}

func (s *HashContext) HRANDFIELD() antlr.TerminalNode {
	return s.GetToken(RedisParserHRANDFIELD, 0)
}

func (s *HashContext) HSCAN() antlr.TerminalNode {
	return s.GetToken(RedisParserHSCAN, 0)
}

func (s *HashContext) HSET() antlr.TerminalNode {
	return s.GetToken(RedisParserHSET, 0)
}

func (s *HashContext) HSETNX() antlr.TerminalNode {
	return s.GetToken(RedisParserHSETNX, 0)
}

func (s *HashContext) HSTRLEN() antlr.TerminalNode {
	return s.GetToken(RedisParserHSTRLEN, 0)
}

func (s *HashContext) HVALS() antlr.TerminalNode {
	return s.GetToken(RedisParserHVALS, 0)
}

func (s *HashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HashContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitHash(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Hash() (localctx IHashContext) {
	this := p
	_ = this

	localctx = NewHashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RedisParserRULE_hash)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-100)&-(0x1f+1)) == 0 && ((1<<uint((_la-100)))&((1<<(RedisParserHDEL-100))|(1<<(RedisParserHEXISTS-100))|(1<<(RedisParserHGET-100))|(1<<(RedisParserHGETALL-100))|(1<<(RedisParserHINCRBY-100))|(1<<(RedisParserHINCRBYFLOAT-100))|(1<<(RedisParserHKEYS-100))|(1<<(RedisParserHLEN-100))|(1<<(RedisParserHMGET-100))|(1<<(RedisParserHMSET-100))|(1<<(RedisParserHRANDFIELD-100))|(1<<(RedisParserHSCAN-100))|(1<<(RedisParserHSET-100))|(1<<(RedisParserHSETNX-100))|(1<<(RedisParserHSTRLEN-100))|(1<<(RedisParserHVALS-100)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHyperloglogContext is an interface to support dynamic dispatch.
type IHyperloglogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHyperloglogContext differentiates from other interfaces.
	IsHyperloglogContext()
}

type HyperloglogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHyperloglogContext() *HyperloglogContext {
	var p = new(HyperloglogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_hyperloglog
	return p
}

func (*HyperloglogContext) IsHyperloglogContext() {}

func NewHyperloglogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HyperloglogContext {
	var p = new(HyperloglogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_hyperloglog

	return p
}

func (s *HyperloglogContext) GetParser() antlr.Parser { return s.parser }

func (s *HyperloglogContext) PFADD() antlr.TerminalNode {
	return s.GetToken(RedisParserPFADD, 0)
}

func (s *HyperloglogContext) PFCOUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserPFCOUNT, 0)
}

func (s *HyperloglogContext) PFDEBUG() antlr.TerminalNode {
	return s.GetToken(RedisParserPFDEBUG, 0)
}

func (s *HyperloglogContext) PFMERGE() antlr.TerminalNode {
	return s.GetToken(RedisParserPFMERGE, 0)
}

func (s *HyperloglogContext) PFSELFTEST() antlr.TerminalNode {
	return s.GetToken(RedisParserPFSELFTEST, 0)
}

func (s *HyperloglogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HyperloglogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HyperloglogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitHyperloglog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Hyperloglog() (localctx IHyperloglogContext) {
	this := p
	_ = this

	localctx = NewHyperloglogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RedisParserRULE_hyperloglog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-116)&-(0x1f+1)) == 0 && ((1<<uint((_la-116)))&((1<<(RedisParserPFADD-116))|(1<<(RedisParserPFCOUNT-116))|(1<<(RedisParserPFDEBUG-116))|(1<<(RedisParserPFMERGE-116))|(1<<(RedisParserPFSELFTEST-116)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) BLMOVE() antlr.TerminalNode {
	return s.GetToken(RedisParserBLMOVE, 0)
}

func (s *ListContext) BLMPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserBLMPOP, 0)
}

func (s *ListContext) BLPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserBLPOP, 0)
}

func (s *ListContext) BRPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserBRPOP, 0)
}

func (s *ListContext) BRPOPLPUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserBRPOPLPUSH, 0)
}

func (s *ListContext) LINDEX() antlr.TerminalNode {
	return s.GetToken(RedisParserLINDEX, 0)
}

func (s *ListContext) LINSERT() antlr.TerminalNode {
	return s.GetToken(RedisParserLINSERT, 0)
}

func (s *ListContext) LLEN() antlr.TerminalNode {
	return s.GetToken(RedisParserLLEN, 0)
}

func (s *ListContext) LMOVE() antlr.TerminalNode {
	return s.GetToken(RedisParserLMOVE, 0)
}

func (s *ListContext) LMPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserLMPOP, 0)
}

func (s *ListContext) LPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserLPOP, 0)
}

func (s *ListContext) LPOS() antlr.TerminalNode {
	return s.GetToken(RedisParserLPOS, 0)
}

func (s *ListContext) LPUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserLPUSH, 0)
}

func (s *ListContext) LPUSHX() antlr.TerminalNode {
	return s.GetToken(RedisParserLPUSHX, 0)
}

func (s *ListContext) LRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserLRANGE, 0)
}

func (s *ListContext) LREM() antlr.TerminalNode {
	return s.GetToken(RedisParserLREM, 0)
}

func (s *ListContext) LSET() antlr.TerminalNode {
	return s.GetToken(RedisParserLSET, 0)
}

func (s *ListContext) LTRIM() antlr.TerminalNode {
	return s.GetToken(RedisParserLTRIM, 0)
}

func (s *ListContext) RPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserRPOP, 0)
}

func (s *ListContext) RPOPLPUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserRPOPLPUSH, 0)
}

func (s *ListContext) RPUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserRPUSH, 0)
}

func (s *ListContext) RPUSHX() antlr.TerminalNode {
	return s.GetToken(RedisParserRPUSHX, 0)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RedisParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-121)&-(0x1f+1)) == 0 && ((1<<uint((_la-121)))&((1<<(RedisParserBLMOVE-121))|(1<<(RedisParserBLMPOP-121))|(1<<(RedisParserBLPOP-121))|(1<<(RedisParserBRPOP-121))|(1<<(RedisParserBRPOPLPUSH-121))|(1<<(RedisParserLINDEX-121))|(1<<(RedisParserLINSERT-121))|(1<<(RedisParserLLEN-121))|(1<<(RedisParserLMOVE-121))|(1<<(RedisParserLMPOP-121))|(1<<(RedisParserLPOP-121))|(1<<(RedisParserLPOS-121))|(1<<(RedisParserLPUSH-121))|(1<<(RedisParserLPUSHX-121))|(1<<(RedisParserLRANGE-121))|(1<<(RedisParserLREM-121))|(1<<(RedisParserLSET-121))|(1<<(RedisParserLTRIM-121))|(1<<(RedisParserRPOP-121))|(1<<(RedisParserRPOPLPUSH-121))|(1<<(RedisParserRPUSH-121))|(1<<(RedisParserRPUSHX-121)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPubSubContext is an interface to support dynamic dispatch.
type IPubSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPubSubContext differentiates from other interfaces.
	IsPubSubContext()
}

type PubSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPubSubContext() *PubSubContext {
	var p = new(PubSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_pubSub
	return p
}

func (*PubSubContext) IsPubSubContext() {}

func NewPubSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PubSubContext {
	var p = new(PubSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_pubSub

	return p
}

func (s *PubSubContext) GetParser() antlr.Parser { return s.parser }

func (s *PubSubContext) PSUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserPSUBSCRIBE, 0)
}

func (s *PubSubContext) PUBLISH() antlr.TerminalNode {
	return s.GetToken(RedisParserPUBLISH, 0)
}

func (s *PubSubContext) PUBSUB() antlr.TerminalNode {
	return s.GetToken(RedisParserPUBSUB, 0)
}

func (s *PubSubContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *PubSubContext) PubSubSub() IPubSubSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPubSubSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPubSubSubContext)
}

func (s *PubSubContext) PUNSUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserPUNSUBSCRIBE, 0)
}

func (s *PubSubContext) SPUBLISH() antlr.TerminalNode {
	return s.GetToken(RedisParserSPUBLISH, 0)
}

func (s *PubSubContext) SSUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserSSUBSCRIBE, 0)
}

func (s *PubSubContext) SUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserSUBSCRIBE, 0)
}

func (s *PubSubContext) SUNSUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserSUNSUBSCRIBE, 0)
}

func (s *PubSubContext) UNSUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(RedisParserUNSUBSCRIBE, 0)
}

func (s *PubSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PubSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PubSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitPubSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) PubSub() (localctx IPubSubContext) {
	this := p
	_ = this

	localctx = NewPubSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RedisParserRULE_pubSub)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserPSUBSCRIBE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(RedisParserPSUBSCRIBE)
		}

	case RedisParserPUBLISH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(RedisParserPUBLISH)
		}

	case RedisParserPUBSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(RedisParserPUBSUB)
		}
		{
			p.SetState(153)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(154)
			p.PubSubSub()
		}

	case RedisParserPUNSUBSCRIBE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
			p.Match(RedisParserPUNSUBSCRIBE)
		}

	case RedisParserSPUBLISH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(156)
			p.Match(RedisParserSPUBLISH)
		}

	case RedisParserSSUBSCRIBE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(157)
			p.Match(RedisParserSSUBSCRIBE)
		}

	case RedisParserSUBSCRIBE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(158)
			p.Match(RedisParserSUBSCRIBE)
		}

	case RedisParserSUNSUBSCRIBE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(159)
			p.Match(RedisParserSUNSUBSCRIBE)
		}

	case RedisParserUNSUBSCRIBE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(160)
			p.Match(RedisParserUNSUBSCRIBE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPubSubSubContext is an interface to support dynamic dispatch.
type IPubSubSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPubSubSubContext differentiates from other interfaces.
	IsPubSubSubContext()
}

type PubSubSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPubSubSubContext() *PubSubSubContext {
	var p = new(PubSubSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_pubSubSub
	return p
}

func (*PubSubSubContext) IsPubSubSubContext() {}

func NewPubSubSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PubSubSubContext {
	var p = new(PubSubSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_pubSubSub

	return p
}

func (s *PubSubSubContext) GetParser() antlr.Parser { return s.parser }

func (s *PubSubSubContext) CHANNELS() antlr.TerminalNode {
	return s.GetToken(RedisParserCHANNELS, 0)
}

func (s *PubSubSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *PubSubSubContext) NUMPAT() antlr.TerminalNode {
	return s.GetToken(RedisParserNUMPAT, 0)
}

func (s *PubSubSubContext) NUMSUB() antlr.TerminalNode {
	return s.GetToken(RedisParserNUMSUB, 0)
}

func (s *PubSubSubContext) SHARDCHANNELS() antlr.TerminalNode {
	return s.GetToken(RedisParserSHARDCHANNELS, 0)
}

func (s *PubSubSubContext) SHARDNUMSUB() antlr.TerminalNode {
	return s.GetToken(RedisParserSHARDNUMSUB, 0)
}

func (s *PubSubSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PubSubSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PubSubSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitPubSubSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) PubSubSub() (localctx IPubSubSubContext) {
	this := p
	_ = this

	localctx = NewPubSubSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RedisParserRULE_pubSubSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-146)&-(0x1f+1)) == 0 && ((1<<uint((_la-146)))&((1<<(RedisParserCHANNELS-146))|(1<<(RedisParserNUMPAT-146))|(1<<(RedisParserNUMSUB-146))|(1<<(RedisParserSHARDCHANNELS-146))|(1<<(RedisParserSHARDNUMSUB-146))|(1<<(RedisParserHELP-146)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IScriptingAndFunctionsContext is an interface to support dynamic dispatch.
type IScriptingAndFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptingAndFunctionsContext differentiates from other interfaces.
	IsScriptingAndFunctionsContext()
}

type ScriptingAndFunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptingAndFunctionsContext() *ScriptingAndFunctionsContext {
	var p = new(ScriptingAndFunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_scriptingAndFunctions
	return p
}

func (*ScriptingAndFunctionsContext) IsScriptingAndFunctionsContext() {}

func NewScriptingAndFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptingAndFunctionsContext {
	var p = new(ScriptingAndFunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_scriptingAndFunctions

	return p
}

func (s *ScriptingAndFunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptingAndFunctionsContext) EVAL() antlr.TerminalNode {
	return s.GetToken(RedisParserEVAL, 0)
}

func (s *ScriptingAndFunctionsContext) EVALSHA() antlr.TerminalNode {
	return s.GetToken(RedisParserEVALSHA, 0)
}

func (s *ScriptingAndFunctionsContext) EVALSHA_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserEVALSHA_RO, 0)
}

func (s *ScriptingAndFunctionsContext) EVAL_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserEVAL_RO, 0)
}

func (s *ScriptingAndFunctionsContext) FCALL() antlr.TerminalNode {
	return s.GetToken(RedisParserFCALL, 0)
}

func (s *ScriptingAndFunctionsContext) FCALL_RO() antlr.TerminalNode {
	return s.GetToken(RedisParserFCALL_RO, 0)
}

func (s *ScriptingAndFunctionsContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(RedisParserFUNCTION, 0)
}

func (s *ScriptingAndFunctionsContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *ScriptingAndFunctionsContext) FunctionSub() IFunctionSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionSubContext)
}

func (s *ScriptingAndFunctionsContext) SCRIPT() antlr.TerminalNode {
	return s.GetToken(RedisParserSCRIPT, 0)
}

func (s *ScriptingAndFunctionsContext) ScriptSub() IScriptSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScriptSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScriptSubContext)
}

func (s *ScriptingAndFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptingAndFunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptingAndFunctionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitScriptingAndFunctions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ScriptingAndFunctions() (localctx IScriptingAndFunctionsContext) {
	this := p
	_ = this

	localctx = NewScriptingAndFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RedisParserRULE_scriptingAndFunctions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserEVAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(RedisParserEVAL)
		}

	case RedisParserEVALSHA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(RedisParserEVALSHA)
		}

	case RedisParserEVALSHA_RO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Match(RedisParserEVALSHA_RO)
		}

	case RedisParserEVAL_RO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Match(RedisParserEVAL_RO)
		}

	case RedisParserFCALL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Match(RedisParserFCALL)
		}

	case RedisParserFCALL_RO:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.Match(RedisParserFCALL_RO)
		}

	case RedisParserFUNCTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(171)
			p.Match(RedisParserFUNCTION)
		}
		{
			p.SetState(172)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(173)
			p.FunctionSub()
		}

	case RedisParserSCRIPT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(174)
			p.Match(RedisParserSCRIPT)
		}
		{
			p.SetState(175)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(176)
			p.ScriptSub()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionSubContext is an interface to support dynamic dispatch.
type IFunctionSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionSubContext differentiates from other interfaces.
	IsFunctionSubContext()
}

type FunctionSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionSubContext() *FunctionSubContext {
	var p = new(FunctionSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_functionSub
	return p
}

func (*FunctionSubContext) IsFunctionSubContext() {}

func NewFunctionSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionSubContext {
	var p = new(FunctionSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_functionSub

	return p
}

func (s *FunctionSubContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionSubContext) DELETE() antlr.TerminalNode {
	return s.GetToken(RedisParserDELETE, 0)
}

func (s *FunctionSubContext) DUMP() antlr.TerminalNode {
	return s.GetToken(RedisParserDUMP, 0)
}

func (s *FunctionSubContext) FLUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserFLUSH, 0)
}

func (s *FunctionSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *FunctionSubContext) KILL() antlr.TerminalNode {
	return s.GetToken(RedisParserKILL, 0)
}

func (s *FunctionSubContext) LIST() antlr.TerminalNode {
	return s.GetToken(RedisParserLIST, 0)
}

func (s *FunctionSubContext) LOAD() antlr.TerminalNode {
	return s.GetToken(RedisParserLOAD, 0)
}

func (s *FunctionSubContext) RESTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserRESTORE, 0)
}

func (s *FunctionSubContext) STATS() antlr.TerminalNode {
	return s.GetToken(RedisParserSTATS, 0)
}

func (s *FunctionSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitFunctionSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) FunctionSub() (localctx IFunctionSubContext) {
	this := p
	_ = this

	localctx = NewFunctionSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RedisParserRULE_functionSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserKILL || _la == RedisParserDUMP || _la == RedisParserRESTORE || (((_la-164)&-(0x1f+1)) == 0 && ((1<<uint((_la-164)))&((1<<(RedisParserDELETE-164))|(1<<(RedisParserFLUSH-164))|(1<<(RedisParserLIST-164))|(1<<(RedisParserLOAD-164))|(1<<(RedisParserSTATS-164))|(1<<(RedisParserHELP-164)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IScriptSubContext is an interface to support dynamic dispatch.
type IScriptSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptSubContext differentiates from other interfaces.
	IsScriptSubContext()
}

type ScriptSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptSubContext() *ScriptSubContext {
	var p = new(ScriptSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_scriptSub
	return p
}

func (*ScriptSubContext) IsScriptSubContext() {}

func NewScriptSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptSubContext {
	var p = new(ScriptSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_scriptSub

	return p
}

func (s *ScriptSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptSubContext) DEBUG() antlr.TerminalNode {
	return s.GetToken(RedisParserDEBUG, 0)
}

func (s *ScriptSubContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(RedisParserEXISTS, 0)
}

func (s *ScriptSubContext) FLUSH() antlr.TerminalNode {
	return s.GetToken(RedisParserFLUSH, 0)
}

func (s *ScriptSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ScriptSubContext) KILL() antlr.TerminalNode {
	return s.GetToken(RedisParserKILL, 0)
}

func (s *ScriptSubContext) LOAD() antlr.TerminalNode {
	return s.GetToken(RedisParserLOAD, 0)
}

func (s *ScriptSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitScriptSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ScriptSub() (localctx IScriptSubContext) {
	this := p
	_ = this

	localctx = NewScriptSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RedisParserRULE_scriptSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserKILL || _la == RedisParserEXISTS || (((_la-165)&-(0x1f+1)) == 0 && ((1<<uint((_la-165)))&((1<<(RedisParserFLUSH-165))|(1<<(RedisParserLOAD-165))|(1<<(RedisParserDEBUG-165))|(1<<(RedisParserHELP-165)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IServerManagementContext is an interface to support dynamic dispatch.
type IServerManagementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServerManagementContext differentiates from other interfaces.
	IsServerManagementContext()
}

type ServerManagementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerManagementContext() *ServerManagementContext {
	var p = new(ServerManagementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_serverManagement
	return p
}

func (*ServerManagementContext) IsServerManagementContext() {}

func NewServerManagementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerManagementContext {
	var p = new(ServerManagementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_serverManagement

	return p
}

func (s *ServerManagementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerManagementContext) ACL() antlr.TerminalNode {
	return s.GetToken(RedisParserACL, 0)
}

func (s *ServerManagementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *ServerManagementContext) AclSub() IAclSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAclSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAclSubContext)
}

func (s *ServerManagementContext) BGREWRITEAOF() antlr.TerminalNode {
	return s.GetToken(RedisParserBGREWRITEAOF, 0)
}

func (s *ServerManagementContext) BGSAVE() antlr.TerminalNode {
	return s.GetToken(RedisParserBGSAVE, 0)
}

func (s *ServerManagementContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(RedisParserCOMMAND, 0)
}

func (s *ServerManagementContext) CommandSub() ICommandSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandSubContext)
}

func (s *ServerManagementContext) CONFIG() antlr.TerminalNode {
	return s.GetToken(RedisParserCONFIG, 0)
}

func (s *ServerManagementContext) ConfigSub() IConfigSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfigSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfigSubContext)
}

func (s *ServerManagementContext) DBSIZE() antlr.TerminalNode {
	return s.GetToken(RedisParserDBSIZE, 0)
}

func (s *ServerManagementContext) DEBUG() antlr.TerminalNode {
	return s.GetToken(RedisParserDEBUG, 0)
}

func (s *ServerManagementContext) FAILOVER() antlr.TerminalNode {
	return s.GetToken(RedisParserFAILOVER, 0)
}

func (s *ServerManagementContext) FLUSHALL() antlr.TerminalNode {
	return s.GetToken(RedisParserFLUSHALL, 0)
}

func (s *ServerManagementContext) FLUSHDB() antlr.TerminalNode {
	return s.GetToken(RedisParserFLUSHDB, 0)
}

func (s *ServerManagementContext) INFO() antlr.TerminalNode {
	return s.GetToken(RedisParserINFO, 0)
}

func (s *ServerManagementContext) LASTSAVE() antlr.TerminalNode {
	return s.GetToken(RedisParserLASTSAVE, 0)
}

func (s *ServerManagementContext) LATENCY() antlr.TerminalNode {
	return s.GetToken(RedisParserLATENCY, 0)
}

func (s *ServerManagementContext) LatencySub() ILatencySubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILatencySubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILatencySubContext)
}

func (s *ServerManagementContext) LOLWUT() antlr.TerminalNode {
	return s.GetToken(RedisParserLOLWUT, 0)
}

func (s *ServerManagementContext) MEMORY() antlr.TerminalNode {
	return s.GetToken(RedisParserMEMORY, 0)
}

func (s *ServerManagementContext) MemorySub() IMemorySubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemorySubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemorySubContext)
}

func (s *ServerManagementContext) MODULE() antlr.TerminalNode {
	return s.GetToken(RedisParserMODULE, 0)
}

func (s *ServerManagementContext) ModuleSub() IModuleSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleSubContext)
}

func (s *ServerManagementContext) MONITOR() antlr.TerminalNode {
	return s.GetToken(RedisParserMONITOR, 0)
}

func (s *ServerManagementContext) PSYNC() antlr.TerminalNode {
	return s.GetToken(RedisParserPSYNC, 0)
}

func (s *ServerManagementContext) REPLCONF() antlr.TerminalNode {
	return s.GetToken(RedisParserREPLCONF, 0)
}

func (s *ServerManagementContext) REPLICAOF() antlr.TerminalNode {
	return s.GetToken(RedisParserREPLICAOF, 0)
}

func (s *ServerManagementContext) RESTOREASKING() antlr.TerminalNode {
	return s.GetToken(RedisParserRESTOREASKING, 0)
}

func (s *ServerManagementContext) ROLE() antlr.TerminalNode {
	return s.GetToken(RedisParserROLE, 0)
}

func (s *ServerManagementContext) SAVE() antlr.TerminalNode {
	return s.GetToken(RedisParserSAVE, 0)
}

func (s *ServerManagementContext) SHUTDOWN() antlr.TerminalNode {
	return s.GetToken(RedisParserSHUTDOWN, 0)
}

func (s *ServerManagementContext) SLAVEOF() antlr.TerminalNode {
	return s.GetToken(RedisParserSLAVEOF, 0)
}

func (s *ServerManagementContext) SLOWLOG() antlr.TerminalNode {
	return s.GetToken(RedisParserSLOWLOG, 0)
}

func (s *ServerManagementContext) SlowlogSub() ISlowlogSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlowlogSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlowlogSubContext)
}

func (s *ServerManagementContext) SWAPDB() antlr.TerminalNode {
	return s.GetToken(RedisParserSWAPDB, 0)
}

func (s *ServerManagementContext) SYNC() antlr.TerminalNode {
	return s.GetToken(RedisParserSYNC, 0)
}

func (s *ServerManagementContext) TIME() antlr.TerminalNode {
	return s.GetToken(RedisParserTIME, 0)
}

func (s *ServerManagementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerManagementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerManagementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitServerManagement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ServerManagement() (localctx IServerManagementContext) {
	this := p
	_ = this

	localctx = NewServerManagementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RedisParserRULE_serverManagement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserACL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(RedisParserACL)
		}
		{
			p.SetState(184)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(185)
			p.AclSub()
		}

	case RedisParserBGREWRITEAOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(RedisParserBGREWRITEAOF)
		}

	case RedisParserBGSAVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(RedisParserBGSAVE)
		}

	case RedisParserCOMMAND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(188)
			p.Match(RedisParserCOMMAND)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RedisParserSPACE {
			{
				p.SetState(189)
				p.Match(RedisParserSPACE)
			}
			{
				p.SetState(190)
				p.CommandSub()
			}

		}

	case RedisParserCONFIG:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.Match(RedisParserCONFIG)
		}
		{
			p.SetState(194)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(195)
			p.ConfigSub()
		}

	case RedisParserDBSIZE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(196)
			p.Match(RedisParserDBSIZE)
		}

	case RedisParserDEBUG:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(197)
			p.Match(RedisParserDEBUG)
		}

	case RedisParserFAILOVER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(198)
			p.Match(RedisParserFAILOVER)
		}

	case RedisParserFLUSHALL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(199)
			p.Match(RedisParserFLUSHALL)
		}

	case RedisParserFLUSHDB:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(200)
			p.Match(RedisParserFLUSHDB)
		}

	case RedisParserINFO:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(201)
			p.Match(RedisParserINFO)
		}

	case RedisParserLASTSAVE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(202)
			p.Match(RedisParserLASTSAVE)
		}

	case RedisParserLATENCY:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(203)
			p.Match(RedisParserLATENCY)
		}
		{
			p.SetState(204)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(205)
			p.LatencySub()
		}

	case RedisParserLOLWUT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(206)
			p.Match(RedisParserLOLWUT)
		}

	case RedisParserMEMORY:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(207)
			p.Match(RedisParserMEMORY)
		}
		{
			p.SetState(208)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(209)
			p.MemorySub()
		}

	case RedisParserMODULE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(210)
			p.Match(RedisParserMODULE)
		}
		{
			p.SetState(211)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(212)
			p.ModuleSub()
		}

	case RedisParserMONITOR:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(213)
			p.Match(RedisParserMONITOR)
		}

	case RedisParserPSYNC:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(214)
			p.Match(RedisParserPSYNC)
		}

	case RedisParserREPLCONF:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(215)
			p.Match(RedisParserREPLCONF)
		}

	case RedisParserREPLICAOF:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(216)
			p.Match(RedisParserREPLICAOF)
		}

	case RedisParserRESTOREASKING:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(217)
			p.Match(RedisParserRESTOREASKING)
		}

	case RedisParserROLE:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(218)
			p.Match(RedisParserROLE)
		}

	case RedisParserSAVE:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(219)
			p.Match(RedisParserSAVE)
		}

	case RedisParserSHUTDOWN:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(220)
			p.Match(RedisParserSHUTDOWN)
		}

	case RedisParserSLAVEOF:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(221)
			p.Match(RedisParserSLAVEOF)
		}

	case RedisParserSLOWLOG:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(222)
			p.Match(RedisParserSLOWLOG)
		}
		{
			p.SetState(223)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(224)
			p.SlowlogSub()
		}

	case RedisParserSWAPDB:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(225)
			p.Match(RedisParserSWAPDB)
		}

	case RedisParserSYNC:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(226)
			p.Match(RedisParserSYNC)
		}

	case RedisParserTIME:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(227)
			p.Match(RedisParserTIME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAclSubContext is an interface to support dynamic dispatch.
type IAclSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAclSubContext differentiates from other interfaces.
	IsAclSubContext()
}

type AclSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAclSubContext() *AclSubContext {
	var p = new(AclSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_aclSub
	return p
}

func (*AclSubContext) IsAclSubContext() {}

func NewAclSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AclSubContext {
	var p = new(AclSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_aclSub

	return p
}

func (s *AclSubContext) GetParser() antlr.Parser { return s.parser }

func (s *AclSubContext) CAT() antlr.TerminalNode {
	return s.GetToken(RedisParserCAT, 0)
}

func (s *AclSubContext) DELUSER() antlr.TerminalNode {
	return s.GetToken(RedisParserDELUSER, 0)
}

func (s *AclSubContext) DRYRUN() antlr.TerminalNode {
	return s.GetToken(RedisParserDRYRUN, 0)
}

func (s *AclSubContext) GETUSER() antlr.TerminalNode {
	return s.GetToken(RedisParserGETUSER, 0)
}

func (s *AclSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *AclSubContext) LIST() antlr.TerminalNode {
	return s.GetToken(RedisParserLIST, 0)
}

func (s *AclSubContext) LOAD() antlr.TerminalNode {
	return s.GetToken(RedisParserLOAD, 0)
}

func (s *AclSubContext) LOG() antlr.TerminalNode {
	return s.GetToken(RedisParserLOG, 0)
}

func (s *AclSubContext) SAVE() antlr.TerminalNode {
	return s.GetToken(RedisParserSAVE, 0)
}

func (s *AclSubContext) SETUSER() antlr.TerminalNode {
	return s.GetToken(RedisParserSETUSER, 0)
}

func (s *AclSubContext) USERS() antlr.TerminalNode {
	return s.GetToken(RedisParserUSERS, 0)
}

func (s *AclSubContext) WHOAMI() antlr.TerminalNode {
	return s.GetToken(RedisParserWHOAMI, 0)
}

func (s *AclSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AclSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AclSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitAclSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) AclSub() (localctx IAclSubContext) {
	this := p
	_ = this

	localctx = NewAclSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RedisParserRULE_aclSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-166)&-(0x1f+1)) == 0 && ((1<<uint((_la-166)))&((1<<(RedisParserLIST-166))|(1<<(RedisParserLOAD-166))|(1<<(RedisParserCAT-166))|(1<<(RedisParserDELUSER-166))|(1<<(RedisParserDRYRUN-166))|(1<<(RedisParserGETUSER-166))|(1<<(RedisParserHELP-166))|(1<<(RedisParserLOG-166))|(1<<(RedisParserSAVE-166))|(1<<(RedisParserSETUSER-166))|(1<<(RedisParserUSERS-166))|(1<<(RedisParserWHOAMI-166)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommandSubContext is an interface to support dynamic dispatch.
type ICommandSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandSubContext differentiates from other interfaces.
	IsCommandSubContext()
}

type CommandSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandSubContext() *CommandSubContext {
	var p = new(CommandSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_commandSub
	return p
}

func (*CommandSubContext) IsCommandSubContext() {}

func NewCommandSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandSubContext {
	var p = new(CommandSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_commandSub

	return p
}

func (s *CommandSubContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandSubContext) COUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserCOUNT, 0)
}

func (s *CommandSubContext) DOCS() antlr.TerminalNode {
	return s.GetToken(RedisParserDOCS, 0)
}

func (s *CommandSubContext) GETKEYS() antlr.TerminalNode {
	return s.GetToken(RedisParserGETKEYS, 0)
}

func (s *CommandSubContext) GETKEYSANDFLAGS() antlr.TerminalNode {
	return s.GetToken(RedisParserGETKEYSANDFLAGS, 0)
}

func (s *CommandSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *CommandSubContext) INFO() antlr.TerminalNode {
	return s.GetToken(RedisParserINFO, 0)
}

func (s *CommandSubContext) LIST() antlr.TerminalNode {
	return s.GetToken(RedisParserLIST, 0)
}

func (s *CommandSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitCommandSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) CommandSub() (localctx ICommandSubContext) {
	this := p
	_ = this

	localctx = NewCommandSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RedisParserRULE_commandSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserINFO || (((_la-166)&-(0x1f+1)) == 0 && ((1<<uint((_la-166)))&((1<<(RedisParserLIST-166))|(1<<(RedisParserHELP-166))|(1<<(RedisParserCOUNT-166))|(1<<(RedisParserDOCS-166))|(1<<(RedisParserGETKEYS-166))|(1<<(RedisParserGETKEYSANDFLAGS-166)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConfigSubContext is an interface to support dynamic dispatch.
type IConfigSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigSubContext differentiates from other interfaces.
	IsConfigSubContext()
}

type ConfigSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigSubContext() *ConfigSubContext {
	var p = new(ConfigSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_configSub
	return p
}

func (*ConfigSubContext) IsConfigSubContext() {}

func NewConfigSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigSubContext {
	var p = new(ConfigSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_configSub

	return p
}

func (s *ConfigSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigSubContext) GET() antlr.TerminalNode {
	return s.GetToken(RedisParserGET, 0)
}

func (s *ConfigSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ConfigSubContext) RESETSTAT() antlr.TerminalNode {
	return s.GetToken(RedisParserRESETSTAT, 0)
}

func (s *ConfigSubContext) REWRITE() antlr.TerminalNode {
	return s.GetToken(RedisParserREWRITE, 0)
}

func (s *ConfigSubContext) SET() antlr.TerminalNode {
	return s.GetToken(RedisParserSET, 0)
}

func (s *ConfigSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitConfigSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ConfigSub() (localctx IConfigSubContext) {
	this := p
	_ = this

	localctx = NewConfigSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RedisParserRULE_configSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-176)&-(0x1f+1)) == 0 && ((1<<uint((_la-176)))&((1<<(RedisParserHELP-176))|(1<<(RedisParserGET-176))|(1<<(RedisParserRESETSTAT-176))|(1<<(RedisParserREWRITE-176))|(1<<(RedisParserSET-176)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILatencySubContext is an interface to support dynamic dispatch.
type ILatencySubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLatencySubContext differentiates from other interfaces.
	IsLatencySubContext()
}

type LatencySubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatencySubContext() *LatencySubContext {
	var p = new(LatencySubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_latencySub
	return p
}

func (*LatencySubContext) IsLatencySubContext() {}

func NewLatencySubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatencySubContext {
	var p = new(LatencySubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_latencySub

	return p
}

func (s *LatencySubContext) GetParser() antlr.Parser { return s.parser }

func (s *LatencySubContext) DOCTOR() antlr.TerminalNode {
	return s.GetToken(RedisParserDOCTOR, 0)
}

func (s *LatencySubContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(RedisParserGRAPH, 0)
}

func (s *LatencySubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *LatencySubContext) HISTOGRAM() antlr.TerminalNode {
	return s.GetToken(RedisParserHISTOGRAM, 0)
}

func (s *LatencySubContext) HISTORY() antlr.TerminalNode {
	return s.GetToken(RedisParserHISTORY, 0)
}

func (s *LatencySubContext) LATEST() antlr.TerminalNode {
	return s.GetToken(RedisParserLATEST, 0)
}

func (s *LatencySubContext) RESET() antlr.TerminalNode {
	return s.GetToken(RedisParserRESET, 0)
}

func (s *LatencySubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatencySubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatencySubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitLatencySub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) LatencySub() (localctx ILatencySubContext) {
	this := p
	_ = this

	localctx = NewLatencySubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RedisParserRULE_latencySub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserRESET || (((_la-176)&-(0x1f+1)) == 0 && ((1<<uint((_la-176)))&((1<<(RedisParserHELP-176))|(1<<(RedisParserDOCTOR-176))|(1<<(RedisParserGRAPH-176))|(1<<(RedisParserHISTOGRAM-176))|(1<<(RedisParserHISTORY-176))|(1<<(RedisParserLATEST-176)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMemorySubContext is an interface to support dynamic dispatch.
type IMemorySubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemorySubContext differentiates from other interfaces.
	IsMemorySubContext()
}

type MemorySubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemorySubContext() *MemorySubContext {
	var p = new(MemorySubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_memorySub
	return p
}

func (*MemorySubContext) IsMemorySubContext() {}

func NewMemorySubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemorySubContext {
	var p = new(MemorySubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_memorySub

	return p
}

func (s *MemorySubContext) GetParser() antlr.Parser { return s.parser }

func (s *MemorySubContext) DOCTOR() antlr.TerminalNode {
	return s.GetToken(RedisParserDOCTOR, 0)
}

func (s *MemorySubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *MemorySubContext) MALLOCSTATS() antlr.TerminalNode {
	return s.GetToken(RedisParserMALLOCSTATS, 0)
}

func (s *MemorySubContext) PURGE() antlr.TerminalNode {
	return s.GetToken(RedisParserPURGE, 0)
}

func (s *MemorySubContext) STATS() antlr.TerminalNode {
	return s.GetToken(RedisParserSTATS, 0)
}

func (s *MemorySubContext) USAGE() antlr.TerminalNode {
	return s.GetToken(RedisParserUSAGE, 0)
}

func (s *MemorySubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemorySubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemorySubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitMemorySub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) MemorySub() (localctx IMemorySubContext) {
	this := p
	_ = this

	localctx = NewMemorySubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RedisParserRULE_memorySub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-168)&-(0x1f+1)) == 0 && ((1<<uint((_la-168)))&((1<<(RedisParserSTATS-168))|(1<<(RedisParserHELP-168))|(1<<(RedisParserDOCTOR-168)))) != 0) || (((_la-206)&-(0x1f+1)) == 0 && ((1<<uint((_la-206)))&((1<<(RedisParserMALLOCSTATS-206))|(1<<(RedisParserPURGE-206))|(1<<(RedisParserUSAGE-206)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IModuleSubContext is an interface to support dynamic dispatch.
type IModuleSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleSubContext differentiates from other interfaces.
	IsModuleSubContext()
}

type ModuleSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleSubContext() *ModuleSubContext {
	var p = new(ModuleSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_moduleSub
	return p
}

func (*ModuleSubContext) IsModuleSubContext() {}

func NewModuleSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleSubContext {
	var p = new(ModuleSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_moduleSub

	return p
}

func (s *ModuleSubContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *ModuleSubContext) LIST() antlr.TerminalNode {
	return s.GetToken(RedisParserLIST, 0)
}

func (s *ModuleSubContext) LOAD() antlr.TerminalNode {
	return s.GetToken(RedisParserLOAD, 0)
}

func (s *ModuleSubContext) LOADEX() antlr.TerminalNode {
	return s.GetToken(RedisParserLOADEX, 0)
}

func (s *ModuleSubContext) UNLOAD() antlr.TerminalNode {
	return s.GetToken(RedisParserUNLOAD, 0)
}

func (s *ModuleSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitModuleSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) ModuleSub() (localctx IModuleSubContext) {
	this := p
	_ = this

	localctx = NewModuleSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RedisParserRULE_moduleSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-166)&-(0x1f+1)) == 0 && ((1<<uint((_la-166)))&((1<<(RedisParserLIST-166))|(1<<(RedisParserLOAD-166))|(1<<(RedisParserHELP-166)))) != 0) || _la == RedisParserLOADEX || _la == RedisParserUNLOAD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISlowlogSubContext is an interface to support dynamic dispatch.
type ISlowlogSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlowlogSubContext differentiates from other interfaces.
	IsSlowlogSubContext()
}

type SlowlogSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlowlogSubContext() *SlowlogSubContext {
	var p = new(SlowlogSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_slowlogSub
	return p
}

func (*SlowlogSubContext) IsSlowlogSubContext() {}

func NewSlowlogSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlowlogSubContext {
	var p = new(SlowlogSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_slowlogSub

	return p
}

func (s *SlowlogSubContext) GetParser() antlr.Parser { return s.parser }

func (s *SlowlogSubContext) GET() antlr.TerminalNode {
	return s.GetToken(RedisParserGET, 0)
}

func (s *SlowlogSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *SlowlogSubContext) LEN() antlr.TerminalNode {
	return s.GetToken(RedisParserLEN, 0)
}

func (s *SlowlogSubContext) RESET() antlr.TerminalNode {
	return s.GetToken(RedisParserRESET, 0)
}

func (s *SlowlogSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlowlogSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SlowlogSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitSlowlogSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) SlowlogSub() (localctx ISlowlogSubContext) {
	this := p
	_ = this

	localctx = NewSlowlogSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RedisParserRULE_slowlogSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserRESET || _la == RedisParserHELP || _la == RedisParserGET || _la == RedisParserLEN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_set
	return p
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_set

	return p
}

func (s *SetContext) GetParser() antlr.Parser { return s.parser }

func (s *SetContext) SADD() antlr.TerminalNode {
	return s.GetToken(RedisParserSADD, 0)
}

func (s *SetContext) SCARD() antlr.TerminalNode {
	return s.GetToken(RedisParserSCARD, 0)
}

func (s *SetContext) SDIFF() antlr.TerminalNode {
	return s.GetToken(RedisParserSDIFF, 0)
}

func (s *SetContext) SDIFFSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserSDIFFSTORE, 0)
}

func (s *SetContext) SINTER() antlr.TerminalNode {
	return s.GetToken(RedisParserSINTER, 0)
}

func (s *SetContext) SINTERCARD() antlr.TerminalNode {
	return s.GetToken(RedisParserSINTERCARD, 0)
}

func (s *SetContext) SINTERSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserSINTERSTORE, 0)
}

func (s *SetContext) SISMEMBER() antlr.TerminalNode {
	return s.GetToken(RedisParserSISMEMBER, 0)
}

func (s *SetContext) SMEMBERS() antlr.TerminalNode {
	return s.GetToken(RedisParserSMEMBERS, 0)
}

func (s *SetContext) SMISMEMBER() antlr.TerminalNode {
	return s.GetToken(RedisParserSMISMEMBER, 0)
}

func (s *SetContext) SMOVE() antlr.TerminalNode {
	return s.GetToken(RedisParserSMOVE, 0)
}

func (s *SetContext) SPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserSPOP, 0)
}

func (s *SetContext) SRANDMEMBER() antlr.TerminalNode {
	return s.GetToken(RedisParserSRANDMEMBER, 0)
}

func (s *SetContext) SREM() antlr.TerminalNode {
	return s.GetToken(RedisParserSREM, 0)
}

func (s *SetContext) SSCAN() antlr.TerminalNode {
	return s.GetToken(RedisParserSSCAN, 0)
}

func (s *SetContext) SUNION() antlr.TerminalNode {
	return s.GetToken(RedisParserSUNION, 0)
}

func (s *SetContext) SUNIONSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserSUNIONSTORE, 0)
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Set() (localctx ISetContext) {
	this := p
	_ = this

	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RedisParserRULE_set)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-225)&-(0x1f+1)) == 0 && ((1<<uint((_la-225)))&((1<<(RedisParserSADD-225))|(1<<(RedisParserSCARD-225))|(1<<(RedisParserSDIFF-225))|(1<<(RedisParserSDIFFSTORE-225))|(1<<(RedisParserSINTER-225))|(1<<(RedisParserSINTERCARD-225))|(1<<(RedisParserSINTERSTORE-225))|(1<<(RedisParserSISMEMBER-225))|(1<<(RedisParserSMEMBERS-225))|(1<<(RedisParserSMISMEMBER-225))|(1<<(RedisParserSMOVE-225))|(1<<(RedisParserSPOP-225))|(1<<(RedisParserSRANDMEMBER-225))|(1<<(RedisParserSREM-225))|(1<<(RedisParserSSCAN-225))|(1<<(RedisParserSUNION-225))|(1<<(RedisParserSUNIONSTORE-225)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISortedSetContext is an interface to support dynamic dispatch.
type ISortedSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSortedSetContext differentiates from other interfaces.
	IsSortedSetContext()
}

type SortedSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortedSetContext() *SortedSetContext {
	var p = new(SortedSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_sortedSet
	return p
}

func (*SortedSetContext) IsSortedSetContext() {}

func NewSortedSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortedSetContext {
	var p = new(SortedSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_sortedSet

	return p
}

func (s *SortedSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SortedSetContext) BZMPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserBZMPOP, 0)
}

func (s *SortedSetContext) BZPOPMAX() antlr.TerminalNode {
	return s.GetToken(RedisParserBZPOPMAX, 0)
}

func (s *SortedSetContext) BZPOPMIN() antlr.TerminalNode {
	return s.GetToken(RedisParserBZPOPMIN, 0)
}

func (s *SortedSetContext) ZADD() antlr.TerminalNode {
	return s.GetToken(RedisParserZADD, 0)
}

func (s *SortedSetContext) ZCARD() antlr.TerminalNode {
	return s.GetToken(RedisParserZCARD, 0)
}

func (s *SortedSetContext) ZCOUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserZCOUNT, 0)
}

func (s *SortedSetContext) ZDIFF() antlr.TerminalNode {
	return s.GetToken(RedisParserZDIFF, 0)
}

func (s *SortedSetContext) ZDIFFSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZDIFFSTORE, 0)
}

func (s *SortedSetContext) ZINCRBY() antlr.TerminalNode {
	return s.GetToken(RedisParserZINCRBY, 0)
}

func (s *SortedSetContext) ZINTER() antlr.TerminalNode {
	return s.GetToken(RedisParserZINTER, 0)
}

func (s *SortedSetContext) ZINTERCARD() antlr.TerminalNode {
	return s.GetToken(RedisParserZINTERCARD, 0)
}

func (s *SortedSetContext) ZINTERSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZINTERSTORE, 0)
}

func (s *SortedSetContext) ZLEXCOUNT() antlr.TerminalNode {
	return s.GetToken(RedisParserZLEXCOUNT, 0)
}

func (s *SortedSetContext) ZMPOP() antlr.TerminalNode {
	return s.GetToken(RedisParserZMPOP, 0)
}

func (s *SortedSetContext) ZMSCORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZMSCORE, 0)
}

func (s *SortedSetContext) ZPOPMAX() antlr.TerminalNode {
	return s.GetToken(RedisParserZPOPMAX, 0)
}

func (s *SortedSetContext) ZPOPMIN() antlr.TerminalNode {
	return s.GetToken(RedisParserZPOPMIN, 0)
}

func (s *SortedSetContext) ZRANDMEMBER() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANDMEMBER, 0)
}

func (s *SortedSetContext) ZRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANGE, 0)
}

func (s *SortedSetContext) ZRANGEBYLEX() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANGEBYLEX, 0)
}

func (s *SortedSetContext) ZRANGEBYSCORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANGEBYSCORE, 0)
}

func (s *SortedSetContext) ZRANGESTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANGESTORE, 0)
}

func (s *SortedSetContext) ZRANK() antlr.TerminalNode {
	return s.GetToken(RedisParserZRANK, 0)
}

func (s *SortedSetContext) ZREM() antlr.TerminalNode {
	return s.GetToken(RedisParserZREM, 0)
}

func (s *SortedSetContext) ZREMRANGEBYLEX() antlr.TerminalNode {
	return s.GetToken(RedisParserZREMRANGEBYLEX, 0)
}

func (s *SortedSetContext) ZREMRANGEBYRANK() antlr.TerminalNode {
	return s.GetToken(RedisParserZREMRANGEBYRANK, 0)
}

func (s *SortedSetContext) ZREMRANGEBYSCORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZREMRANGEBYSCORE, 0)
}

func (s *SortedSetContext) ZREVRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserZREVRANGE, 0)
}

func (s *SortedSetContext) ZREVRANGEBYLEX() antlr.TerminalNode {
	return s.GetToken(RedisParserZREVRANGEBYLEX, 0)
}

func (s *SortedSetContext) ZREVRANGEBYSCORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZREVRANGEBYSCORE, 0)
}

func (s *SortedSetContext) ZREVRANK() antlr.TerminalNode {
	return s.GetToken(RedisParserZREVRANK, 0)
}

func (s *SortedSetContext) ZSCAN() antlr.TerminalNode {
	return s.GetToken(RedisParserZSCAN, 0)
}

func (s *SortedSetContext) ZSCORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZSCORE, 0)
}

func (s *SortedSetContext) ZUNION() antlr.TerminalNode {
	return s.GetToken(RedisParserZUNION, 0)
}

func (s *SortedSetContext) ZUNIONSTORE() antlr.TerminalNode {
	return s.GetToken(RedisParserZUNIONSTORE, 0)
}

func (s *SortedSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortedSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortedSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitSortedSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) SortedSet() (localctx ISortedSetContext) {
	this := p
	_ = this

	localctx = NewSortedSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RedisParserRULE_sortedSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-242)&-(0x1f+1)) == 0 && ((1<<uint((_la-242)))&((1<<(RedisParserBZMPOP-242))|(1<<(RedisParserBZPOPMAX-242))|(1<<(RedisParserBZPOPMIN-242))|(1<<(RedisParserZADD-242))|(1<<(RedisParserZCARD-242))|(1<<(RedisParserZCOUNT-242))|(1<<(RedisParserZDIFF-242))|(1<<(RedisParserZDIFFSTORE-242))|(1<<(RedisParserZINCRBY-242))|(1<<(RedisParserZINTER-242))|(1<<(RedisParserZINTERCARD-242))|(1<<(RedisParserZINTERSTORE-242))|(1<<(RedisParserZLEXCOUNT-242))|(1<<(RedisParserZMPOP-242))|(1<<(RedisParserZMSCORE-242))|(1<<(RedisParserZPOPMAX-242))|(1<<(RedisParserZPOPMIN-242))|(1<<(RedisParserZRANDMEMBER-242))|(1<<(RedisParserZRANGE-242))|(1<<(RedisParserZRANGEBYLEX-242))|(1<<(RedisParserZRANGEBYSCORE-242))|(1<<(RedisParserZRANGESTORE-242))|(1<<(RedisParserZRANK-242))|(1<<(RedisParserZREM-242))|(1<<(RedisParserZREMRANGEBYLEX-242))|(1<<(RedisParserZREMRANGEBYRANK-242))|(1<<(RedisParserZREMRANGEBYSCORE-242))|(1<<(RedisParserZREVRANGE-242))|(1<<(RedisParserZREVRANGEBYLEX-242))|(1<<(RedisParserZREVRANGEBYSCORE-242))|(1<<(RedisParserZREVRANK-242))|(1<<(RedisParserZSCAN-242)))) != 0) || (((_la-274)&-(0x1f+1)) == 0 && ((1<<uint((_la-274)))&((1<<(RedisParserZSCORE-274))|(1<<(RedisParserZUNION-274))|(1<<(RedisParserZUNIONSTORE-274)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStreamContext is an interface to support dynamic dispatch.
type IStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamContext differentiates from other interfaces.
	IsStreamContext()
}

type StreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamContext() *StreamContext {
	var p = new(StreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_stream
	return p
}

func (*StreamContext) IsStreamContext() {}

func NewStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamContext {
	var p = new(StreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_stream

	return p
}

func (s *StreamContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamContext) XACK() antlr.TerminalNode {
	return s.GetToken(RedisParserXACK, 0)
}

func (s *StreamContext) XADD() antlr.TerminalNode {
	return s.GetToken(RedisParserXADD, 0)
}

func (s *StreamContext) XAUTOCLAIM() antlr.TerminalNode {
	return s.GetToken(RedisParserXAUTOCLAIM, 0)
}

func (s *StreamContext) XCLAIM() antlr.TerminalNode {
	return s.GetToken(RedisParserXCLAIM, 0)
}

func (s *StreamContext) XDEL() antlr.TerminalNode {
	return s.GetToken(RedisParserXDEL, 0)
}

func (s *StreamContext) XGROUP() antlr.TerminalNode {
	return s.GetToken(RedisParserXGROUP, 0)
}

func (s *StreamContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RedisParserSPACE, 0)
}

func (s *StreamContext) XgroupSub() IXgroupSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXgroupSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXgroupSubContext)
}

func (s *StreamContext) XINFO() antlr.TerminalNode {
	return s.GetToken(RedisParserXINFO, 0)
}

func (s *StreamContext) XinfoSub() IXinfoSubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXinfoSubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXinfoSubContext)
}

func (s *StreamContext) XLEN() antlr.TerminalNode {
	return s.GetToken(RedisParserXLEN, 0)
}

func (s *StreamContext) XPENDING() antlr.TerminalNode {
	return s.GetToken(RedisParserXPENDING, 0)
}

func (s *StreamContext) XRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserXRANGE, 0)
}

func (s *StreamContext) XREAD() antlr.TerminalNode {
	return s.GetToken(RedisParserXREAD, 0)
}

func (s *StreamContext) XREADGROUP() antlr.TerminalNode {
	return s.GetToken(RedisParserXREADGROUP, 0)
}

func (s *StreamContext) XREVRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserXREVRANGE, 0)
}

func (s *StreamContext) XSETID() antlr.TerminalNode {
	return s.GetToken(RedisParserXSETID, 0)
}

func (s *StreamContext) XTRIM() antlr.TerminalNode {
	return s.GetToken(RedisParserXTRIM, 0)
}

func (s *StreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Stream() (localctx IStreamContext) {
	this := p
	_ = this

	localctx = NewStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RedisParserRULE_stream)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RedisParserXACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(RedisParserXACK)
		}

	case RedisParserXADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.Match(RedisParserXADD)
		}

	case RedisParserXAUTOCLAIM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)
			p.Match(RedisParserXAUTOCLAIM)
		}

	case RedisParserXCLAIM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(251)
			p.Match(RedisParserXCLAIM)
		}

	case RedisParserXDEL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(252)
			p.Match(RedisParserXDEL)
		}

	case RedisParserXGROUP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(253)
			p.Match(RedisParserXGROUP)
		}
		{
			p.SetState(254)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(255)
			p.XgroupSub()
		}

	case RedisParserXINFO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(256)
			p.Match(RedisParserXINFO)
		}
		{
			p.SetState(257)
			p.Match(RedisParserSPACE)
		}
		{
			p.SetState(258)
			p.XinfoSub()
		}

	case RedisParserXLEN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(259)
			p.Match(RedisParserXLEN)
		}

	case RedisParserXPENDING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(260)
			p.Match(RedisParserXPENDING)
		}

	case RedisParserXRANGE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(261)
			p.Match(RedisParserXRANGE)
		}

	case RedisParserXREAD:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(262)
			p.Match(RedisParserXREAD)
		}

	case RedisParserXREADGROUP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(263)
			p.Match(RedisParserXREADGROUP)
		}

	case RedisParserXREVRANGE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(264)
			p.Match(RedisParserXREVRANGE)
		}

	case RedisParserXSETID:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(265)
			p.Match(RedisParserXSETID)
		}

	case RedisParserXTRIM:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(266)
			p.Match(RedisParserXTRIM)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IXgroupSubContext is an interface to support dynamic dispatch.
type IXgroupSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXgroupSubContext differentiates from other interfaces.
	IsXgroupSubContext()
}

type XgroupSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXgroupSubContext() *XgroupSubContext {
	var p = new(XgroupSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_xgroupSub
	return p
}

func (*XgroupSubContext) IsXgroupSubContext() {}

func NewXgroupSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XgroupSubContext {
	var p = new(XgroupSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_xgroupSub

	return p
}

func (s *XgroupSubContext) GetParser() antlr.Parser { return s.parser }

func (s *XgroupSubContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RedisParserCREATE, 0)
}

func (s *XgroupSubContext) CREATECONSUMER() antlr.TerminalNode {
	return s.GetToken(RedisParserCREATECONSUMER, 0)
}

func (s *XgroupSubContext) DELCONSUMER() antlr.TerminalNode {
	return s.GetToken(RedisParserDELCONSUMER, 0)
}

func (s *XgroupSubContext) DESTROY() antlr.TerminalNode {
	return s.GetToken(RedisParserDESTROY, 0)
}

func (s *XgroupSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *XgroupSubContext) SETID() antlr.TerminalNode {
	return s.GetToken(RedisParserSETID, 0)
}

func (s *XgroupSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XgroupSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XgroupSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitXgroupSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) XgroupSub() (localctx IXgroupSubContext) {
	this := p
	_ = this

	localctx = NewXgroupSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RedisParserRULE_xgroupSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserHELP || (((_la-283)&-(0x1f+1)) == 0 && ((1<<uint((_la-283)))&((1<<(RedisParserCREATE-283))|(1<<(RedisParserCREATECONSUMER-283))|(1<<(RedisParserDELCONSUMER-283))|(1<<(RedisParserDESTROY-283))|(1<<(RedisParserSETID-283)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IXinfoSubContext is an interface to support dynamic dispatch.
type IXinfoSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXinfoSubContext differentiates from other interfaces.
	IsXinfoSubContext()
}

type XinfoSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXinfoSubContext() *XinfoSubContext {
	var p = new(XinfoSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_xinfoSub
	return p
}

func (*XinfoSubContext) IsXinfoSubContext() {}

func NewXinfoSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XinfoSubContext {
	var p = new(XinfoSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_xinfoSub

	return p
}

func (s *XinfoSubContext) GetParser() antlr.Parser { return s.parser }

func (s *XinfoSubContext) CONSUMERS() antlr.TerminalNode {
	return s.GetToken(RedisParserCONSUMERS, 0)
}

func (s *XinfoSubContext) GROUPS() antlr.TerminalNode {
	return s.GetToken(RedisParserGROUPS, 0)
}

func (s *XinfoSubContext) HELP() antlr.TerminalNode {
	return s.GetToken(RedisParserHELP, 0)
}

func (s *XinfoSubContext) STREAM() antlr.TerminalNode {
	return s.GetToken(RedisParserSTREAM, 0)
}

func (s *XinfoSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XinfoSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XinfoSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitXinfoSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) XinfoSub() (localctx IXinfoSubContext) {
	this := p
	_ = this

	localctx = NewXinfoSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, RedisParserRULE_xinfoSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserHELP || (((_la-289)&-(0x1f+1)) == 0 && ((1<<uint((_la-289)))&((1<<(RedisParserCONSUMERS-289))|(1<<(RedisParserGROUPS-289))|(1<<(RedisParserSTREAM-289)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringCmdContext is an interface to support dynamic dispatch.
type IStringCmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringCmdContext differentiates from other interfaces.
	IsStringCmdContext()
}

type StringCmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringCmdContext() *StringCmdContext {
	var p = new(StringCmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_stringCmd
	return p
}

func (*StringCmdContext) IsStringCmdContext() {}

func NewStringCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringCmdContext {
	var p = new(StringCmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_stringCmd

	return p
}

func (s *StringCmdContext) GetParser() antlr.Parser { return s.parser }

func (s *StringCmdContext) APPEND() antlr.TerminalNode {
	return s.GetToken(RedisParserAPPEND, 0)
}

func (s *StringCmdContext) DECR() antlr.TerminalNode {
	return s.GetToken(RedisParserDECR, 0)
}

func (s *StringCmdContext) DECRBY() antlr.TerminalNode {
	return s.GetToken(RedisParserDECRBY, 0)
}

func (s *StringCmdContext) GET() antlr.TerminalNode {
	return s.GetToken(RedisParserGET, 0)
}

func (s *StringCmdContext) GETDEL() antlr.TerminalNode {
	return s.GetToken(RedisParserGETDEL, 0)
}

func (s *StringCmdContext) GETEX() antlr.TerminalNode {
	return s.GetToken(RedisParserGETEX, 0)
}

func (s *StringCmdContext) GETRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserGETRANGE, 0)
}

func (s *StringCmdContext) GETSET() antlr.TerminalNode {
	return s.GetToken(RedisParserGETSET, 0)
}

func (s *StringCmdContext) INCR() antlr.TerminalNode {
	return s.GetToken(RedisParserINCR, 0)
}

func (s *StringCmdContext) INCRBY() antlr.TerminalNode {
	return s.GetToken(RedisParserINCRBY, 0)
}

func (s *StringCmdContext) INCRBYFLOAT() antlr.TerminalNode {
	return s.GetToken(RedisParserINCRBYFLOAT, 0)
}

func (s *StringCmdContext) LCS() antlr.TerminalNode {
	return s.GetToken(RedisParserLCS, 0)
}

func (s *StringCmdContext) MGET() antlr.TerminalNode {
	return s.GetToken(RedisParserMGET, 0)
}

func (s *StringCmdContext) MSET() antlr.TerminalNode {
	return s.GetToken(RedisParserMSET, 0)
}

func (s *StringCmdContext) MSETNX() antlr.TerminalNode {
	return s.GetToken(RedisParserMSETNX, 0)
}

func (s *StringCmdContext) PSETEX() antlr.TerminalNode {
	return s.GetToken(RedisParserPSETEX, 0)
}

func (s *StringCmdContext) SET() antlr.TerminalNode {
	return s.GetToken(RedisParserSET, 0)
}

func (s *StringCmdContext) SETEX() antlr.TerminalNode {
	return s.GetToken(RedisParserSETEX, 0)
}

func (s *StringCmdContext) SETNX() antlr.TerminalNode {
	return s.GetToken(RedisParserSETNX, 0)
}

func (s *StringCmdContext) SETRANGE() antlr.TerminalNode {
	return s.GetToken(RedisParserSETRANGE, 0)
}

func (s *StringCmdContext) STRLEN() antlr.TerminalNode {
	return s.GetToken(RedisParserSTRLEN, 0)
}

func (s *StringCmdContext) SUBSTR() antlr.TerminalNode {
	return s.GetToken(RedisParserSUBSTR, 0)
}

func (s *StringCmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringCmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringCmdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitStringCmd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) StringCmd() (localctx IStringCmdContext) {
	this := p
	_ = this

	localctx = NewStringCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RedisParserRULE_stringCmd)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RedisParserGET || _la == RedisParserSET || (((_la-300)&-(0x1f+1)) == 0 && ((1<<uint((_la-300)))&((1<<(RedisParserAPPEND-300))|(1<<(RedisParserDECR-300))|(1<<(RedisParserDECRBY-300))|(1<<(RedisParserGETDEL-300))|(1<<(RedisParserGETEX-300))|(1<<(RedisParserGETRANGE-300))|(1<<(RedisParserGETSET-300))|(1<<(RedisParserINCR-300))|(1<<(RedisParserINCRBY-300))|(1<<(RedisParserINCRBYFLOAT-300))|(1<<(RedisParserLCS-300))|(1<<(RedisParserMGET-300))|(1<<(RedisParserMSET-300))|(1<<(RedisParserMSETNX-300))|(1<<(RedisParserPSETEX-300))|(1<<(RedisParserSETEX-300))|(1<<(RedisParserSETNX-300))|(1<<(RedisParserSETRANGE-300))|(1<<(RedisParserSTRLEN-300))|(1<<(RedisParserSUBSTR-300)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITransactionsContext is an interface to support dynamic dispatch.
type ITransactionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionsContext differentiates from other interfaces.
	IsTransactionsContext()
}

type TransactionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionsContext() *TransactionsContext {
	var p = new(TransactionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RedisParserRULE_transactions
	return p
}

func (*TransactionsContext) IsTransactionsContext() {}

func NewTransactionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionsContext {
	var p = new(TransactionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RedisParserRULE_transactions

	return p
}

func (s *TransactionsContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionsContext) DISCARD() antlr.TerminalNode {
	return s.GetToken(RedisParserDISCARD, 0)
}

func (s *TransactionsContext) EXEC() antlr.TerminalNode {
	return s.GetToken(RedisParserEXEC, 0)
}

func (s *TransactionsContext) MULTI() antlr.TerminalNode {
	return s.GetToken(RedisParserMULTI, 0)
}

func (s *TransactionsContext) UNWATCH() antlr.TerminalNode {
	return s.GetToken(RedisParserUNWATCH, 0)
}

func (s *TransactionsContext) WATCH() antlr.TerminalNode {
	return s.GetToken(RedisParserWATCH, 0)
}

func (s *TransactionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RedisParserVisitor:
		return t.VisitTransactions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RedisParser) Transactions() (localctx ITransactionsContext) {
	this := p
	_ = this

	localctx = NewTransactionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RedisParserRULE_transactions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-320)&-(0x1f+1)) == 0 && ((1<<uint((_la-320)))&((1<<(RedisParserDISCARD-320))|(1<<(RedisParserEXEC-320))|(1<<(RedisParserMULTI-320))|(1<<(RedisParserUNWATCH-320))|(1<<(RedisParserWATCH-320)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
