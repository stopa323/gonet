package objects

import (
	"fmt"
	"strings"
)

type EthernetConnection struct {
	InterfaceName string `yaml:"name"`
	MTU           uint32 `yaml:"mtu"`
}

func (e *EthernetConnection) Type() ConnectionType {
	return TypeEthernet
}

func (e *EthernetConnection) Object() interface{} {
	return e
}

func (e *EthernetConnection) ToString() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%-15s %10s\n", "InterfaceName:", e.InterfaceName)
	fmt.Fprintf(&b, "%-15s %10d\n", "MTU:", e.MTU)
	return b.String()
}
