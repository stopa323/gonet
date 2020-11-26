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
}
