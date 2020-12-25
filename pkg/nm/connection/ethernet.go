package connection

import (
	"fmt"

	"github.com/stopa323/gonet/pkg/nm/internal/dbusproxy"
)

type EthernetConnection struct {
	InterfaceName string
	MTU           uint16
}

func (e *EthernetConnection) Serialize() (
	rv dbusproxy.ConnectionSettings, err error) {
	rv = dbusproxy.ConnectionSettings{
		"connection": {
			"id":             fmt.Sprintf("conn-%s", e.InterfaceName),
			"interface-name": e.InterfaceName,
			"type":           ConnectionTypeEthernet,
			"autoconnect":    false,
		},
		"802-3-ethernet": {
			"mtu": e.MTU,
		},
	}
	return
}

func (e *EthernetConnection) Name() string {
	return e.InterfaceName
}
