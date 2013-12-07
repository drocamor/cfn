package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"

)

func ex_make_cmd_dot() *commander.Command {
	cmd := &commander.Command{
		Run:       ex_run_cmd_dot,
		UsageLine: "dot template1 [templateN...]",
		Short:     "generates Graphviz dot files from templates",
		Long: `
generates Graphviz dot files from templates.

ex:
 $ cfn dot foo.template bar.template
`,
		Flag: *flag.NewFlagSet("my-cmd-cmd1", flag.ExitOnError),
	}
	
	return cmd
}

func ex_run_cmd_dot(cmd *commander.Command, args []string) {
	name := "cfn-" + cmd.Name()
	
	fmt.Printf("%s: hello from dot \n", name)
}
