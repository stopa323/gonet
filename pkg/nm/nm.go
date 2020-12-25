package nm

import (
	log "github.com/sirupsen/logrus"
	conn "github.com/stopa323/gonet/pkg/nm/connection"
)

type NetworkManager struct {
	Connections conn.ConnectionController
}

func NewNetworkManager() (nm *NetworkManager, err error) {
	cc, err := conn.NewConnectionController()
	if err != nil {
		log.Error("NetworkManager init failed")
		return
	}

	nm = &NetworkManager{
		Connections: cc,
	}
	return
}
