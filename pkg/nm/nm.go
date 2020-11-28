package nm

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nmdbus"
)

type NetworkManager struct {
	settingsProxy nmdbus.SettingsProxy
}

func NewNetworkManager() (nm *NetworkManager, err error) {
	sp, err := nmdbus.NewSettings()
	if err != nil {
		log.Error("NetworkManager init failed")
		return
	}

	nm = &NetworkManager{
		settingsProxy: sp,
	}
	return
}

type ConnectionType string

const (
	ConnectionTypeEthernet ConnectionType = "802-3-ethernet"
)

type ConnectionIntent interface {
	Serialize() (nmdbus.ConnectionSettings, error)
}

type EthernetConnection struct {
	InterfaceName string
	MTU           uint16
}

func (c *EthernetConnection) Serialize() (
	rv nmdbus.ConnectionSettings, err error) {
	rv = nmdbus.ConnectionSettings{
		"connection": {
			"id":             fmt.Sprintf("conn-%s", c.InterfaceName),
			"interface-name": c.InterfaceName,
			"type":           ConnectionTypeEthernet,
			"autoconnect":    false,
		},
		"802-3-ethernet": {
			"mtu": c.MTU,
		},
	}
	return
}

func (nm *NetworkManager) CreateConnection(conn ConnectionIntent) (
	err error) {
	var connSettings nmdbus.ConnectionSettings
	connSettings, err = conn.Serialize()
	if err != nil {
		return
	}
	return nm.settingsProxy.AddConnection(connSettings)
}
