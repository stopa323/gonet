package nm

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nm/internal/dbusproxy"
	obj "github.com/stopa323/gonet/pkg/objects"
)

type ConnectionController interface {
	// Create new connection
	Create(obj.ConnectionBase) error
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

func (cc *connectionController) Create(c obj.ConnectionBase) error {
	connection, err := SerializeConnection(c)
	if err != nil {
		return fmt.Errorf("serialize connection: %w", err)
	}

	err = cc.settings.AddConnection(connection)
	if err != nil {
		return fmt.Errorf("add connection: %w", err)
	}
	return nil
}
