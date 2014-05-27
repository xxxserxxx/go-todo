package main

import (
	"fmt"
	"os"

	"github.com/toffanin/go-todo/commands"
	"github.com/toffanin/go-todo/utils"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

func main() {

	// Load Todo.txt CLI environment variables
	utils.LoadConfig()

	// Initialize the app CLI
	app := cli.NewApp()
	app.Name = "go-todo"
	app.Usage = "A simple and extensible utility for managing your todo.txt files"
	app.Version = "1.0.1"
	app.Author = "Mauro Toffanin"
	app.Email = "toffanin.mauro@gmail.com"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{"t", "Prepend the current date to a task automatically when it's added"},
		cli.BoolFlag{"T", "Do not prepend the current date to a task automatically when it's added."},
		cli.BoolFlag{"f", "Forces actions without confirmation or interactive input."},
	}
	app.Commands = []cli.Command{
		commands.GetEnv(),
		commands.GetInit(),
		commands.GetShorthelp(),
		commands.GetAdd(),
		commands.GetAddm(),
		commands.GetList(),
		/*{
			Name:  "status",
			Usage: "Obtain a summary of the todo.txt structure",
			Flags: []cli.Flag{
				cli.StringFlag{"dest, d", "", "specifies a different destination path"},
			},
			Action: func(c *cli.Context) {
				//fmt.Println("status: ", c.Args().First())
			},
		},*/
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
