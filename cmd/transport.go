package cmd

import (
	"context"
	"github.com/Microsoft/go-winio"
	"net"
	"net/http"
)

func transport(pipePath string) *http.Transport {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return winio.DialPipeContext(ctx, pipePath)
		},
	}
}
