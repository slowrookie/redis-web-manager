module github.com/slowrookie/redis-web-manager

go 1.18

require (
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/slowrookie/redis-web-manager/api v0.0.0-20220525113429-86ac49708877
	go.lsp.dev/jsonrpc2 v0.10.0
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220527190237-ee62e23da966 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/slowrookie/redis-web-manager/api => ../api
