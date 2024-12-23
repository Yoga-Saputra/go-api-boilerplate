package config

// Database configuration key value
type grpc struct {
	// Host name where the Database is hosted
	GrpcTargetHost string `json:"grpcTargetHost" yaml:"grpcTargetHost"`

	// Host name where the Database is hosted
	GrpcTargetPort int `json:"grpcTargetPort" yaml:"grpcTargetPort"`

	// Host name where the Database is hosted
	GrpcTargetToken string `json:"grpcTargetToken" yaml:"grpcTargetToken"`

	// Host name where the Database is hosted
	GrpcTargetSignature string `json:"grpcTargetSignature" yaml:"grpcTargetSignature"`
}
