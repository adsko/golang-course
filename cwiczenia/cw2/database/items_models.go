package database

import "errors"

type Priority string

const(
	LOW Priority = "low"
	MEDIUM Priority = "medium"
	HIGH Priority = "high"
)

func IsValidPriority (s string) (Priority, bool) {
	switch Priority(s) {
	case LOW, MEDIUM, HIGH:
		return Priority(s), true
	}

	return LOW, false
}

func PriorityAsString(i uint) (string, error) {
	switch i {
	case 0:
		return string(LOW), nil
	case 1:
		return string(MEDIUM), nil
	case 2:
		return string(HIGH), nil
	default:
		return "", errors.New("not defined priority")
	}
}

type DBItem struct {
	Name string
	Description string
	Priority uint
}

// It's easier to make one function for creating in case of further structure expansion.
// Because structure fields are optional when creating, it's better to do it in this way.
func makeDBItem(name, description string, priority uint) DBItem {
	return DBItem{
		name,
		description,
		priority,
	}
}

func getItemOnlyWithID(name string) DBItem {
	return DBItem{
		name,
		"",
		1,
	}
}

func (d DBItem) ID() (jsonField string, value interface{}) {
	jsonField = "Name"
	value = d.Name

	return
}
