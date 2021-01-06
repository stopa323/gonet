package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/stopa323/gonet/pkg/manifest"
	"github.com/stopa323/gonet/pkg/nm"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	mft, err := manifest.Load("docs/nmcli-samples/eth-conn.yml")
	if err != nil {
		panic(err)
	}

	var netManager *nm.NetworkManager
	netManager, err = nm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	for _, c := range mft.Connections {
		log.Debugf("Creating new connection: ")
		if err = netManager.Connections.Create(c); err != nil {
			log.Errorf("Failed to create connection : %s", err)
			return
		}
	}
}
