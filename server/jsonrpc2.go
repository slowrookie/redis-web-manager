package main

import (
	"context"
	"errors"
	"io"
	"sync"

	"go.lsp.dev/jsonrpc2"
)

var handlers = []jsonrpc2.Handler{}

// RegisterHandler .
func RegisterHandler(handler jsonrpc2.Handler) {
	handlers = append(handlers, handler)
}

// DispatchHandlers .
func DispatchHandlers() jsonrpc2.Handler {
	return func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if len(handlers) <= 0 {
			return errors.New("no handler found")
		}
		var err error
		for _, h := range handlers {
			err = h(ctx, reply, req)
			if err != nil {
				break
			}
		}
		return err
	}
}

// ServerConn over http or websocket
func ServerConn(ctx context.Context, netConn io.ReadWriteCloser, server jsonrpc2.StreamServer) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	closedConns := make(chan error)
	stream := jsonrpc2.NewRawStream(netConn)
	go func() {
		conn := jsonrpc2.NewConn(stream)
		closedConns <- server.ServeStream(ctx, conn)
		stream.Close()
	}()
	wg.Wait()
}
