package models

type Config struct {
	Url      string   `yaml:"url"`
	Port     string   `yaml:"port"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Address  string   `yaml:"address"`
	Auth     bool     `yaml:"auth"`
	Keys     []string `yaml:"keys"`
}

type Short struct {
	Short string `db:"short"`
	Url   string `db:"url"`
	//Timestamp time.Time `db:"timestamp"`
	Timestamp []uint8 `db:"timestamp"`
}
