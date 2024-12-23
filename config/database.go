package config

// Database configuration key value
type database struct {
	// Host name where the Database is hosted
	Host string `json:"host"`

	// Port number of Database connection
	Port int `json:"port"`

	// User name of Database connection
	User string `json:"user"`

	// Password of Database conenction
	Password string `json:"password"`

	// Name of Database that want to connect
	Name string `json:"name"`

	// Dialect is varian or type of database query language
	Dialect string `json:"dialect"`

	// Identifier represent custom identifier of this connection
	Identifier string `json:"identifier" yaml:"identifier"`

	// Database resolver
	//
	// Can be called as other database sources
	Resolver resolver `json:"resolver"`
}

// Database configuration key value
type database2 struct {
	// Host name where the Database is hosted
	Host string `json:"host"`

	// Port number of Database connection
	Port int `json:"port"`

	// User name of Database connection
	User string `json:"user"`

	// Password of Database conenction
	Password string `json:"password"`

	// Name of Database that want to connect
	Name string `json:"name"`

	// Dialect is varian or type of database query language
	Dialect string `json:"dialect"`

	// Identifier represent custom identifier of this connection
	Identifier string `json:"identifier" yaml:"identifier"`

	// Database resolver
	//
	// Can be called as other database sources
	Resolver resolver `json:"resolver"`
}

type resolver struct {
	// Array of other database sources
	Sources []database `json:"sources"`

	// Array of other database that used as replications
	Replicas []database `json:"replicas"`
}
