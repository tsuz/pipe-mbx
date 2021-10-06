package main

import (
	"os"
	"pipe-mbx/repo"
	"pipe-mbx/service"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	var dt string
	var raw string

	app := &cli.App{
		Name:  "pipe-mbx",
		Usage: "Pipe data into Mapbox ready layers",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Usage:       "data type to generate",
				Destination: &dt,
			},
			&cli.StringFlag{
				Name:        "raw",
				Aliases:     []string{"r"},
				Usage:       "data path to read raw data",
				Destination: &raw,
			},
		},
		Action: func(c *cli.Context) error {
			repo, err := repo.NewSaveRepo()
			if err != nil {
				log.Fatal(errors.Wrap(err, "Error creating new repo"))
			}

			opts := service.Opts{
				DataType:    dt,
				RawDataPath: raw,
			}

			s, err := service.NewService(repo, opts)
			if err != nil {
				log.Fatal(errors.Wrap(err, "Error creating new service"))
			}

			if err := s.Run(); err != nil {
				log.Fatal(errors.Wrap(err, "Error running service"))
			}

			log.Infof("Succesfully Ran")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error running command line"))
	}
}
