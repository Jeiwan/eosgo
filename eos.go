package eosgo

// EOS implements the main interface
type EOS struct {
	Config EOSConfig
}

// EOSConfig stores config options
type EOSConfig struct {
	NodeosURL string
	KeosURL   string
}

// New returns an EOS
func New(config EOSConfig) *EOS {
	return &EOS{
		Config: config,
	}
}
