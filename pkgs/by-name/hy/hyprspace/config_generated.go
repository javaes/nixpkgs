// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"

type Config struct {
	// List of addresses to listen on for libp2p traffic.
	ListenAddresses []string `json:"listenAddresses,omitempty"`

	// Trusted peers in the network.
	Peers []ConfigPeersElem `json:"peers,omitempty"`

	// This node's private key.
	PrivateKey string `json:"privateKey"`

	// The services this node provides via the Service Network.
	Services ConfigServices `json:"services,omitempty"`
}

type ConfigPeersElem struct {
	// PeerID of this peer.
	Id string `json:"id"`

	// Friendly name for this peer. (optional)
	Name string `json:"name,omitempty"`

	// Networks to route to this peer. (optional)
	Routes []ConfigPeersElemRoutesElem `json:"routes,omitempty"`
}

type ConfigPeersElemRoutesElem struct {
	// Network specification.
	Net string `json:"net"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfigPeersElemRoutesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["net"]; raw != nil && !ok {
		return fmt.Errorf("field net in ConfigPeersElemRoutesElem: required")
	}
	type Plain ConfigPeersElemRoutesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConfigPeersElemRoutesElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfigPeersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in ConfigPeersElem: required")
	}
	type Plain ConfigPeersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		plain.Name = ""
	}
	if v, ok := raw["routes"]; !ok || v == nil {
		plain.Routes = []ConfigPeersElemRoutesElem{}
	}
	*j = ConfigPeersElem(plain)
	return nil
}

// The services this node provides via the Service Network.
type ConfigServices map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *Config) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["privateKey"]; raw != nil && !ok {
		return fmt.Errorf("field privateKey in Config: required")
	}
	type Plain Config
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["listenAddresses"]; !ok || v == nil {
		plain.ListenAddresses = []string{
			"/ip4/0.0.0.0/tcp/8001",
			"/ip4/0.0.0.0/udp/8001/quic-v1",
			"/ip6/::/tcp/8001",
			"/ip6/::/udp/8001/quic-v1",
		}
	}
	if v, ok := raw["peers"]; !ok || v == nil {
		plain.Peers = []ConfigPeersElem{}
	}
	if v, ok := raw["services"]; !ok || v == nil {
		plain.Services = map[string]string{}
	}
	*j = Config(plain)
	return nil
}
