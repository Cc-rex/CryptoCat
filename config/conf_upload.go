package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` //image size
	Path string `yaml:"path" json:"path"` //image storage path
}
