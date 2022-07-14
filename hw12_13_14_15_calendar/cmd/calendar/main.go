package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/home/sergey/Projects/otus-hw/hw12_13_14_15_calendar/configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)

	logg, err := logger.New(config.Logger.Type, config.Logger.Directory, config.Logger.Level)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var storage storage.Storage
	storageType := config.Storage.Type
	if storageType == "MEMO" {
		storage = memorystorage.New()
	} else if storageType == "DB" {
		storage, err = sqlstorage.New(config.Storage.GetDbConnectionString())
		if err != nil {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}

	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar, config.Server.Host, config.Server.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
