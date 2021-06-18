package configs

type Config struct {
	Debug bool `json:"debug"`
	Port  int  `json:"port"`

	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		SSL      struct {
			Enabled  bool   `json:"enabled"`
			Cert     string `json:"cert"`
			Key      string `json:"key2"`
			RootCert string `json:"root_cert"`
		}
	} `json:"database"`
}
