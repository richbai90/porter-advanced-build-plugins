package build

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

// PluginInterface for the builder. This first part of the
// three-part plugin key is only seen/used by the plugins when the host is
// communicating with the plugin and is not exposed to users.
const PluginInterface = "builder"

var _ plugin.Plugin = &Plugin{}

// Plugin is a generic type of plugin for working with any implementation of a secret store.
type Plugin struct {
	Impl BuilderPlugin
}

type BuilderPlugin interface {
	Command() string
}

type Server struct {
	Impl BuilderPlugin
}

type Client struct {
	client *rpc.Client
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &Server{Impl: p.Impl}, nil
}

func (Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return Client{client: c}, nil
}

