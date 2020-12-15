package main

import (
	"challenge-SmartMEI/app"
	"context"
	"os"
	"os/signal"

	"github.com/micro/go-micro/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	var flagConfig string
	var flagFileName string
	log.SetFormatter(&log.JSONFormatter{})
	cliApp := cli.NewApp()
	cliApp.Name = "Test"
	cliApp.Description = "Challenge for HASH by Ednailson Junior"
	cliApp.Version = "Version"
	cliApp.EnableBashCompletion = true
	cliApp.Commands = []cli.Command{
		{
			Name:    "config sample generator",
			Aliases: []string{"csg"},
			Action: func(cli *cli.Context) error {
				return configSampleGenerator(flagFileName)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file-name, f",
					Value:       "./config.json",
					Usage:       "Config sample file name",
					Destination: &flagFileName,
				},
			},
		},
		{
			Name:    "run application",
			Aliases: []string{"run"},
			Action: func(cli *cli.Context) error {
				return runApplication(flagConfig)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Value:       "./config.json",
					Usage:       "config cliApp file",
					Destination: &flagConfig,
				},
			},
		},
	}
	err := cliApp.Run(os.Args)
	if err != nil {
		log.WithField("error", err.Error()).Errorf("error on running application")
	}
}

func runApplication(flagConfig string) error {
	ctx := gracefullyShutdown()
	var cfg app.Config
	if err := errors.Wrap(config.LoadFile(flagConfig), "failed to get config file"); err != nil {
		return err
	}
	if err := errors.Wrap(config.Scan(&cfg), "failed to read from config file"); err != nil {
		return err
	}
	log.SetLevel(log.Level(cfg.LogLevel))
	application, err := app.LoadApp(cfg)
	if err != nil {
		return err
	}
	log.Infof("application running")
	chErr := application.Run()
	select {
	case err := <-chErr:
		if err != nil {
			log.WithField("error", err.Error()).Errorf("something went wrong on the application")
			application.Close()
			return err
		}
	case <-ctx.Done():
		application.Close()
	}
	log.Infof("application closed")
	return nil
}

func gracefullyShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("gracefully shutdown")
		cancel()
	}()
	return ctx
}

func configSampleGenerator(flagFileName string) error {
	return errors.Wrap(app.NewConfigFile(flagFileName), "could not create a new config file")
}
