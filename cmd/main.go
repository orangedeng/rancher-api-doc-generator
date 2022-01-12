package main

import (
	"os"

	"github.com/orangedeng/rancher-api-doc-generator/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	VERSION = "v0.0.0"
	path    string
)

func main() {
	app := cli.NewApp()
	app.Usage = "This is a rancher norman stype API Docs generator."
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug,D",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:        "output,o",
			EnvVar:      "FILE",
			Usage:       "the output file name",
			Value:       "api-docs.md",
			Destination: &path,
		},
	}
	app.Version = VERSION
	app.Flags = append(app.Flags, utils.GetFlags()...)
	app.Action = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		md, err := utils.GenerateDocs()
		if err != nil {
			return err
		}
		logrus.Infof("Writing docs to file %s", path)
		return md.Write(path)
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Panic(err)
	}
}
