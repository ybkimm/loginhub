package configs

type Config struct {
	Debug bool `json:"debug"`
	Port  int  `json:"port"`

	Database struct {
		DBName   string `json:"db_name"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		SSL      struct {
			Enabled  bool   `json:"enabled"`
			Cert     string `json:"cert"`
			Key      string `json:"key2"`
			RootCert string `json:"root_cert"`
		}
	} `json:"database"`
}
