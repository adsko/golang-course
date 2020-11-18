package database

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

func (d DBItem) ID() (jsonField string, value interface{}) {
	jsonField = "id"
	value = d.Name

	return
}
