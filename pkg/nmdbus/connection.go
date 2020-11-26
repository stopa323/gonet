package nmdbus

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

const (
	ConnectionInterface = SettingsInterface + ".Connection"

	/* Methods */
	ConnectionGetSettings = ConnectionInterface + ".GetSettings"
)

type ConnectionSettings map[string]map[string]interface{}

type ConnectionProxy interface {
	GetSettings() (ConnectionSettings, error)
	ToJSON() ([]byte, error)
}

func NewConnection(objectPath dbus.ObjectPath) (ConnectionProxy, error) {
	var c connectionProxy
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type connectionProxy struct {
	dbusBase
}

func (c *connectionProxy) GetSettings() (ConnectionSettings, error) {
	var settings map[string]map[string]dbus.Variant
	err := c.obj.Call(ConnectionGetSettings, 0).Store(&settings)

	if err != nil {
		return nil, err
	}

	rv := make(ConnectionSettings)

	for k1, v1 := range settings {
		rv[k1] = make(map[string]interface{})

		for k2, v2 := range v1 {
			rv[k1][k2] = v2.Value()
		}
	}

	return rv, nil
}

func (c *connectionProxy) ToJSON() ([]byte, error) {
	s, _ := c.GetSettings()
	return json.Marshal(s)
}
