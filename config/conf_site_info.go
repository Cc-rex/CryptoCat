package config

type SiteInfo struct {
	CreatedAt string `yaml:"created_at" json:"created_at"`
	Title     string `yaml:"title" json:"title"`
	Version   string `yaml:"version" json:"version"`
	Email     string `yaml:"email" json:"email"`
	Name      string `yaml:"name" json:"name"`
	Job       string `yaml:"job" json:"job"`
	Addr      string `yaml:"addr" json:"addr"`
	Slogan    string `yaml:"slogan" json:"slogan"`
}
