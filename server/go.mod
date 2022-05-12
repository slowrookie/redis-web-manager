module github.com/slowrookie/redis-web-manager

go 1.18

require (
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/slowrookie/redis-web-manager/api v0.0.0-00010101000000-000000000000
	go.lsp.dev/jsonrpc2 v0.9.0
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220418222510-f25a4f6275ed // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/segmentio/encoding v0.2.7 // indirect
	go.lsp.dev/pkg v0.0.0-20210125030640-b6310ac75a91 // indirect
	golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/slowrookie/redis-web-manager/api => ../api
