// reference: https://redis.io/commands
parser grammar RedisParser;

options { tokenVocab=RedisLexer; }

command
    : bitmap
    | clusterManagement
    | connectionManagement
    | generic
    | geospatialIndices
    | hash
    | hyperloglog
    | list
    | pubSub
    | scriptingAndFunctions
    | serverManagement
    | set
    | sortedSet
    | stream
    | stringCmd
    | transactions
    ;

// Bitmap
bitmap
    : BITCOUNT
    | BITFIELD
    | BITFIELD_RO
    | BITOP
    | BITPOS
    | GETBIT
    | SETBIT
    ;

// Cluster management
clusterManagement
    : ASKING
    | CLUSTER SPACE clusterSub
    ;

clusterSub
    : ADDSLOTS
    | ADDSLOTSRANGE
    | BUMPEPOCH
    | COUNTFAILUREREPORTS
    | COUNTKEYSINSLOT
    | DELSLOTS
    | DELSLOTSRANGE
    | FAILOVER
    | FLUSHSLOTS
    | FORGET
    | GETKEYSINSLOT
    | HELP
    | INFO
    | KEYSLOT
    | LINKS
    | MEET
    | MYID
    | NODES
    | REPLICAS
    | REPLICATE
    | RESET
    | SAVECONFIG
    | SETCONFIGEPOCH
    | SETSLOT
    | SHARDS
    | SLAVES
    | SLOTS
    | READONLY
    | READWRITE
    ;

// Connection management
connectionManagement
    : AUTH
    | CLIENT SPACE clientSub
    | ECHO
    | HELLO
    | PING
    | QUIT
    | RESET
    | SELECT SPACE DB
    ;

clientSub
    :CACHING
    |GETNAME
    |GETREDIR
    |HELP
    |ID
    |INFO
    |KILL
    |LIST
    |NOEVICT
    |PAUSE
    |REPLY
    |SETNAME
    |TRACKING
    |TRACKINGINFO
    |UNBLOCK
    |UNPAUSE
    ;

// Generic
generic
    : COPY
    | DEL
    | DUMP
    | EXISTS
    | EXPIRE
    | EXPIREAT
    | EXPIRETIME
    | KEYS
    | MIGRATE
    | MOVE
    | OBJECT SPACE objectSub
    | PERSIST
    | PEXPIRE
    | PEXPIREAT
    | PEXPIRETIME
    | PTTL
    | RANDOMKEY
    | RENAME
    | RENAMENX
    | RESTORE
    | SCAN
    | SORT
    | SORT_RO
    | TOUCH
    | TTL
    | TYPE
    | UNLINK
    | WAIT
    ;

objectSub
    : ENCODING
    | FREQ
    | HELP
    | IDLETIME
    | REFCOUNT
    ;

// Geospatial Indices
geospatialIndices
    : GEOADD
    | GEODIST
    | GEOHASH
    | GEOPOS
    | GEORADIUS
    | GEORADIUSBYMEMBER
    | GEORADIUSBYMEMBER_RO
    | GEORADIUS_RO
    | GEOSEARCH
    | GEOSEARCHSTORE
    ;

// Hash
hash
    : HDEL
    | HEXISTS
    | HGET
    | HGETALL
    | HINCRBY
    | HINCRBYFLOAT
    | HKEYS
    | HLEN
    | HMGET
    | HMSET
    | HRANDFIELD
    | HSCAN
    | HSET
    | HSETNX
    | HSTRLEN
    | HVALS
    ;

// HyperLogLog
hyperloglog
    : PFADD
    | PFCOUNT
    | PFDEBUG
    | PFMERGE
    | PFSELFTEST
    ;

// List
list
    : BLMOVE
    | BLMPOP
    | BLPOP
    | BRPOP
    | BRPOPLPUSH
    | LINDEX
    | LINSERT
    | LLEN
    | LMOVE
    | LMPOP
    | LPOP
    | LPOS
    | LPUSH
    | LPUSHX
    | LRANGE
    | LREM
    | LSET
    | LTRIM
    | RPOP
    | RPOPLPUSH
    | RPUSH
    | RPUSHX
    ;

// Pub/Sub
pubSub
    : PSUBSCRIBE
    | PUBLISH
    | PUBSUB SPACE pubSubSub
    | PUNSUBSCRIBE
    | SPUBLISH
    | SSUBSCRIBE
    | SUBSCRIBE
    | SUNSUBSCRIBE
    | UNSUBSCRIBE
    ;

pubSubSub
    : CHANNELS
    | HELP
    | NUMPAT
    | NUMSUB
    | SHARDCHANNELS
    | SHARDNUMSUB
    ;

// Scripting and functions
scriptingAndFunctions
    : EVAL
    | EVALSHA
    | EVALSHA_RO
    | EVAL_RO
    | FCALL
    | FCALL_RO
    | FUNCTION SPACE functionSub
    | SCRIPT SPACE scriptSub
    ;

