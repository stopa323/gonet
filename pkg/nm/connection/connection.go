package connection

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nm/internal/dbusproxy"
)

type ConnectionType string

const (
	ConnectionTypeEthernet ConnectionType = "802-3-ethernet"
)

type ConnectionIntent interface {
	Serialize() (dbusproxy.ConnectionSettings, error)
	Name() string
}

type ConnectionController interface {
	// Create new connection
	Create(ConnectionIntent) error
}

func NewConnectionController() (ConnectionController, error) {
	// acquire NetworkManager.Settings proxy through D-Bus.
	s, err := dbusproxy.NewSettings()
	if err != nil {
		log.Debugf("failed to acquire settings D-Bus proxy object: %s", err)
		return nil, fmt.Errorf("Connection controller init failed")
	}

	cc := connectionController{
		settings: s,
	}
	return &cc, nil
}

type connectionController struct {
	settings dbusproxy.SettingsProxy
}

func (cc *connectionController) Create(conn ConnectionIntent) error {
	log.Debugf("creating connection: %s", conn.Name())
	connSettings, err := conn.Serialize()
	if err != nil {
		log.Debug(err)
		return fmt.Errorf("failed to serialize connection: %s", conn.Name())
	}

	err = cc.settings.AddConnection(connSettings)
	if err != nil {
		log.Errorf("failed to create connection %s: %s", conn.Name(), err)
		return err
	}
	return err
}
