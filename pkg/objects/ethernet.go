package objects

type EthernetConnection struct {
	InterfaceName string `yaml:"name"`
	MTU           uint16 `yaml:"mtu"`
}

func (e *EthernetConnection) Type() ConnectionType {
	return TypeEthernet
}

func (e *EthernetConnection) Object() interface{} {
	return e
}
