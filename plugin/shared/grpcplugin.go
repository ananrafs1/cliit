package shared

import (
	"context"
	"io"
	"log"

	pg "github.com/ananrafs1/cliit/plugin"
	proto "github.com/ananrafs1/cliit/plugin/shared/grpc"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCPlugin struct {
	plugin.Plugin
	Impl pg.Executor
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewPluginClient(c)}, nil
}

type GRPCClient struct{ client proto.PluginClient }

func (m *GRPCClient) GetActionMetadata() map[string]map[string]string {
	grpcResp, err := m.client.GetActionMetadata(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil
	}

	resp := make(map[string]map[string]string)
	for k, paramMetadata := range grpcResp.GetActionParameter() {
		resp[k] = paramMetadata.GetParameterMetadata()
	}

	return resp
}

func (m *GRPCClient) GetMetadata() pg.Metadata {
	resp, err := m.client.GetMetadata(context.Background(), &emptypb.Empty{})
	if err != nil {
		return pg.Metadata{}
	}

	return pg.Metadata{
		Title:    resp.GetTitle(),
		Subtitle: resp.GetSubtitle(),
	}
}

func (m *GRPCClient) Execute(act string, params map[string]string) <-chan string {
	stream, err := m.client.Execute(context.Background(), &proto.ExecuteRequest{
		Action: act,
		Params: &proto.ParameterMetadata{
			ParameterMetadata: params,
		},
	})
	if err != nil {
		return nil
	}

	out := make(chan string, 2)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}

			if err == nil {
				out <- resp.GetMessage()
			} else {
				log.Println(err) // dont use panic in your real project
			}

		}
	}()

	return out
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl pg.Executor
}

func (m *GRPCServer) GetActionMetadata(
	ctx context.Context, req *emptypb.Empty) (*proto.ActionMetadataResponse, error) {
	val := m.Impl.GetActionMetadata()

	protoResponse := new(proto.ActionMetadataResponse)
	protoResponse.ActionParameter = map[string]*proto.ParameterMetadata{}
	for action, params := range val {
		protoResponse.ActionParameter[action] = &proto.ParameterMetadata{
			ParameterMetadata: params,
		}
	}
	return protoResponse, nil
}

func (m *GRPCServer) GetMetadata(
	ctx context.Context, req *emptypb.Empty) (*proto.MetadataResponse, error) {
	v := m.Impl.GetMetadata()

	return &proto.MetadataResponse{
		Title:    v.Title,
		Subtitle: v.Subtitle,
	}, nil

}

func (m *GRPCServer) Execute(req *proto.ExecuteRequest, stream proto.Plugin_ExecuteServer) error {
	v := m.Impl.Execute(req.GetAction(), req.GetParams().GetParameterMetadata())

	for msg := range v {
		stream.Send(&proto.ExecuteResponse{
			Message: msg,
		})
	}

	return nil
}
