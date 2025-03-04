package blockchain

// NetworkConfig holds blockchain settings
type NetworkConfig struct {
	ChainID  int
	NodeURL  string
	Protocol string
}

// GetDefaultConfig returns default blockchain settings
func GetDefaultConfig() NetworkConfig {
	return NetworkConfig{
		ChainID:  1,
		NodeURL:  "https://rpc.goai.network",
		Protocol: "PoS",
	}
}
