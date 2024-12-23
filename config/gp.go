package config

// Game provider configuration key value
type gameProvider struct {
	// API Host
	APIHost map[string]string `json:"apiHost"`

	// API Token
	APIToken map[string]string `json:"apiToken"`
}
