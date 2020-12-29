package glue

import (
	m "github.com/stopa323/gonet/pkg/manifest"
	c "github.com/stopa323/gonet/pkg/nm/connection"
)

func ToEthernetConnectionIntent(e *m.EthernetConnection) c.ConnectionIntent {
	intent := c.EthernetConnection{
		InterfaceName: e.InterfaceName,
		MTU:           e.Mtu,
	}
	return &intent
}
