package main

import (
	"cw1/database"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func showWrongAction() {
	fmt.Printf("Provide action: %s, %s or %s", "add", "remove", "list")
}

func main() {
	if len(os.Args) == 1 {
		showWrongAction()
		return
	}

	fmt.Println("Executing action:", os.Args[1])

	switch os.Args[1] {
	case "add":
		name, description, priority, err := initAddFlags(os.Args[2:])

		if err != nil {
			log.Panic("There was an issue with parsing your data: ", err, ". Please fix your input")
		}

		fmt.Printf("Created action: %s (%s) with priority %s", name, description, priority)

		_, err = database.AddItem(name, description, priority)

		if err != nil {
			log.Panic(err)
		}
	case "remove":
		name, err := initRemoveFlags(os.Args[2:])
		if err != nil {
			log.Panic("Could not run remove: ", err)
		}
		err = database.Delete(name)

		if err != nil {
			log.Panic("Could not run remove: ", err)
		}
	case "list":
		items, err := database.GetAllItems()
		if err != nil {
			log.Panic("Error when receiving data from db:", err, ". Program is going to stop")
		}

		for index, item := range items {
			fmt.Println(index + 1, ": ", item.Name, "\t", item.Description, "\t", item.Priority)
		}

	default:
		showWrongAction()
	}
}

func initAddFlags(args []string) (name, description string, priority database.Priority, err error) {
	addFlags := flag.NewFlagSet("add", flag.ExitOnError)

	priority = database.LOW

	var priorityStr string
	priorityDesc := fmt.Sprintf(
		"Pirority of the issue. Possible values: %s, %s, %s",
		database.LOW,
		database.MEDIUM,
		database.HIGH,
	)
	addFlags.StringVar(&priorityStr, "priority", string(database.MEDIUM), priorityDesc)
	addFlags.StringVar(&name, "name", "", "Name of the issue. (REQUIRED)")
	addFlags.StringVar(&description, "description", "", "Description of the issue.")

	err = addFlags.Parse(args)
	if err != nil {
		return
	}

	priority, isOk := database.IsValidPriority(priorityStr)

	switch {
	case len(name) == 0:
		err = errors.New("name of the name is empty but required")
		return
	case len(description) == 0:
		err = fmt.Errorf("name of the %s is empty but required", "description")
		return
	case !isOk:
		err = fmt.Errorf("priority %s is not allowed", priorityStr)
		return
	}

	// TODO: Add estimated time to complete

	return
}


func initRemoveFlags(args []string) (name string, err error) {
	removeFlags := flag.NewFlagSet("remove", flag.ExitOnError)

	removeFlags.StringVar(&name, "name", "", "Name of the issue to delete")

	err = removeFlags.Parse(args)

	if err != nil {
		return
	}

	if len(name) == 0 {
		err = errors.New("name can not be empty")
		return
	}

	return
}