package database

func AddItem (name, description string, priority Priority) (err error, item *DBItem) {
	err, item = addToDB(name, description, priority)

	if err != nil {
		return
	}

	return
}