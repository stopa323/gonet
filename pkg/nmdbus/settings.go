package nmdbus

import "github.com/godbus/dbus/v5"

const (
	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	/* Methods */
	SettingsListConnections = SettingsInterface + ".ListConnections"
	SettingsAddConnection   = SettingsInterface + ".AddConnection"
)

type SettingsProxy interface {
	ListConnections() ([]ConnectionProxy, error)
	AddConnection(ConnectionSettings) error
}

func NewSettings() (SettingsProxy, error) {
	var s settingsProxy
	return &s, s.init(NetworkManagerInterface, SettingsObjectPath)
}

type settingsProxy struct {
	dbusBase
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

func (s *settingsProxy) AddConnection(settings ConnectionSettings) error {
	var path dbus.ObjectPath
	err := s.obj.Call(SettingsAddConnection, 0, settings).Store(&path)
	if err != nil {
		return err
	}
	return nil
}
