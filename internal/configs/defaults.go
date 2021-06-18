package configs

import (
	"encoding/json"
	"log"
)

const DefaultConfigFilePath = "/etc/loginhub/loginhub.conf"

var DefaultConfig Config

func init() {
	err := json.Unmarshal([]byte(`
		{
			"debug": true,
			"port": 7280,
			"database": {
				"db_name": "loginhub",
				"host": "localhost",
				"port": "5432",
				"user": "loginhub",
				"password": "",
				"ssl": {
					"enabled": false
				}
			}
		}
	`), &DefaultConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
