package nm

import (
	"fmt"

	"github.com/stopa323/gonet/pkg/nm/internal/dbusproxy"
	o "github.com/stopa323/gonet/pkg/objects"
)

type nmConnectionType string

const (
	Ethernet nmConnectionType = "802-3-ethernet"
)

func SerializeConnection(conn o.ConnectionBase) (
	dbusproxy.ConnectionSettings, error) {
	switch t := conn.Type(); t {
	case o.TypeEthernet:
		c, ok := conn.Object().(*o.EthernetConnection)
		if !ok {
			return nil, fmt.Errorf("convert to EthernetConnection")
		}
		out := serializeEthernet(c)
		return out, nil
	default:
		return nil, fmt.Errorf("unknown object type")
	}
}

func serializeEthernet(c *o.EthernetConnection) dbusproxy.ConnectionSettings {
	out := dbusproxy.ConnectionSettings{
		"connection": {
			"id":             fmt.Sprintf("conn-%s", c.InterfaceName),
			"interface-name": c.InterfaceName,
			"type":           Ethernet,
			"autoconnect":    false,
		},
		"802-3-ethernet": {
			"mtu": c.MTU,
		},
	}
	return out
}
