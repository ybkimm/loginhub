package main

import (
	"flag"
	"os"

	"github.com/ybkimm/loginhub/internal/app"
	"github.com/ybkimm/loginhub/internal/configs"

	_ "github.com/lib/pq"
)

var flagConfig *string

func init() {
	flagConfig = flag.String(
		"c",
		configs.DefaultConfigFilePath,
		"location of configuration file",
	)
}

func main() {
	flag.Parse()

	instance := app.New(&app.Options{
		ConfigPath: *flagConfig,
	})

	err := instance.Run()
	if err != nil {
		os.Exit(1)
	}
}
