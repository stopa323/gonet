package manifest

import (
	"fmt"
	"io/ioutil"

	"github.com/prometheus/common/log"
	yaml "gopkg.in/yaml.v3"
)

type ConnectionType string

const (
	Ethernet ConnectionType = "Ethernet"
)

type EthernetConnection struct {
	InterfaceName string `yaml:"name"`
	Mtu           uint16 `yaml:"mtu"`
}

type Manifest struct {
	EthernetConnections []EthernetConnection
}

func Load(path string) (mft *Manifest, err error) {
	log.Debugf("reading manifest: %s", path)

	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("manifest read: %w", err)
	}

	mft, err = unmarshalManifest(yamlFile)
	if err != nil {
		return
	}

	log.Infof("loaded manifest: %s", path)
	return
}

func unmarshalManifest(in []byte) (*Manifest, error) {
	type manifest struct {
		Connections []yaml.Node `yaml:"connections"`
	}
	var (
		err error
		m   manifest
		mft Manifest
	)

	err = yaml.Unmarshal(in, &m)
	if err != nil {
		return nil, fmt.Errorf("unmarshal manifest: %w", err)
	}

	for _, c := range m.Connections {
		err = unmarshalConnection(&c, &mft)
		if err != nil {
			return nil, err
		}
	}
	return &mft, nil
}

func unmarshalConnection(conn *yaml.Node, out *Manifest) (err error) {
	type connectionBase struct {
		Type ConnectionType `yaml:"type"`
	}
	var cb connectionBase

	err = conn.Decode(&cb)
	if err != nil {
		return fmt.Errorf("unmarshal base connection: %w", err)
	}

	switch t := cb.Type; t {
	case Ethernet:
		var ethConn EthernetConnection
		err = conn.Decode(&ethConn)
		if err != nil {
			return fmt.Errorf("unmarshal ethernet connection: %w", err)
		}

		out.EthernetConnections = append(out.EthernetConnections, ethConn)
		log.Debugf("unmarshalled new ethernet connection: %s",
			ethConn.InterfaceName)
		return
	default:
		return fmt.Errorf("unmarshal connection: unknown connection type (%s)", t)
	}
}
