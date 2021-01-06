package objects

type ConnectionType string

const (
	TypeEthernet ConnectionType = "Ethernet"
)

type ConnectionBase interface {
	Type() ConnectionType
	Object() interface{}
}
