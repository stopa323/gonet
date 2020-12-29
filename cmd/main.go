package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/glue"
	"github.com/stopa323/gonet/pkg/manifest"
	"github.com/stopa323/gonet/pkg/nm"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	mft, err := manifest.Load("docs/nmcli-samples/eth-conn.yml")
	if err != nil {
		log.Error(err)
		return
	}

	var netManager *nm.NetworkManager
	netManager, err = nm.NewNetworkManager()
	if err != nil {
		log.Error(err)
		return
	}

	for _, eth := range mft.EthernetConnections {
		intent := glue.ToEthernetConnectionIntent(&eth)
		if err = netManager.Connections.Create(intent); err != nil {
			log.Error(err)
			return
		}
	}

}
