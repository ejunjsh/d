package main

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/ejunjsh/d/commands"
	_ "github.com/ejunjsh/d/nsenter"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

const usage = `a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "d"
	app.Usage = usage

	app.Commands = []cli.Command{
		InitCommand,
		RunCommand,
		ListCommand,
		ExecCommand,
		RemoveCommand,
		NetworkCommand,
		InstallCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	app.Version = "0.0.1"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
