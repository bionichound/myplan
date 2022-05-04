package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	var msg string
	app := &cli.App{
		Name:      "myplan",
		HelpName:  "myplan",
		Usage:     "Simple time planner and task monitoring for your daily needs",
		Copyright: "Bionichound 2022",
		Commands: []*cli.Command{
			{
				Name:      "new",
				Aliases:   []string{"n"},
				ArgsUsage: "title",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "message",
						Usage:       "Add a description message to your task",
						Aliases:     []string{"m"},
						Destination: &msg,
					},
				},
				Usage: "Add a new task to your daily",
				Action: func(ctx *cli.Context) error {
					fmt.Println("Added a new task")
					task := NewTask(ctx.Args().First(), msg)
					AddToStore(task)
					PrintStore()
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "Remove a task from today",
				Action: func(ctx *cli.Context) error {
					fmt.Println("This is a list of current tasks")
					PrintEnumerated()
					fmt.Println("Please enter the index of the value to remove")
					reader := bufio.NewReader(os.Stdin)
					ind, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					intInd, err := strconv.Atoi(ind[:len(ind)-1])
					if err != nil {
						log.Fatal(err)
					}
					if intInd >= len(store.Tasks) || intInd < 0 {
						fmt.Println("Please make sure the index is valid")
						return nil
					}
					RemoveFromStore(intInd)
					SaveStore()
					fmt.Println("Item removed")
					fmt.Println("")
					PrintStore()
					return nil
				},
			},
			{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "Print the list of tasks",
				Action: func(ctx *cli.Context) error {
					PrintStore()
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "Mark as done",
				Action: func(ctx *cli.Context) error {
					fmt.Println("This is a list of current tasks")
					PrintEnumerated()
					fmt.Println("Please enter the index of the value to mark as done")
					reader := bufio.NewReader(os.Stdin)
					ind, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					intInd, err := strconv.Atoi(ind[:len(ind)-1])
					if err != nil {
						log.Fatal(err)
					}
					if intInd >= len(store.Tasks) || intInd < 0 {
						fmt.Println("Please make sure the index is valid")
						return nil
					}
					MarkAsDone(intInd)
					SaveStore()
					fmt.Println("")
					PrintStore()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
