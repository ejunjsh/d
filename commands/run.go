package commands

import "github.com/urfave/cli"

var runCommand = cli.Command{
	Name:  "run",
	Usage: `d run -ti <image> <command>`,
}
