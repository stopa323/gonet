package dbusproxy

import (
	"github.com/godbus/dbus/v5"
)

const (
	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	/* Methods */
	SettingsAddConnection   = SettingsInterface + ".AddConnection"
	SettingsListConnections = SettingsInterface + ".ListConnections"
)

type SettingsProxy interface {
	AddConnection(ConnectionSettings) error
	ListConnections() ([]ConnectionProxy, error)
}

func NewSettings() (SettingsProxy, error) {
	var s settingsProxy
	return &s, s.init(NetworkManagerInterface, SettingsObjectPath)
}

type settingsProxy struct {
	dbusProxyCommon
}

// AddConnection adds new connection and save it to disk. It returns any D-Bus
// error encountered.
func (s *settingsProxy) AddConnection(settings ConnectionSettings) error {
	var dummy dbus.ObjectPath
	return s.obj.Call(SettingsAddConnection, 0, settings).Store(&dummy)
}

func (s *settingsProxy) ListConnections() ([]ConnectionProxy, error) {
	var connectionPaths []dbus.ObjectPath

	err := s.obj.Call(SettingsListConnections, 0).Store(&connectionPaths)
	if err != nil {
		return nil, err
	}

	connections := make([]ConnectionProxy, len(connectionPaths))

	for i, path := range connectionPaths {
		connections[i], err = NewConnection(path)
		if err != nil {
			return connections, err
		}
	}

	return connections, nil
}
