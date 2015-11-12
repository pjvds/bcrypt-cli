package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "cost",
			Value: 10,
		},
	}
	app.Action = func(ctx *cli.Context) {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fail("failed to read from stdin: " + err.Error())
		}

		hash, err := bcrypt.GenerateFromPassword(input, ctx.Int("cost"))
		if err != nil {
			fail("failed to hash: " + err.Error())
		}

		os.Stdout.Write(hash)
	}
	app.RunAndExitOnError()
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
