package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

var InstallCommand = cli.Command{
	Name:  "install",
	Usage: "install an image into d",
	Action: func(context *cli.Context) error {
		return installImage(context)
	},
}

const MISS_IMAGE_TAR = "Missing image tar"

func installImage(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		return fmt.Errorf(MISS_IMAGE_TAR)
	}
	var cmdArray []string
	for _, arg := range ctx.Args() {
		cmdArray = append(cmdArray, arg)
	}

	tar := cmdArray[0]

	f, err := os.OpenFile(tar, os.O_WRONLY, 0)
	if err != nil {
		return nil
	}

	os.MkdirAll("/var/lib/d/i/"+strings.TrimSuffix(f.Name(), ".tar"), 0662)

	if strings.HasSuffix(f.Name(), ".tar") {
		if _, err := exec.Command("tar", "-xvf", tar, "-C", "/var/lib/d/i/"+strings.TrimSuffix(f.Name(), ".tar")).CombinedOutput(); err != nil {
			return err
		}
	} else {
		return fmt.Errorf(MISS_IMAGE_TAR)
	}
	return nil
}
