package main

import (
	"cw1/database"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(fmt.Sprintf("Executing action: \"%s\"", os.Args[1]))
	switch os.Args[1] {
	case "add":
		name, description, priority := initAddFlags(os.Args[2:])
		fmt.Printf("Created action: %s (%s) with priority %s", name, description, priority)

		err, _ := database.AddItem(name, description, priority)

		if err != nil {
			log.Panic(err)
		}
	case "remove":
		// TODO: Write it

	case "list":
		// TODO: Write it

	default:
		fmt.Printf("Provide action: %s, %s or %s", "add", "remove", "list")
	}
}

func initAddFlags(args []string) (name, description string, priority database.Priority) {
	addFlags := flag.NewFlagSet("add", flag.ExitOnError)

	priority = database.LOW // TODO: Add validation

	var priorityStr string
	addFlags.StringVar(
		&priorityStr,
		"priority",
		string(database.MEDIUM),
		fmt.Sprintf(
			"Pirority of the issue. Possible values: %s, %s, %s",
			database.LOW,
			database.MEDIUM,
			database.HIGH,
		),
	)
	addFlags.StringVar(&name, "name", "", "Name of the issue. (REQUIRED)")
	addFlags.StringVar(&description, "description", "", "Description of the issue.")

	err := addFlags.Parse(args)
	if err != nil {
		log.Panic(err)
	}

	// TODO: Add validation of priority, name and description len > 1
	// TODO: Add estimated time to complete

	return
}
