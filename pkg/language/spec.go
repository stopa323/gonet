package language

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Connections *Connections
}

type Connections struct {
	Ethernet map[string]EthernetConnection
}

func NewConfig() *Config {
	conn := Connections{
		Ethernet: make(map[string]EthernetConnection),
	}
	cfg := Config{
		Connections: &conn,
	}
	return &cfg
}

func (c *Config) AddConnection(kind string, name string, body hcl.Body) (
	diags hcl.Diagnostics) {
	if c.ContainsConnection(kind, name) {
		diags = append(diags, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "connection already exists",
			Detail:   fmt.Sprintf("%s connection: %s already exists", kind, name),
		})
		return
	}

	switch kind {
	case "ethernet":
		log.Debugf("found `%s` connection named `%s`", kind, name)
		var connection EthernetConnection
		diags = gohcl.DecodeBody(body, nil, &connection)
		if diags.HasErrors() {
			return
		}
		c.Connections.Ethernet[name] = connection
		return
	default:
		panic("you fucked up somewhere")
	}
}

func (c *Config) ContainsConnection(kind string, name string) bool {
	switch kind {
	case "ethernet":
		_, found := c.Connections.Ethernet[name]
		return found
	default:
		return false
	}
}

func isValidConnectionKind(kind string) bool {
	switch kind {
	case "ethernet":
		return true
	default:
		return false
	}
}

type EthernetConnection struct {
	InterfaceName string `hcl:"ifname,attr"`
	MTU           uint16 `hcl:"mtu,attr"`
}
