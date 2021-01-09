package nm

import (
	"fmt"

	"github.com/stopa323/gonet/pkg/nm/internal/dbusproxy"
	obj "github.com/stopa323/gonet/pkg/objects"
)

func DeserializeConnection(c dbusproxy.ConnectionSettings) (
	obj.ConnectionBase, error) {
	connType := c["connection"]["type"]
	switch connType {
	case string(Ethernet):
		return deserializeEthernet(c), nil
	default:
		return nil, fmt.Errorf("unknown connection type: %s", connType)
	}
}

func deserializeEthernet(c dbusproxy.ConnectionSettings) obj.ConnectionBase {
	mtu, ok := c["802-3-ethernet"]["mtu"].(uint32)
	if !ok {
		mtu = 0
	}

	ethernetConnection := obj.EthernetConnection{
		InterfaceName: c["connection"]["interface-name"].(string),
		MTU:           mtu,
	}
	return &ethernetConnection
}