functionSub
    : DELETE
    | DUMP
    | FLUSH
    | HELP
    | KILL
    | LIST
    | LOAD
    | RESTORE
    | STATS
    ;

scriptSub
    : DEBUG
    | EXISTS
    | FLUSH
    | HELP
    | KILL
    | LOAD
    ;

// Server management
serverManagement
    : ACL SPACE aclSub
    | BGREWRITEAOF
    | BGSAVE
    | COMMAND (SPACE commandSub)?
    | CONFIG SPACE configSub
    | DBSIZE
    | DEBUG
    | FAILOVER
    | FLUSHALL
    | FLUSHDB
    | INFO
    | LASTSAVE
    | LATENCY SPACE latencySub
    | LOLWUT
    | MEMORY SPACE memorySub
    | MODULE SPACE moduleSub
    | MONITOR
    | PSYNC
    | REPLCONF
    | REPLICAOF
    | RESTOREASKING
    | ROLE
    | SAVE
    | SHUTDOWN
    | SLAVEOF
    | SLOWLOG SPACE slowlogSub
    | SWAPDB
    | SYNC
    | TIME
    ;

aclSub
    : CAT
    | DELUSER
    | DRYRUN
    | GETUSER
    | HELP
    | LIST
    | LOAD
    | LOG
    | SAVE
    | SETUSER
    | USERS
    | WHOAMI
    ;

commandSub
    : COUNT
    | DOCS
    | GETKEYS
    | GETKEYSANDFLAGS
    | HELP
    | INFO
    | LIST
    ;

configSub
    : GET
    | HELP
    | RESETSTAT
    | REWRITE
    | SET
    ;

latencySub
    : DOCTOR
    | GRAPH
    | HELP
    | HISTOGRAM
    | HISTORY
    | LATEST
    | RESET
    ;

memorySub
    : DOCTOR
    | HELP
    | MALLOCSTATS
    | PURGE
    | STATS
    | USAGE
    ;

moduleSub
    : HELP
    | LIST
    | LOAD
    | LOADEX
    | UNLOAD
    ;

slowlogSub
    : GET
    | HELP
    | LEN
    | RESET
    ;

// Set
set
    : SADD
    | SCARD
    | SDIFF
    | SDIFFSTORE
    | SINTER
    | SINTERCARD
    | SINTERSTORE
    | SISMEMBER
    | SMEMBERS
    | SMISMEMBER
    | SMOVE
    | SPOP
    | SRANDMEMBER
    | SREM
    | SSCAN
    | SUNION
    | SUNIONSTORE
    ;

// Sorted set
sortedSet
    : BZMPOP
    | BZPOPMAX
    | BZPOPMIN
    | ZADD
    | ZCARD
    | ZCOUNT
    | ZDIFF
    | ZDIFFSTORE
    | ZINCRBY
    | ZINTER
    | ZINTERCARD
    | ZINTERSTORE
    | ZLEXCOUNT
    | ZMPOP
    | ZMSCORE
    | ZPOPMAX
    | ZPOPMIN
    | ZRANDMEMBER
    | ZRANGE
    | ZRANGEBYLEX
    | ZRANGEBYSCORE
    | ZRANGESTORE
    | ZRANK
    | ZREM
    | ZREMRANGEBYLEX
    | ZREMRANGEBYRANK
    | ZREMRANGEBYSCORE
    | ZREVRANGE
    | ZREVRANGEBYLEX
    | ZREVRANGEBYSCORE
    | ZREVRANK
    | ZSCAN
    | ZSCORE
    | ZUNION
    | ZUNIONSTORE
    ;

// Stream
stream
    : XACK
    | XADD
    | XAUTOCLAIM
    | XCLAIM
    | XDEL
    | XGROUP SPACE xgroupSub
    | XINFO SPACE xinfoSub
    | XLEN
    | XPENDING
    | XRANGE
    | XREAD
    | XREADGROUP
    | XREVRANGE
    | XSETID
    | XTRIM
    ;

xgroupSub
    : CREATE
    | CREATECONSUMER
    | DELCONSUMER
    | DESTROY
    | HELP
    | SETID
    ;

xinfoSub
    : CONSUMERS
    | GROUPS
    | HELP
    | STREAM
    ;

// String
stringCmd
    : APPEND
    | DECR
    | DECRBY
    | GET
    | GETDEL
    | GETEX
    | GETRANGE
    | GETSET
    | INCR
    | INCRBY
    | INCRBYFLOAT
    | LCS
    | MGET
    | MSET
    | MSETNX
    | PSETEX
    | SET
    | SETEX
    | SETNX
    | SETRANGE
    | STRLEN
    | SUBSTR
    ;

// Transactions
transactions
    : DISCARD
    | EXEC
    | MULTI
    | UNWATCH
    | WATCH
    ;