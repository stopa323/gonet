package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nmdbus"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	var (
		sp  nmdbus.SettingsProxy
		err error
		c   nmdbus.ConnectionProxy
	)

	sp, err = nmdbus.NewSettings()
	if err != nil {
		log.Error("could not initialize Settings: ", err)
		panic(err)
	}

	connections, err := sp.ListConnections()
	for _, conn := range connections {
		log.Info(conn.GetSettings())
	}

	conSet := nmdbus.ConnectionSettings{
		"connection": {
			"id":             "conn-eth-test",
			"interface-name": "enp0s8",
			"type":           "802-3-ethernet",
		},
	}
	c, err = sp.AddConnection(conSet)
	if err != nil {
		log.Error("failed to create connection: ", err)
		panic(err)
	}
	log.Info(c)
}
