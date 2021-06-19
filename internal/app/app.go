package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ybkimm/loginhub/internal/app/routes"
	"github.com/ybkimm/loginhub/internal/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	readTimeout     = 45 * time.Second
	writeTimeout    = 15 * time.Second
	shutdownTimeout = 46 * time.Second
)

type Application struct {
	opts   *Options
	cfgs   *configs.Config
	logger *zap.Logger
	server *http.Server
	db     *sql.DB
}

type Options struct {
	ConfigPath string
}

func New(opts *Options) *Application {
	return &Application{
		opts: opts,
	}
}

func (app *Application) Run() error {
	var jobs = []func() error{
		app.loadConfigs,
		app.makeLogger,
		app.makeServer,
		app.makeDB,
	}
	for _, job := range jobs {
		err := job()
		if err != nil {
			if app.logger == nil {
				fmt.Fprintf(os.Stderr, "[CRIT] Setup failed: %s\n", err.Error())
			} else {
				app.logger.Error(
					"Setup failed",
					zap.Error(err),
					zap.String("method", funcName(job)),
				)
			}
			return err
		}
	}
	return app.begin()
}

func (app *Application) loadConfigs() error {
	cfgs, err := configs.Load(app.opts.ConfigPath)
	if err != nil {
		return wrapErr(err, "failed to load configuration file")
	}
	app.cfgs = cfgs
	return nil
}

func (app *Application) makeLogger() error {
	if app.cfgs == nil {
		return ErrConfigNotLoaded
	}

	var logcfg zap.Config
	if app.cfgs.Debug {
		logcfg = zap.NewDevelopmentConfig()
		logcfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		logcfg = zap.NewProductionConfig()
	}

	logger, err := logcfg.Build()
	if err != nil {
		return wrapErr(err, "failed to build logger")
	}

	app.logger = logger
	return nil
}

func (app *Application) makeServer() error {
	if app.cfgs == nil {
		return ErrConfigNotLoaded
	}

	handler := routes.Handler()
	svr := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfgs.Port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      handler,
	}
	app.server = svr
	return nil
}

func (app *Application) makeDB() error {
	if app.cfgs == nil {
		return ErrConfigNotLoaded
	}

	// Build connection string

	explen := 33 +
		len(app.cfgs.Database.Host) +
		len(app.cfgs.Database.Port) +
		len(app.cfgs.Database.User)
	if len(app.cfgs.Database.Password) > 0 {
		explen += 10 +
			len(app.cfgs.Database.Password)
	}
	if app.cfgs.Database.SSL.Enabled {
		explen += 33 +
			len(app.cfgs.Database.SSL.Cert) +
			len(app.cfgs.Database.SSL.Key) +
			len(app.cfgs.Database.SSL.RootCert)
	}

	buf := make([]byte, 0, explen)
	buf = append(buf, "host="...)
	buf = append(buf, app.cfgs.Database.Host...)
	buf = append(buf, " port="...)
	buf = append(buf, app.cfgs.Database.Port...)
	buf = append(buf, " user="...)
	buf = append(buf, app.cfgs.Database.User...)
	if len(app.cfgs.Database.Password) > 0 {
		buf = append(buf, " password="...)
		buf = append(buf, app.cfgs.Database.Password...)
	}
	if app.cfgs.Database.SSL.Enabled {
		buf = append(buf, " sslmode=verify-full sslcert="...)
		buf = append(buf, app.cfgs.Database.SSL.Cert...)
		buf = append(buf, " sslkey="...)
		buf = append(buf, app.cfgs.Database.SSL.Key...)
		buf = append(buf, " sslrootcert="...)
		buf = append(buf, app.cfgs.Database.SSL.RootCert...)
	} else {
		buf = append(buf, " sslmode=disable"...)
	}

	app.logger.Debug(
		"Trying to connect to the db",
		zap.String("connection string", string(buf)),
	)

	// Make connection

	conn, err := sql.Open("postgres", string(buf))
	if err != nil {
		return wrapErr(err, "db connection error")
	}

	err = conn.Ping()
	if err != nil {
		return wrapErr(err, "db ping error")
	}

	app.db = conn
	return nil
}

func (app *Application) begin() error {
	sigchan := make(chan os.Signal, 1)
	defer close(sigchan)

	signal.Notify(sigchan, os.Interrupt, syscall.SIGHUP)

	errchan := make(chan error, 1)
	defer close(errchan)

	go func() {
		app.logger.Info("Server started", zap.Int("port", app.cfgs.Port))
		errchan <- app.server.ListenAndServe()
	}()

	for {
		select {
		case sig := <-sigchan:
			switch sig {
			case os.Interrupt:
				err := app.shutdown()
				if err != nil {
					app.logger.Error("Shutdown error", zap.Error(err))
					errchan <- err
				}

			case syscall.SIGHUP:
				// TODO: Reload configuration
			}

		case err := <-errchan:
			if err != nil && err != http.ErrServerClosed {
				app.logger.Error("http server throwns an error", zap.Error(err))
			}
			return nil
		}
	}
}

func (app *Application) shutdown() error {
	app.logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	return app.server.Shutdown(ctx)
}
