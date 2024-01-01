package shared

import (
	pg "github.com/ananrafs1/cliit/plugin"
	plugin "github.com/hashicorp/go-plugin"
)

func Serve(exec pg.Executor) func() {
	return func() {
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: Handshake,
			Plugins: map[string]plugin.Plugin{
				"grpc_plugin": &GRPCPlugin{Impl: exec},
			},

			// A non-nil value here enables gRPC serving for this plugin...
			GRPCServer: plugin.DefaultGRPCServer,
		})
	}
}
