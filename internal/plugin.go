package internal

import (
	"errors"
	"os"
	"os/exec"

	"github.com/ananrafs1/cliit/plugin"
	"github.com/ananrafs1/cliit/plugin/shared"
	pg "github.com/hashicorp/go-plugin"
)

func GetPluginList() []string {
	files, _ := os.ReadDir("./bin/plugin")
	if len(files) < 1 {
		return []string{}
	}

	res := make([]string, 0, len(files))
	for _, file := range files {
		res = append(res, file.Name())
	}
	return res
}

func AccessPlugin(fileName string) (p plugin.Executor, closer func(), err error) {
	client := pg.NewClient(&pg.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command(fileName),
		AllowedProtocols: []pg.Protocol{
			pg.ProtocolGRPC},
	})

	closer = func() { client.Kill() }
	defer func() {
		if err != nil {
			closer()
		}
	}()

	rpcClient, err := client.Client()
	if err != nil {
		return nil, func() {}, err
	}

	raw, err := rpcClient.Dispense(fileName)
	if err != nil {
		return nil, func() {}, err
	}

	res, ok := raw.(plugin.Executor)
	if !ok {
		return nil, func() {}, errors.New("invalid plugin")
	}

	return res, closer, nil
}
