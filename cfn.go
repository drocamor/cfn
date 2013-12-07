package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"path"
	"github.com/gonuts/flag"
)

var g_cmd *commander.Commander

func init() {
	g_cmd = &commander.Commander{
		Name: path.Base(os.Args[0]),
		Commands: []*commander.Command{
			ex_make_cmd_dot(),
		},

		Flag: flag.NewFlagSet("cfn", flag.ExitOnError),
	}
}

func main() {
	err := g_cmd.Flag.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("**err**: %v\n", err)
		os.Exit(1)
	}

	args := g_cmd.Flag.Args()
	err = g_cmd.Run(args)
	if err != nil {
		fmt.Printf("**err**: %v\n", err)
		os.Exit(1)
	}

	return

}
