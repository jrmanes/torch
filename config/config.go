package config

// MutualPeersConfig represents the configuration structure.
type MutualPeersConfig struct {
	MutualPeers []*MutualPeer `yaml:"mutualPeers"` // MutualPeers list of mutual peers.
}

// MutualPeer represents a mutual peer structure.
type MutualPeer struct {
	ConsensusNode    string `yaml:"consensusNode,omitempty"`    // ConsensusNode name
	Peers            []Peer `yaml:"peers"`                      //  Peer list of peers.
	TrustedPeersPath string `yaml:"trustedPeersPath,omitempty"` // TrustedPeersPath specify the path to keep the files
}

// Peer represents a peer structure.
type Peer struct {
	// NodeName of the peer node.
	NodeName         string   `yaml:"nodeName"`                   // NodeName name of the sts/deployment
	ContainerName    string   `yaml:"containerName"`              // ContainerName name of the main container
	ConnectsAsEnvVar bool     `yaml:"connectsAsEnvVar,omitempty"` // ConnectsAsEnvVar use the value as env var
	ConnectsTo       []string `yaml:"connectsTo,omitempty"`       // ConnectsTo list of nodes that it will connect to
}
