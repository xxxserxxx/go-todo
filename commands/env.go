package commands

import (
	"fmt"

	"../utils"

	"github.com/codegangsta/cli"
)

func GetEnvCommand() cli.Command {

	return cli.Command{
		Name:        "env",
		Usage:       "Print `todo` environment information",
		Description: "By default env prints information as a shell script.\n   The environment info will be dumped in straight-forward\n   form suitable for sourcing into a shell script.\n\n   If one or more variable names is given as arguments, env\n   prints the value of each named variable on its own line.",
		Action: func(c *cli.Context) {
			// collect all the user-submitted arguments in an array
			args := c.Args()

			// debugging
			/*fmt.Printf("pwd: %s\n", pwd)
			fmt.Printf("HOME=\"%s\"\n", ENV["HOME"])*/

			switch args.Present() {
			case true:
				// print only the required environment variables
				for _, arg := range args {
					fmt.Printf("%s=\"%s\"\n", arg, utils.TODOENV[arg])
				}
			case false:
				// print all the environment variables
				for k, v := range utils.TODOENV {
					fmt.Printf("%s=\"%s\"\n", k, v)
				}
			}
		},
	}
}
