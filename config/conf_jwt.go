package config

type Jwt struct {
	PublicKeyPath  string `json:"public_key_path" yaml:"public_key_path"`   //public key
	PrivateKeyPath string `json:"private_key_path" yaml:"private_key_path"` //public key
	Expires        int    `json:"expires" yaml:"expires"`                   // 过期时间
	Issuer         string `json:"issuer" yaml:"issuer"`                     // 颁发人
}
