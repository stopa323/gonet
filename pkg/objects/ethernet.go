package objects

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
