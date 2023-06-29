package tunnelserver

import (
	"context"
	"io"
	"net"

	"github.com/loft-sh/devpod/pkg/agent/tunnel"
	"github.com/loft-sh/devpod/pkg/stdio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTunnelClient(reader io.Reader, writer io.WriteCloser, exitOnClose bool) (tunnel.TunnelClient, error) {
	pipe := stdio.NewStdioStream(reader, writer, exitOnClose)

	// Set up a connection to the server.
	conn, err := grpc.Dial("", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
		return pipe, nil
	}))
	if err != nil {
		return nil, err
	}

	return tunnel.NewTunnelClient(conn), nil
}
