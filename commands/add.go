package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"../utils"

	"github.com/codegangsta/cli"
)

func GetAddCommand() cli.Command {

	return cli.Command{
		Name:      "add",
		ShortName: "a",
		Usage:     "Add a task to your todo.txt file",
		Description: `
   This command can be used to add the specified task to your todo.txt file on
   its own line.

   Project and content notation are optional. Quotation marks are optional too.

EXAMPLES

   Adds a simple task (quotes are optional):

	  $ todo add "Move out cardboard boxes from the garage"

   Adds tasks with a project notation (quotes are optional):

	  $ todo add "Move out cardboard boxes from the garage +cleaning"
	  $ todo add "Do a load of laundry +cleaning"
	  $ todo add "Vacuum the house +cleaning"

   Adds tasks with a context notation (quotes are optional):

	  $ todo add "Buy eggs and milk @grocery"
	  $ todo add "Buy a cake for friday's dinner party with friends @backery"

   Adds tasks with both project and context notation (quotes are optional):

	  $ todo add "Feed the kitten +BellyOfTheBeast"
	  $ todo add "Buy food with amino acid taurine @petshop +BellyOfTheBeast"
	  $ todo add "Buy huge amont of meat @butcher +BellyOfTheBeast"
	  $ todo add "Hire a bouncer to protect @kitchen cupboard from the cat +BellyOfTheBeast"
`,
		Action: func(c *cli.Context) {
			// collect all the user-submitted arguments in an array
			args := c.Args()

			// debugging
			/*fmt.Printf("(add::Action) args (%d): %s\n", len(args), args)
			fmt.Printf("(add::Action) global flag: %s (%t)\n", "-t", c.GlobalBool("t"))*/

			// task mangler
			task := ""
			switch {
			case len(args) == 0: // no options specified

				// check incorrect usage of the command
				if c.GlobalBool("f") {
					fmt.Print("\nDetected missing option with command \"add [task]\"\n")
					fmt.Print("Usage: todo -f add [task]\"\n\n")
					cli.ShowCommandHelp(c, "add")
					return
				}

				// using interactive input
				fmt.Print("Add: ")
				task, _ = bufio.NewReader(os.Stdin).ReadString('\n')
				//check(err)

			default: // collect all the arguments into a single string
				task = strings.Join(args[0:], " ")
			}

			// TODO: validating input as a task

			// replace return carriage chars with spaces
			// and trim leading / ending spaces
			task = strings.NewReplacer("\n", " ", "\t", " ", "\r", " ").Replace(task)
			task = strings.TrimSpace(task)

			// honor the -t global flag
			if c.GlobalBool("t") {
				date := time.Now().Format("2006-01-02 ")
				task = date + task
			}

			// save the new task
			addAction(task)
		},
	}
}

// Adds a task to a todo.txt file.
func addAction(task string) {
	file := utils.GetTodoEnv("TODO_FILE")

	// TODO: TODOENV["TODO_FILE"] path should be validated somehow
	// before to be stated by os.OpenFile
	//path.Clean(TODOENV["TODO_FILE"])

	// determine the number of tasks in todo.txt
	// TODO: with NewReadWriter the code should be more compact
	//       buf := bufio.NewReadWriter(bufio.NewReader(r), bufio.NewWriter(w))
	fd, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0600)
	utils.Check(err)
	scanner := bufio.NewScanner(fd)
	ntasks := 1
	for scanner.Scan() {
		ntasks++
	}
	if err := scanner.Err(); err != nil {
		fd.Close()
		utils.Check(err)
	}
	//fmt.Printf("n. lines: %d\n", ntasks)
	err = fd.Close()
	utils.Check(err)

	// Open todo.txt in append mode only
	fd, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	utils.Check(err)
	defer fd.Close()

	// use buffered I/O
	writer := bufio.NewWriter(fd)

	// add the task to todo.txt
	_, err = writer.WriteString(task + "\n")
	utils.Check(err)
	err = writer.Flush()
	utils.Check(err)

	// print summary
	fmt.Printf("%d: %s\n", ntasks, task)
	fmt.Printf("TODO: %d added\n", ntasks)
}
