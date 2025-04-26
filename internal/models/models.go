package models

type Config struct {
	Url      string   `yaml:"url"`
	Port     string   `yaml:"port"`
	Type     string   `yaml:"dbtype"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Address  string   `yaml:"address"`
	Auth     bool     `yaml:"auth"`
	Keys     []string `yaml:"keys"`
}

type Short struct {
	Short     string  `db:"short"`
	Url       string  `db:"url"`
	Timestamp []uint8 `db:"timestamp"`
	Ip        string  `db:"ip"`
}
